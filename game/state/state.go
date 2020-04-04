package state

import (
	"fmt"
	"math"
	Beatmap "musicaltyper-go/game/beatmap"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/view/mainview"
	Rank "musicaltyper-go/game/rank"
	"musicaltyper-go/game/sehelper"
	"time"
)

// ResultType is unused
type ResultType int

// Result is unused
type Result struct {
	Count     int
	MissCount int
}

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
}

// NewGameState makes GameState from Beatmap
func NewGameState(Map *Beatmap.Beatmap) *GameState {
	r := new(GameState)
	r.Beatmap = Map
	r.KeyInputs = make([]time.Time, 0)
	r.IsInputDisabled = Map.Notes[0].Type != Beatmap.NORMAL
	return r
}

// Update overrides current time and updates current note
func (s *GameState) Update(CurrentTime float64) {
	s.CurrentTime = CurrentTime
	if len(s.Beatmap.Notes) > s.CurrentSentenceIndex+1 && s.Beatmap.Notes[s.CurrentSentenceIndex+1].Time <= CurrentTime {
		fmt.Println("Updated index")

		Note := s.Beatmap.Notes[s.CurrentSentenceIndex]
		CurrentSentence := Note.Sentence
		if !CurrentSentence.IsFinished && Note.Type == Beatmap.NORMAL {
			mainview.AddEffector(mainview.FOREGROUND, 120, tleTextEffect())
			mainview.AddEffector(mainview.BACKGROUND, 15, tleBackgroundEffect())
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
	s.KeyInputs = append(s.KeyInputs, time.Now())
}

// GetKeyTypePerSecond calculates typed times per second from records from CountKeyType
func (s *GameState) GetKeyTypePerSecond() int {
	now := time.Now()
	remains := make([]time.Time, 0, len(s.KeyInputs))

	for _, v := range s.KeyInputs {
		if now.Sub(v).Milliseconds() < 1000 {
			remains = append(remains, v)
		}
	}
	s.KeyInputs = remains
	return len(remains)
}

// AddPoint decides and adds point with flags
func (s *GameState) AddPoint(isTypeOK, isThisSentenceEnded bool) (point int) {
	CurrentSentence := s.Beatmap.Notes[s.CurrentSentenceIndex].Sentence

	if isTypeOK {
		s.TotalCorrectCount++
		s.Combo++

		point = int(float64(Constants.OneCharPoint*10*s.GetKeyTypePerSecond()) * float64(s.Combo/10))
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
