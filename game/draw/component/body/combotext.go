package body

import (
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"
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
			helper.DrawText(Renderer,
				pos.FromXY(Constants.Margin-12, 157),
				helper.LeftAlign, helper.FullFont,
				strconv.Itoa(Combo), Constants.ComboTextColor)

		helper.DrawText(Renderer,
			pos.FromXY(ComboTextWidth+5, 165),
			helper.LeftAlign, helper.SystemFont,
			"chain", comboTextColor())
	}
}
