package mainview

import (
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/view/game/component/effects"
)

/*
Q. Why is this separated between GameState's method?
A. To prevent from cyclic dependencies.
*/

var (
	successEffect = effects.NewSlideFadeoutText(
		"Pass",
		Constants.GreenThinColor.Darker(50),
		helper.AlphabetFont,
		pos.FromXY(-150, -383), 10,
	)

	pointOnKeyboardEffect = effects.NewAbsoluteFadeout(
		"",
		Constants.BlueThickColor,
		helper.FullFont,
		pos.FromXY(0, 0), 15,
	)

	acTextEffect = effects.NewSlideFadeoutText(
		"AC",
		Constants.GreenThickColor,
		helper.AlphabetFont,
		pos.FromXY(-170, -222), 20,
	)

	acBackgroundEffect = effects.NewBlinkRect(
		Constants.GreenThinColor.Brighter(50),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)

	waTextEffect = effects.NewSlideFadeoutText(
		"WA",
		Constants.BlueThickColor.Brighter(100),
		helper.AlphabetFont,
		pos.FromXY(-170, -222), 20,
	)

	waBackgroundEffect = effects.NewBlinkRect(
		Constants.BlueThickColor.Brighter(100),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)

	missTypeTextEffect = effects.NewSlideFadeoutText(
		"MISS",
		Constants.RedColor.Brighter(50),
		helper.AlphabetFont,
		pos.FromXY(-150, -222), 10,
	)

	missTypeBackgroundEffect = effects.NewBlinkRect(
		color.FromRGB(255, 200, 200),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)

	tleTextEffect = effects.NewSlideFadeoutText(
		"TLE",
		Constants.RedColor.Darker(50),
		helper.AlphabetFont,
		pos.FromXY(-150, -222), 10,
	)

	tleBackgroundEffect = effects.NewBlinkRect(
		Constants.RedColor.Brighter(50),
		area.FromXYWH(0, 60, Constants.WindowWidth, 130),
	)
)
