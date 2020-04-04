package state

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/component/effects"
	"musicaltyper-go/game/draw/component/keyboard"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/draw/view/mainview"
	SEHelper "musicaltyper-go/game/sehelper"

	"github.com/veandco/go-sdl2/sdl"
)

/*
Q. Why is this separated between GameState's method?
A. To prevent from cyclic dependencies.
*/

func successEffect() component.DrawableEffect {
	return effects.NewSlideFadeoutText(
		"Pass",
		Constants.GreenThinColor.Darker(50),
		helper.AlphabetFont,
		pos.FromXY(-150, -383), 10,
	)
}

func pointOnKeyboardEffect() component.DrawableEffect {
	return effects.NewAbsoluteFadeout(
		"",
		Constants.BlueThickColor,
		helper.FullFont,
		pos.FromXY(0, 0), 15,
	)
}

func acTextEffect() component.DrawableEffect {
	return effects.NewSlideFadeoutText(
		"AC",
		Constants.GreenThickColor,
		helper.AlphabetFont,
		pos.FromXY(-170, -222), 20,
	)
}

func acBackgroundEffect() component.DrawableEffect {
	return effects.NewBlinkRect(
		Constants.GreenThinColor.Brighter(50),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)
}

func waTextEffect() component.DrawableEffect {
	return effects.NewSlideFadeoutText(
		"WA",
		Constants.BlueThickColor.Brighter(100),
		helper.AlphabetFont,
		pos.FromXY(-170, -222), 20,
	)
}

func waBackgroundEffect() component.DrawableEffect {
	return effects.NewBlinkRect(
		Constants.BlueThickColor.Brighter(100),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)
}

func missTypeTextEffect() component.DrawableEffect {
	return effects.NewSlideFadeoutText(
		"MISS",
		Constants.RedColor.Brighter(50),
		helper.AlphabetFont,
		pos.FromXY(-150, -222), 10,
	)
}

func missTypeBackgroundEffect() component.DrawableEffect {
	return effects.NewBlinkRect(
		color.FromRGB(255, 200, 200),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)
}

func tleTextEffect() component.DrawableEffect {
	return effects.NewSlideFadeoutText(
		"TLE",
		Constants.RedColor.Darker(50),
		helper.AlphabetFont,
		pos.FromXY(-150, -222), 10,
	)
}

func tleBackgroundEffect() component.DrawableEffect {
	return effects.NewBlinkRect(
		Constants.RedColor.Brighter(50),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)
}

// ParseKeyInput handles key input event from sdl
func (s *GameState) ParseKeyInput(renderer *sdl.Renderer, code sdl.Keycode, PrintLyric bool) {
	if !((code >= 'a' && code <= 'z') || (code >= '0' && code <= '9') || code == '[' || code == ']' || code == ',' || code == '.' || code == ' ' || code == '-') {
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
		mainview.AddEffector(mainview.FOREGROUND, 120, missTypeTextEffect())
		mainview.AddEffector(mainview.BACKGROUND, 15, missTypeBackgroundEffect())
		SEHelper.Play(SEHelper.FailedSE)
		return
	}

	s.CountKeyType()
	mainview.AddEffector(mainview.FOREGROUND, 30, successEffect())

	if !PrintLyric {
		KeyPos := keyboard.GetKeyPos(KeyChar)
		text := fmt.Sprintf("+%d", Point)
		textwidth := helper.GetTextSize(renderer, helper.FullFont, text, Constants.BlueThickColor).W()
		KeyPos = pos.FromXY(KeyPos.X()-textwidth/2, KeyPos.Y())

		mainview.AddEffector(mainview.FOREGROUND, 30, effects.NewAbsoluteFadeout(
			text,
			Constants.BlueThickColor,
			helper.FullFont,
			KeyPos, 15,
		))
	}

	if SentenceEnded {
		s.IsInputDisabled = true

		if CurrentSentence.MissCount == 0 {
			mainview.AddEffector(mainview.FOREGROUND, 120, acTextEffect())
			mainview.AddEffector(mainview.BACKGROUND, 15, acBackgroundEffect())
			SEHelper.Play(SEHelper.AcSE)
		} else {
			mainview.AddEffector(mainview.FOREGROUND, 120, waTextEffect())
			mainview.AddEffector(mainview.BACKGROUND, 15, waBackgroundEffect())
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
