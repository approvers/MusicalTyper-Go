package beatmap

import (
	"bufio"
	"fmt"
	"io/ioutil"
	Logger "musicaltyper-go/game/logger"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Beatmap has its properties, defined notes, and existing sections
type Beatmap struct {
	Properties map[string]string
	Notes      []*Note
	Sections   []*Section
}

// NewBeatmap makes empty Beatmap
func NewBeatmap() *Beatmap {
	Result := Beatmap{}

	Result.Properties = map[string]string{}
	Result.Notes = make([]*Note, 0)
	Result.Sections = make([]*Section, 0)
	return &Result
}

// LoadMap makes Beatmap from file in passed path
func LoadMap(path string) *Beatmap {
	logger := Logger.NewLogger("LoadMap")

	File, Err := os.OpenFile(path, os.O_RDONLY, 0666)
	logger.CheckError(Err)
	defer File.Close()

	var Scanner *bufio.Scanner
	if detectEncoding(path) {
		Scanner = bufio.NewScanner(File)
	} else {
		Scanner = bufio.NewScanner(transform.NewReader(File, japanese.ShiftJIS.NewDecoder()))
	}

	var (
		Result                  = NewBeatmap()
		LineCount               = 1
		CurrentMinute           = 0
		CurrentTime     float64 = 0
		isInSong                = false
		isDetectedError         = false
		TempLyric               = ""
		TempPron                = ""
	)

	Scanner.Split(bufio.ScanLines)
	for Scanner.Scan() {
		Line := strings.TrimSpace(Scanner.Text())
		if Line == "" || strings.HasPrefix(Line, "#") {
			continue
		}

		isSpecialCommand := strings.HasPrefix(Line, "[") && strings.HasSuffix(Line, "]")
		if !isInSong {
			switch {
			case strings.HasPrefix(Line, ":"):
				Line = Line[1:]
				Split := strings.Split(Line, " ")
				//":author      test" -> ["author", "", "", ... "test"]
				Result.Properties[Split[0]] = Split[len(Split)-1]

			case isSpecialCommand:
				if Command := parseSpecialCommand(Line); Command == "start" {
					isInSong = true
				} else {
					logger.Warn(fmt.Sprintf("Line%d: Unknown command \"%s\" outside of song section.", LineCount, Command))
				}

			default:
				isDetectedError = true
				logger.FatalErrorWithoutExit(fmt.Sprintf("Line%d: Unknown text \"%s\" outside of song section.", LineCount, Line))
			}

		} else {

			switch {
			case isSpecialCommand:
				switch Command := parseSpecialCommand(Line); Command {
				case "break":
					Result.Notes = append(Result.Notes, newBlankNote(CurrentTime))
				case "end":
					Result.Notes = append(Result.Notes, endMap(CurrentTime))
					isInSong = false
				}

			case strings.HasPrefix(Line, ">>"):
				Result.Notes = append(Result.Notes, newCaptionNote(CurrentTime, Line[2:]))

			case strings.HasPrefix(Line, "|"):
				CurrentMinute, _ = strconv.Atoi(Line[1:])

			case strings.HasPrefix(Line, "*"):
				if TempLyric != "" {
					if TempPron != "" {
						Result.Notes = append(Result.Notes, newNote(CurrentTime, TempLyric, TempPron))
					} else {
						isDetectedError = true
						logger.FatalErrorWithoutExit(fmt.Sprintf("Line%d: Lyric provided, but pronunciation data doensn't provided.", LineCount))
					}
				}
				TempLyric = ""
				TempPron = ""

				NewSec, Error := strconv.ParseFloat(Line[1:], 64)
				logger.CheckError(Error)
				CurrentTime = float64(60*CurrentMinute) + NewSec

			case strings.HasPrefix(Line, "@"):
				Result.Sections = append(Result.Sections, newSection(CurrentTime, Line[1:]))

			case strings.HasPrefix(Line, ":"):
				TempPron += Line[1:]

			default:
				TempLyric += Line

			}
		}
		LineCount++
	}
	if v, Exist := Result.Properties["song_data"]; !Exist {
		isDetectedError = true
		logger.FatalErrorWithoutExit("The song_data property is not defined. It is required to play a song.")
	} else {
		Path, Error := filepath.Abs(path)
		logger.CheckError(Error)

		Dir := filepath.Dir(Path)
		SongPath := filepath.Join(Dir, v)
		Info, Error := os.Stat(SongPath)
		if Error != nil || !Info.Mode().IsRegular() {
			isDetectedError = true
			logger.FatalErrorWithoutExit("The path which is in song_data property is invalid.")
			logger.FatalErrorWithoutExit("Path: " + SongPath)
		} else {
			Result.Properties["song_data"] = SongPath
		}
	}

	if isDetectedError {
		logger.FatalError("Please fix above issues. Exiting.")
	}

	return Result
}

//Because regex in Golang is slow.
func parseSpecialCommand(s string) string {
	command := strings.ReplaceAll(s, "[", "")
	command = strings.ReplaceAll(command, "]", "")
	command = strings.TrimSpace(command)
	return command
}

//return -> isUTF-8. true means UTF-8, false means Shift_JIS.
//if detected something else, consider as an error to prevent depressing errors.
func detectEncoding(path string) bool {
	Logger := Logger.NewLogger("detectEncoding")

	Data, Err := ioutil.ReadFile(path)
	Logger.CheckError(Err)

	Encoding, Err := chardet.NewTextDetector().DetectBest(Data)
	if Err != nil {
		Logger.Warn("Failed to detect Beatmap's text encoding. Trying to parse as UTF-8.")
		return true //consider as UTF-8
	}

	switch Encoding.Charset {
	case "UTF-8":
		return true
	case "Shift_JIS":
		Logger.Warn("Detected that Beatmap's text encoding is Shift_JIS. Please consider using UTF-8.")
		return false
	default:
		Logger.FatalError("Detected that Beatmap's text encoding is not neither UTF-8 or Shift_JIS. Exiting.")
		return false //won't reach here.
	}
}
