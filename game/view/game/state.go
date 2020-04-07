package mainview

import (
	"fmt"
	"math"
	Beatmap "musicaltyper-go/game/beatmap"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	Rank "musicaltyper-go/game/rank"
	"musicaltyper-go/game/sehelper"
	"musicaltyper-go/game/view/game/component/effects"
	"musicaltyper-go/game/view/game/component/keyboard"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// GameState has whole of state to manage game logic
type GameState struct {
	Beatmap *Beatmap.Beatmap

	CurrentSentenceIndex int
	CurrentTime          float64

	Combo        int
	Point        int
	PerfectPoint int

	TotalCorrectCount int
	TotalMissCount    int

	IsInputDisabled bool

	KeyInputs []time.Time

	KeyInputSpeedMutex sync.Mutex
	KeyInputSpeedSum   int
	KeyInputSpeedCount int

	stopTypeSpeedCalcDaemon chan bool
}

// NewGameState makes GameState from Beatmap
func NewGameState(Map *Beatmap.Beatmap) *GameState {
	r := new(GameState)
	r.Beatmap = Map
	r.KeyInputs = make([]time.Time, 0)
	r.IsInputDisabled = Map.Notes[0].Type != Beatmap.NORMAL
	r.stopTypeSpeedCalcDaemon = make(chan bool)

	go calcTypeSpeed(r)

	return r
}

func calcTypeSpeed(state *GameState) {
	ticker := time.NewTicker(100 * time.Millisecond)
	Continue := true
	DontErase := false

	for Continue {
		select {
		case <-state.stopTypeSpeedCalcDaemon:
			Continue = false

		case <-ticker.C:
			if !state.IsInputDisabled {
				state.KeyInputSpeedMutex.Lock()

				now := time.Now()
				remains := make([]time.Time, 0, len(state.KeyInputs))

				for _, v := range state.KeyInputs {
					if (now.Sub(v).Milliseconds() < 100) && !DontErase {
						remains = append(remains, v)
					}
				}
				state.KeyInputs = remains
				state.KeyInputSpeedSum += len(remains) * 10
				state.KeyInputSpeedCount++

				fmt.Print("[CalcTypeSpeedDaemon] ThisTime:", len(remains)*10, "Average:", float64(state.KeyInputSpeedSum)/float64(state.KeyInputSpeedCount))

				if DontErase {
					fmt.Println(" Didn't Erase.")
					DontErase = false
				}

				state.KeyInputSpeedMutex.Unlock()
			} else {
				DontErase = true
			}

		}
	}
}

// Update overrides current time and updates current note
func (s *GameState) Update(CurrentTime float64) {
	s.CurrentTime = CurrentTime
	if len(s.Beatmap.Notes) > s.CurrentSentenceIndex+1 && s.Beatmap.Notes[s.CurrentSentenceIndex+1].Time <= CurrentTime {
		fmt.Println("Updated index")

		Note := s.Beatmap.Notes[s.CurrentSentenceIndex]
		CurrentSentence := Note.Sentence
		if !CurrentSentence.IsFinished && Note.Type == Beatmap.NORMAL {
			AddEffector(FOREGROUND, 120, tleTextEffect)
			AddEffector(BACKGROUND, 15, tleBackgroundEffect)
			sehelper.Play(sehelper.TleSE)
		}

		s.CurrentSentenceIndex++
		s.IsInputDisabled = s.Beatmap.Notes[s.CurrentSentenceIndex].Type != Beatmap.NORMAL
	}
}

// GetAccuracy calculates accuracy
func (s *GameState) GetAccuracy() float64 {
	if s.TotalCorrectCount == 0 {
		return 0
	}

	return float64(s.TotalCorrectCount) / float64(s.TotalMissCount+s.TotalCorrectCount)
}

// GetAchievementRate calculates achievement rate
func (s *GameState) GetAchievementRate(Limit bool) float64 {
	Acc := s.GetAccuracy()
	PerfectScore := s.PerfectPoint + s.TotalCorrectCount*45
	Score := float64(s.Point) * Acc
	if Score <= 0 {
		return 0
	}

	if Limit {
		Score = math.Min(Score, float64(PerfectScore))
	}
	return Score / float64(PerfectScore)
}

// GetRank decides player rank
func (s *GameState) GetRank() Rank.Rank {
	Rate := s.GetAchievementRate(false)
	return Rank.FromAchievementRate(Rate)
}

// CountKeyType records time when typed any key
func (s *GameState) CountKeyType() {
	s.KeyInputSpeedMutex.Lock()
	defer s.KeyInputSpeedMutex.Unlock()

	s.KeyInputs = append(s.KeyInputs, time.Now())
}

// GetKeyTypePerSecond calculates typed times per second from records from CountKeyType
func (s *GameState) GetKeyTypePerSecond() float64 {
	s.KeyInputSpeedMutex.Lock()
	defer s.KeyInputSpeedMutex.Unlock()

	if s.KeyInputSpeedCount == 0 {
		return 0
	}

	return float64(s.KeyInputSpeedSum) / float64(s.KeyInputSpeedCount)
}

// AddPoint decides and adds point with flags
func (s *GameState) AddPoint(isTypeOK, isThisSentenceEnded bool) (point int) {
	CurrentSentence := s.Beatmap.Notes[s.CurrentSentenceIndex].Sentence

	if isTypeOK {
		s.TotalCorrectCount++
		s.Combo++

		point = int(Constants.OneCharPoint * 10 * s.GetKeyTypePerSecond() * float64(s.Combo/10))
		s.Point += point
		s.PerfectPoint += Constants.OneCharPoint * 10 * Constants.IdealTypeSpeed * s.Combo / 10

		if isThisSentenceEnded {
			s.PerfectPoint += Constants.ClearPoint + Constants.PerfectPoint
			s.Point += Constants.ClearPoint
			if CurrentSentence.MissCount == 0 {
				s.Point += Constants.PerfectPoint
			}
		} else {
			CurrentSentence.TypeCount++
		}
	} else {
		s.TotalMissCount++
		CurrentSentence.MissCount++
		point = Constants.MissPoint
		s.Point += point
		s.Combo = 0
	}
	return
}

// AddTLEPoint adds points when failed to type all
func (s *GameState) AddTLEPoint() {
	CurrentSentence := s.Beatmap.Notes[s.CurrentSentenceIndex].Sentence
	TextLen := len(CurrentSentence.GetRemainingRoma())

	s.Point += Constants.CouldntTypeCount * TextLen
	s.PerfectPoint += Constants.OneCharPoint*TextLen*40 + Constants.ClearPoint + Constants.PerfectPoint
	s.TotalMissCount += TextLen
	CurrentSentence.MissCount += TextLen
}

// ParseKeyInput handles key input event from sdl
func (s *GameState) ParseKeyInput(renderer *sdl.Renderer, code sdl.Keycode, PrintLyric bool) {
	if !((code >= 'a' && code <= 'z') || (code >= '0' && code <= '9') || code == '[' || code == ']' || code == ',' || code == '.' || code == ' ' || code == '-') {
		return
	}

	if s.IsInputDisabled {
		sehelper.Play(sehelper.UnneccesarySE)
		return
	}

	KeyChar := string(code)
	CurrentSentence := s.Beatmap.Notes[s.CurrentSentenceIndex].Sentence
	ok, SentenceEnded := CurrentSentence.JudgeKeyInput(KeyChar)

	Point := s.AddPoint(ok, SentenceEnded)

	if !ok {
		AddEffector(FOREGROUND, 120, missTypeTextEffect)
		AddEffector(BACKGROUND, 15, missTypeBackgroundEffect)
		sehelper.Play(sehelper.FailedSE)
		return
	}

	s.CountKeyType()
	AddEffector(FOREGROUND, 30, successEffect)

	if !PrintLyric {
		KeyPos := keyboard.GetKeyPos(KeyChar)
		text := fmt.Sprintf("+%d", Point)
		textwidth := helper.GetTextSize(renderer, helper.FullFont, text, Constants.BlueThickColor).W()
		KeyPos = pos.FromXY(KeyPos.X()-textwidth/2, KeyPos.Y())

		AddEffector(FOREGROUND, 30, effects.NewAbsoluteFadeout(
			text,
			Constants.BlueThickColor,
			helper.FullFont,
			KeyPos, 15,
		))
	}

	if SentenceEnded {
		s.IsInputDisabled = true

		if CurrentSentence.MissCount == 0 {
			AddEffector(FOREGROUND, 120, acTextEffect)
			AddEffector(BACKGROUND, 15, acBackgroundEffect)
			sehelper.Play(sehelper.AcSE)
		} else {
			AddEffector(FOREGROUND, 120, waTextEffect)
			AddEffector(BACKGROUND, 15, waBackgroundEffect)
			sehelper.Play(sehelper.WaSE)
		}
	} else {
		if s.GetKeyTypePerSecond() > 4 {
			sehelper.Play(sehelper.FastSE)
		} else {
			sehelper.Play(sehelper.SuccessSE)
		}
	}
}
