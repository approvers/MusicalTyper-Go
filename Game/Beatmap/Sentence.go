package Beatmap

import (
	"MusicalTyper-Go/Game/Util"
)

type Sentence struct {
	OriginalSentence      string
	HiraganaSentence      string
	SolvedSentence        []Character
	CurrentCharacterIndex int
	TypeCount             int
	MissCount             int
}

type Character struct {
	//例: ち
	Character string

	//ti chi
	RomaStyles []string
	//特定されたローマ字。 未特定の場合は-1 0:ti 1:chi
	RomaStyleIndex int
	//ローマ字を特定するポイント。 tiとchiなら最初の文字で特定できるので0。siとshiなら、2文字目で特定できるので1。
	RomaStyleCheckPoint int
	//どこまで入力したか。
	TypingIndex int
}

func NewSentence(OriginalSentence, HiraganaSentence string) Sentence {
	Result := Sentence{}
	Result.HiraganaSentence = HiraganaSentence
	Result.OriginalSentence = OriginalSentence
	Result.SolvedSentence = Solve(HiraganaSentence)
	return Result
}

func (s *Sentence) GetTypedText() string {
	return Util.Substring(s.HiraganaSentence, 0, s.CurrentCharacterIndex)
}

func (s *Sentence) GetRemainingText() string {
	return Util.Substring(s.HiraganaSentence, s.CurrentCharacterIndex, Util.Length(s.HiraganaSentence))
}

func (s *Sentence) GetTypedRoma() string {
	Result := ""
	for i := 0; i < s.CurrentCharacterIndex; i++ {
		CurrentCharacter := s.SolvedSentence[i]
		var RomaIndex int

		if CurrentCharacter.RomaStyleIndex == -1 {
			RomaIndex = 0
		} else {
			RomaIndex = CurrentCharacter.RomaStyleIndex
		}

		if i == s.CurrentCharacterIndex {
			Result += Util.Substring(CurrentCharacter.RomaStyles[RomaIndex], 0, CurrentCharacter.TypingIndex)
		} else {
			Result += CurrentCharacter.RomaStyles[RomaIndex]
		}
	}

	return Result
}

func (s *Sentence) GetRemainingRoma() string {
	Result := ""
	for i := 0; i < len(s.SolvedSentence); i++ {
		CurrentCharacter := s.SolvedSentence[i]
		var RomaIndex int

		if CurrentCharacter.RomaStyleIndex == -1 {
			RomaIndex = 0
		} else {
			RomaIndex = CurrentCharacter.RomaStyleIndex
		}

		Result += CurrentCharacter.RomaStyles[RomaIndex]
	}
	return Util.Substring(Result, Util.Length(s.GetTypedRoma()), Util.Length(Result))
}

func (s *Sentence) GetRoma() string {
	return s.GetTypedRoma() + s.GetRemainingRoma()
}

func (s *Sentence) GetAccuracy() float64 {
	Misses, Types := 1, 1
	if s.TypeCount > 0 {
		Types = s.TypeCount
	}
	if s.MissCount > 0 {
		Misses = s.MissCount
	}

	return float64(Misses) / float64(Types)
}
