package GameState

import (
	"MusicalTyper-Go/Game/Beatmap"
	"MusicalTyper-Go/Game/Constants"
	"fmt"
	"math"
	"time"
)

type ResultType int

type Result struct {
	Count     int
	MissCount int
}

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

func NewGameState(Map *Beatmap.Beatmap) *GameState {
	r := new(GameState)
	r.Beatmap = Map
	r.KeyInputs = make([]time.Time, 0)
	r.IsInputDisabled = Map.Notes[0].Type != Beatmap.NORMAL
	return r
}

//Sync between GameState's current time and realtime, then Update current note.
func (s *GameState) Update(CurrentTime float64) {
	s.CurrentTime = CurrentTime
	if len(s.Beatmap.Notes) > s.CurrentSentenceIndex+1 && s.Beatmap.Notes[s.CurrentSentenceIndex+1].Time <= CurrentTime {
		fmt.Println("Updated index")

		s.CurrentSentenceIndex++
		s.IsInputDisabled = s.Beatmap.Notes[s.CurrentSentenceIndex].Type != Beatmap.NORMAL
	}
}

func (s *GameState) GetAccuracy() float64 {
	Misses, Types := 0, 1
	if s.TotalCorrectCount > 0 {
		Types = s.TotalCorrectCount
	}
	if s.TotalMissCount > 0 {
		Misses = s.TotalMissCount
	}

	return float64(Types) / float64(Misses+Types)
}

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

func (s *GameState) GetRank() int {
	Rate := s.GetAchievementRate(false)
	for i, v := range Constants.RankPoints {
		if v < Rate*100 {
			return i
		}
	}
	return len(Constants.RankPoints) - 1
}

func (s *GameState) CountKeyType() {
	s.KeyInputs = append(s.KeyInputs, time.Now())
}

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
