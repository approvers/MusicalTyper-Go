package beatmap

import (
	"fmt"
	DrawHelper "musicaltyper-go/game/draw/helper"
	"strings"
)

type Sentence struct {
	OriginalSentence string
	HiraganaSentence string

	SolvedSentence []*Character

	CurrentCharacterIndex int

	TypeCount  int
	MissCount  int
	IsFinished bool
}

type Character struct {
	//例: ち
	Character string

	//ti chi
	RomaStyles []string

	//どこまで入力したか。
	TypingIndex int
}

func NewSentence(OriginalSentence, HiraganaSentence string) *Sentence {
	Result := new(Sentence)
	Result.HiraganaSentence = HiraganaSentence
	Result.OriginalSentence = OriginalSentence
	Result.SolvedSentence = Solve(HiraganaSentence)
	Result.IsFinished = false
	return Result
}

func (s *Sentence) GetTypedText() string {
	return DrawHelper.Substring(s.HiraganaSentence, 0, s.CurrentCharacterIndex)
}

func (s *Sentence) GetRemainingText() string {
	return DrawHelper.Substring(s.HiraganaSentence, s.CurrentCharacterIndex, DrawHelper.Length(s.HiraganaSentence))
}

func (s *Sentence) GetTypedRoma() string {
	if len(s.SolvedSentence) == 0 {
		return ""
	}

	Result := ""
	for i := 0; i < s.CurrentCharacterIndex; i++ {
		Result += s.SolvedSentence[i].RomaStyles[0]
	}

	if len(s.SolvedSentence) > s.CurrentCharacterIndex {
		CurrentCharacter := s.SolvedSentence[s.CurrentCharacterIndex]
		Result += DrawHelper.Substring(CurrentCharacter.RomaStyles[0], 0, CurrentCharacter.TypingIndex)
	}

	return Result
}

func (s *Sentence) GetRemainingRoma() string {
	if len(s.SolvedSentence) == 0 {
		return ""
	}

	Result := ""
	for _, v := range s.SolvedSentence {
		Result += v.RomaStyles[0]
	}
	return DrawHelper.Substring(Result, DrawHelper.Length(s.GetTypedRoma()), DrawHelper.Length(Result))
}

func (s *Sentence) GetRoma() string {
	Result := ""
	for i := 0; i < s.CurrentCharacterIndex; i++ {
		Result += s.SolvedSentence[i].RomaStyles[0]
	}
	return Result
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

func (s *Sentence) JudgeKeyInput(input string) (ok, isThisSentenceEnded bool) {
	CurrentChar := s.SolvedSentence[s.CurrentCharacterIndex]
	fmt.Printf("Char:%d Typ: %d\n", s.CurrentCharacterIndex, CurrentChar.TypingIndex)

	var (
		RemainSuggests = make([]string, 0)
		Len            = 0
		isCharEnded    = false
	)

	fmt.Print(CurrentChar.RomaStyles)
	for _, v := range CurrentChar.RomaStyles {
		if strings.Split(v, "")[CurrentChar.TypingIndex] == input {
			if CurrentChar.TypingIndex+1 == len(v) {
				isCharEnded = true
				break
			}

			RemainSuggests = append(RemainSuggests, v)
		}
	}
	fmt.Println(Len)

	if !isCharEnded {
		if len(RemainSuggests) == 0 {
			return false, false
		}

		CurrentChar.RomaStyles = RemainSuggests
		CurrentChar.TypingIndex++
		return true, false
	} else {
		s.CurrentCharacterIndex++
		if len(s.SolvedSentence) == s.CurrentCharacterIndex {
			s.IsFinished = true
			return true, true
		}
		return true, false
	}
}
