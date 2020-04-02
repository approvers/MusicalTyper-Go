package system

import (
	"fmt"
	Beatmap "musicaltyper-go/game/beatmap"
	Constants "musicaltyper-go/game/constants"
	Effects "musicaltyper-go/game/draw/component/effects"
	DrawHelper "musicaltyper-go/game/draw/helper"
	DrawManager "musicaltyper-go/game/draw/manager"
	SEHelper "musicaltyper-go/game/sehelper"
	GameState "musicaltyper-go/game/state"

	"github.com/veandco/go-sdl2/sdl"
)

/*
Q. Why is this separated between GameState's method?
A. To prevent from cyclic dependencies.
*/

var (
	successEffect = Effects.NewSlideFadeoutText(
		"Pass",
		*DrawHelper.GetMoreBlackishColor(Constants.GreenThinColor, 50),
		DrawHelper.AlphabetFont,
		-150, -383, 10,
	)

	pointOnKeyboardEffect = Effects.NewAbsoluteFadeout(
		"",
		*Constants.BlueThickColor,
		DrawHelper.FullFont,
		0, 0, 15,
	)

	acTextEffect = Effects.NewSlideFadeoutText(
		"AC",
		*Constants.GreenThickColor,
		DrawHelper.AlphabetFont,
		-170, -222, 20,
	)

	acBackgroundEffect = Effects.NewBlinkRect(
		*DrawHelper.GetMoreWhitishColor(Constants.GreenThinColor, 50),
		&sdl.Rect{X: 0, Y: 60, W: Constants.WindowWidth, H: 130},
	)

	waTextEffect = Effects.NewSlideFadeoutText(
		"WA",
		*DrawHelper.GetMoreWhitishColor(Constants.BlueThickColor, 100),
		DrawHelper.AlphabetFont,
		-170, -222, 20,
	)

	waBackgroundEffect = Effects.NewBlinkRect(
		*DrawHelper.GetMoreWhitishColor(Constants.BlueThickColor, 100),
		&sdl.Rect{X: 0, Y: 60, W: Constants.WindowWidth, H: 130},
	)

	missTypeTextEffect = Effects.NewSlideFadeoutText(
		"MISS",
		*DrawHelper.GetMoreWhitishColor(Constants.RedColor, 50),
		DrawHelper.AlphabetFont,
		-150, -222, 10,
	)

	missTypeBackgroundEffect = Effects.NewBlinkRect(
		sdl.Color{255, 200, 200, 255},
		&sdl.Rect{X: 0, Y: 60, W: Constants.WindowWidth, H: 130},
	)

	tleTextEffect = Effects.NewSlideFadeoutText(
		"TLE",
		*DrawHelper.GetMoreBlackishColor(Constants.RedColor, 50),
		DrawHelper.AlphabetFont,
		-150, -222, 10,
	)

	tleBackgroundEffect = Effects.NewBlinkRect(
		*DrawHelper.GetMoreWhitishColor(Constants.RedColor, 50),
		&sdl.Rect{X: 0, Y: 60, W: Constants.WindowWidth, H: 130},
	)
)

//Sync between GameState's current time and realtime, then Update current note.
func Update(s *GameState.GameState, CurrentTime float64) {
	s.CurrentTime = CurrentTime
	if len(s.Beatmap.Notes) > s.CurrentSentenceIndex+1 && s.Beatmap.Notes[s.CurrentSentenceIndex+1].Time <= CurrentTime {
		fmt.Println("Updated index")

		Note := s.Beatmap.Notes[s.CurrentSentenceIndex]
		CurrentSentence := Note.Sentence
		if !CurrentSentence.IsFinished && Note.Type == Beatmap.NORMAL {
			DrawManager.AddEffector(DrawManager.FOREGROUND, 120, tleTextEffect)
			DrawManager.AddEffector(DrawManager.BACKGROUND, 15, tleBackgroundEffect)
			SEHelper.Play(SEHelper.TleSE)
		}

		s.CurrentSentenceIndex++
		s.IsInputDisabled = s.Beatmap.Notes[s.CurrentSentenceIndex].Type != Beatmap.NORMAL
	}
}

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
	ok, SentenceEnded := CurrentSentence.JudgeKeyInput(KeyChar)

	Point := s.AddPoint(ok, SentenceEnded)

	if !ok {
		DrawManager.AddEffector(DrawManager.FOREGROUND, 120, missTypeTextEffect)
		DrawManager.AddEffector(DrawManager.BACKGROUND, 15, missTypeBackgroundEffect)
		SEHelper.Play(SEHelper.FailedSE)
		return
	}

	s.CountKeyType()
	DrawManager.AddEffector(DrawManager.FOREGROUND, 30, successEffect)

	if !PrintLyric {
		x, y := DrawHelper.GetKeyPos(KeyChar)
		text := fmt.Sprintf("+%d", Point)
		textwidth, _ := DrawHelper.GetTextSize(renderer, DrawHelper.FullFont, text, Constants.BlueThickColor)
		x -= textwidth / 2

		pointOnKeyboardEffect.Text = text
		pointOnKeyboardEffect.BaseX = x
		pointOnKeyboardEffect.BaseY = y
		DrawManager.AddEffector(DrawManager.FOREGROUND, 30, *pointOnKeyboardEffect)
	}

	if SentenceEnded {
		if CurrentSentence.MissCount == 0 {
			DrawManager.AddEffector(DrawManager.FOREGROUND, 120, acTextEffect)
			DrawManager.AddEffector(DrawManager.BACKGROUND, 15, acBackgroundEffect)
			SEHelper.Play(SEHelper.AcSE)
		} else {
			DrawManager.AddEffector(DrawManager.FOREGROUND, 120, waTextEffect)
			DrawManager.AddEffector(DrawManager.BACKGROUND, 15, waBackgroundEffect)
			SEHelper.Play(SEHelper.WaSE)
		}
	} else {
		if s.GetKeyTypePerSecond() > 4 {
			SEHelper.Play(SEHelper.FastSE)
		} else {
			SEHelper.Play(SEHelper.SuccessSE)
		}
	}

	//if isEndOfSentence {
	//	s.CurrentSentenceIndex++
	//}
}
