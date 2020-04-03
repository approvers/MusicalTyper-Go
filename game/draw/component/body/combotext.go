package body

import (
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

func comboTextColor() color.Color {
	return color.FromRGB(126, 126, 132)
}

// ComboText draws combo indication text
func ComboText(Combo int) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		ComboTextWidth, _ :=
			DrawHelper.DrawText(Renderer,
				pos.FromXY(Constants.Margin-12, 157),
				DrawHelper.LeftAlign, DrawHelper.FullFont,
				strconv.Itoa(Combo), Constants.ComboTextColor)

		DrawHelper.DrawText(Renderer,
			pos.FromXY(ComboTextWidth+5, 165),
			DrawHelper.LeftAlign, DrawHelper.SystemFont,
			"chain", comboTextColor())
	}
}
