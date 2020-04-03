package system

import (
	"fmt"
	Beatmap "musicaltyper-go/game/beatmap"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	Effects "musicaltyper-go/game/draw/component/effects"
	DrawHelper "musicaltyper-go/game/draw/helper"
	DrawManager "musicaltyper-go/game/draw/manager"
	"musicaltyper-go/game/draw/pos"
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
		Constants.GreenThinColor.Darker(50),
		DrawHelper.AlphabetFont,
		pos.FromXY(-150, -383), 10,
	)

	pointOnKeyboardEffect = Effects.NewAbsoluteFadeout(
		"",
		Constants.BlueThickColor,
		DrawHelper.FullFont,
		pos.FromXY(0, 0), 15,
	)

	acTextEffect = Effects.NewSlideFadeoutText(
		"AC",
		Constants.GreenThickColor,
		DrawHelper.AlphabetFont,
		pos.FromXY(-170, -222), 20,
	)

	acBackgroundEffect = Effects.NewBlinkRect(
		Constants.GreenThinColor.Brighter(50),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)

	waTextEffect = Effects.NewSlideFadeoutText(
		"WA",
		Constants.BlueThickColor.Brighter(100),
		DrawHelper.AlphabetFont,
		pos.FromXY(-170, -222), 20,
	)

	waBackgroundEffect = Effects.NewBlinkRect(
		Constants.BlueThickColor.Brighter(100),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)

	missTypeTextEffect = Effects.NewSlideFadeoutText(
		"MISS",
		Constants.RedColor.Brighter(50),
		DrawHelper.AlphabetFont,
		pos.FromXY(-150, -222), 10,
	)

	missTypeBackgroundEffect = Effects.NewBlinkRect(
		color.FromRGB(255, 200, 200),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)

	tleTextEffect = Effects.NewSlideFadeoutText(
		"TLE",
		Constants.RedColor.Darker(50),
		DrawHelper.AlphabetFont,
		pos.FromXY(-150, -222), 10,
	)

	tleBackgroundEffect = Effects.NewBlinkRect(
		Constants.RedColor.Brighter(50),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)
)

// Update overrides current time and updates current note
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

// ParseKeyInput handles key input event from sdl
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
		KeyPos := DrawHelper.GetKeyPos(KeyChar)
		text := fmt.Sprintf("+%d", Point)
		textwidth := DrawHelper.GetTextSize(renderer, DrawHelper.FullFont, text, Constants.BlueThickColor).W()
		KeyPos = pos.FromXY(KeyPos.X()-textwidth/2, KeyPos.Y())

		pointOnKeyboardEffect = Effects.NewAbsoluteFadeout(
			text,
			Constants.BlueThickColor,
			DrawHelper.FullFont,
			KeyPos, 15,
		)
		DrawManager.AddEffector(DrawManager.FOREGROUND, 30, pointOnKeyboardEffect)
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
