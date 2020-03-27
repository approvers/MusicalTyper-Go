package Game

import (
	"MusicalTyper-Go/Game/Beatmap"
	"MusicalTyper-Go/Game/Constants"
	"fmt"
	"math"
)

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

	KeyInputs []int
}

func NewGameState(Map *Beatmap.Beatmap) *GameState {
	r := GameState{}
	r.Beatmap = Map
	r.KeyInputs = make([]int, 0)
	return &r
}

//Sync between GameState's current time and realtime, then Update current note.
func (s *GameState) Update(CurrentTime float64) {
	s.CurrentTime = CurrentTime
	if len(s.Beatmap.Notes) > s.CurrentSentenceIndex+1 && s.Beatmap.Notes[s.CurrentSentenceIndex+1].Time <= CurrentTime {
		fmt.Println("Updated index")
		s.CurrentSentenceIndex++
	}
}

func (s *GameState) GetSentenceTimeRemainRatio() float64 {
	if len(s.Beatmap.Notes) <= s.CurrentSentenceIndex+1 {
		return 1
	}

	CurrentSentenceStartTime := s.Beatmap.Notes[s.CurrentSentenceIndex].Time
	NextSentenceStartTime := s.Beatmap.Notes[s.CurrentSentenceIndex+1].Time
	CurrentSentenceDuration := NextSentenceStartTime - CurrentSentenceStartTime
	CurrentTimeInCurrentSentence := CurrentSentenceDuration - s.CurrentTime + CurrentSentenceStartTime
	return CurrentTimeInCurrentSentence / CurrentSentenceDuration
}

func (s *GameState) GetAccuracy() float64 {
	Misses, Types := 0, 1
	if s.TotalCorrectCount > 0 {
		Types = s.TotalCorrectCount
	}
	if s.TotalMissCount > 0 {
		Misses = s.TotalMissCount
	}

	return float64(Misses) / float64(Types)
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

func (s *GameState) GetKeyTypePerSecond() float64 {
	if len(s.KeyInputs) == 0 {
		return 0
	}
	Sum := 0
	for _, v := range s.KeyInputs {
		Sum += v
	}
	return 1.0 / (float64(Sum) / float64(len(s.KeyInputs)))
}
