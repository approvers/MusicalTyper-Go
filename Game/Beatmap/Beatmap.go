package Beatmap

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"MusicalTyper-Go/Game/Logger"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Beatmap struct {
	Properties map[string]string
	Notes      []Note
	Zones      []Zone
	Sections   []Section
}

func NewBeatmap() *Beatmap {
	Result := Beatmap{}

	Result.Properties = map[string]string{}
	Result.Notes = make([]Note, 0)
	Result.Zones = make([]Zone, 0)
	Result.Sections = make([]Section, 0)
	return &Result
}

func LoadMap(path string) *Beatmap {
	Logger := Logger.NewLogger("LoadMap")

	File, Err := os.OpenFile(path, os.O_RDONLY, 0666)
	Logger.CheckError(Err)
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
					Logger.Warn(fmt.Sprintf("Line%d: Unknown command \"%s\" outside of song section.", LineCount, Command))
				}

			default:
				isDetectedError = true
				Logger.FatalErrorWithoutExit(fmt.Sprintf("Line%d: Unknown text \"%s\" outside of song section.", LineCount, Line))
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
						Logger.FatalErrorWithoutExit(fmt.Sprintf("Line%d: Lyric provided, but pronunciation data doensn't provided.", LineCount))
					}
				}
				TempLyric = ""
				TempPron = ""

				NewSec, Error := strconv.ParseFloat(Line[1:], 64)
				Logger.CheckError(Error)
				CurrentTime = float64(60*CurrentMinute) + NewSec

			case strings.HasPrefix(Line, "@"):
				Result.Sections = append(Result.Sections, newSection(CurrentTime, Line[1:]))

			case strings.HasPrefix(Line, "!"):
				Split := strings.Split(Line[1:], " ")
				Flag, ZoneName := Split[0], Split[1]
				switch Flag := strings.ToLower(Flag); Flag {
				case "start":
					Result.Zones = append(Result.Zones, newZone(CurrentTime, true, ZoneName))
				case "end":
					Result.Zones = append(Result.Zones, newZone(CurrentTime, false, ZoneName))
				default:
					isDetectedError = true
					Logger.FatalErrorWithoutExit(fmt.Sprintf("Line%d: Zone flag \"%s\" is invalid. Allowed values are only \"start\" and \"end\".", LineCount, Flag))
				}

			case strings.HasPrefix(Line, ":"):
				TempPron += Line[1:]

			default:
				TempLyric += Line

			}
		}
		LineCount++
	}
	if _, Exist := Result.Properties["song_data"]; !Exist {
		isDetectedError = true
		Logger.FatalErrorWithoutExit("The song_data property is not defined. It is required to play a song.")
	}

	if isDetectedError {
		Logger.FatalError("Please fix above issues. Exiting.")
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
	Logger.CheckError(Err)

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
