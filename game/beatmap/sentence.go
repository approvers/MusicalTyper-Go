package beatmap

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"musicaltyper-go/game/draw/helper"
)

// Sentence has state for typing roman string
type Sentence struct {
	OriginalSentence string
	HiraganaSentence string

	SolvedSentence []*Character

	CurrentCharacterIndex int

	TypeCount  int
	MissCount  int
	IsFinished bool
}

// Character has japanese character, its roman styles, and typed index
type Character struct {
	//例: ち
	Character   string
	TypingIndex int
	RomaStyles  []*RomaStyle
}

// RomaStyle has a single roman style, and how many characters are typed with this style.
type RomaStyle struct {
	Forwards int
	Roma     string
}

// NewSentence makes Sentence from japanese and hiragana
func NewSentence(OriginalSentence, HiraganaSentence string) *Sentence {
	Result := new(Sentence)
	Result.HiraganaSentence = HiraganaSentence
	Result.OriginalSentence = OriginalSentence
	Result.SolvedSentence = Solve(HiraganaSentence)
	Result.IsFinished = false
	return Result
}

// GetTypedText returns substring of hiragana before typed index
func (s *Sentence) GetTypedText() string {
	return helper.Substring(s.HiraganaSentence, 0, s.CurrentCharacterIndex)
}

// GetRemainingText returns substring of hiragana after typed index
func (s *Sentence) GetRemainingText() string {
	return helper.Substring(s.HiraganaSentence, s.CurrentCharacterIndex, length(s.HiraganaSentence))
}

// GetTypedRoma returns typed roman
func (s *Sentence) GetTypedRoma() string {
	if len(s.SolvedSentence) == 0 {
		return ""
	}

	Result := ""
	i := 0
	for i < s.CurrentCharacterIndex {
		Style := s.SolvedSentence[i].RomaStyles[0]
		Result += Style.Roma

		i += Style.Forwards
	}

	if i == s.CurrentCharacterIndex && len(s.SolvedSentence) > s.CurrentCharacterIndex {
		CurrentCharacter := s.SolvedSentence[s.CurrentCharacterIndex]
		Result += helper.Substring(CurrentCharacter.RomaStyles[0].Roma, 0, CurrentCharacter.TypingIndex)
	}

	return Result
}

// GetRemainingRoma returns standard style roman string to be inputted
func (s *Sentence) GetRemainingRoma() string {
	if len(s.SolvedSentence) == 0 {
		return ""
	}

	Result := ""

	for i := 0; i < len(s.SolvedSentence); {
		Style := s.SolvedSentence[i].RomaStyles[0]
		Result += Style.Roma

		i += Style.Forwards
	}
	return helper.Substring(Result, length(s.GetTypedRoma()), length(Result))
}

// GetRoma returns whole of roman string
func (s *Sentence) GetRoma() string {
	Result := ""
	for i := 0; i < s.CurrentCharacterIndex; i++ {
		Result += s.SolvedSentence[i].RomaStyles[0].Roma
	}
	return Result
}

// GetAccuracy return accuracy of user typed
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

// JudgeKeyInput decides and mutates Sentence state by inputted characters
func (s *Sentence) JudgeKeyInput(input string) (ok, isThisSentenceEnded bool) {
	var (
		CurrentChar   = s.SolvedSentence[s.CurrentCharacterIndex]
		RemainGuesses = make([]*RomaStyle, 0)
		isCharEnded   = false
	)

	fmt.Print("Input:\"" + input + "\" Guesses: [ ")
	for _, v := range CurrentChar.RomaStyles {
		if strings.Split(v.Roma, "")[CurrentChar.TypingIndex] == input {
			fmt.Print("\"" + v.Roma + "\" ")
			RemainGuesses = append(RemainGuesses, v)
			if CurrentChar.TypingIndex+1 == len(v.Roma) {
				isCharEnded = true
				break
			}
		}
	}
	fmt.Print("]")

	if !isCharEnded {
		if len(RemainGuesses) == 0 {
			fmt.Println(" denied.")
			return false, false
		}

		CurrentChar.RomaStyles = RemainGuesses
		CurrentChar.TypingIndex++
		fmt.Println(" approved.")
		return true, false
	}
	fmt.Println(" approved.")

	s.CurrentCharacterIndex += RemainGuesses[0].Forwards

	if len(s.SolvedSentence) == s.CurrentCharacterIndex {
		s.IsFinished = true
		return true, true
	}
	return true, false
}

func length(s string) int {
	return utf8.RuneCountInString(s)
}
