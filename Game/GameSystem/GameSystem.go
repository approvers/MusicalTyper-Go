package GameSystem

import (
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/DrawComponents/DrawManager"
	"MusicalTyper-Go/Game/DrawComponents/Effects"
	"MusicalTyper-Go/Game/DrawHelper"
	"MusicalTyper-Go/Game/GameState"
	"MusicalTyper-Go/Game/SEHelper"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

/*
Q. Why is this separated between GameState's method?
A. To prevent from cyclic dependencies.
*/

var (
	successEffect = Effects.NewSlideFadeoutText(
		"Pass",
		DrawHelper.GetMoreBlackishColor(Constants.GreenThinColor, 50),
		DrawHelper.AlphabetFont,
		-150, -383, 10)

	pointOnKeyboardEffect = *Effects.NewAbsoluteFadeout("", Constants.TextColor, DrawHelper.FullFont, 0, 0, 15)
)

func ParseKeyInput(renderer *sdl.Renderer, s *GameState.GameState, code sdl.Keycode, PrintLyric bool) {
	if !((code >= 'a' && code <= 'z') || (code >= '0' && code <= '9') || code == '[' || code == ']' || code == ',' || code == '.' || code == ' ') {
		return
	}

	if s.IsInputDisabled {
		SEHelper.Play(SEHelper.UnneccesarySE)
		return
	}

	KeyChar := string(code)
	CurrentSentence := s.Beatmap.Notes[s.CurrentSentenceIndex].Sentence
	ok, SentenceEnded := CurrentSentence.IsExceptedKey(KeyChar)

	Point := s.AddPoint(ok, SentenceEnded)

	if !ok {
		SEHelper.Play(SEHelper.FailedSE)
		return
	}

	s.CountKeyType()
	SEHelper.Play(SEHelper.SuccessSE)
	DrawManager.AddEffector(DrawManager.FOREGROUND, 30, successEffect)

	if !PrintLyric {
		x, y := DrawHelper.GetKeyPlace(KeyChar)
		text := fmt.Sprintf("+%d", Point)
		textwidth, _ := DrawHelper.GetTextSize(renderer, DrawHelper.FullFont, text, Constants.BlueThickColor)
		x -= textwidth / 2

		pointOnKeyboardEffect.Text = text
		pointOnKeyboardEffect.BaseX = x
		pointOnKeyboardEffect.BaseY = y
		DrawManager.AddEffector(DrawManager.FOREGROUND, 30, pointOnKeyboardEffect)
	}

	//if isEndOfSentence {
	//	s.CurrentSentenceIndex++
	//}
}
