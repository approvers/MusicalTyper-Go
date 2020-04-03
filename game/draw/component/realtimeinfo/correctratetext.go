package realtimeinfo

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

func correctRateTextBaseColor() color.Color {
	return Constants.RedColor.Darker(50)
}

// CorrectRateText draws correctness rate by percent text
func CorrectRateText(Acc float64) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		helper.DrawText(Renderer,
			pos.FromXY(Constants.Margin, 430),
			helper.LeftAlign, helper.SystemFont,
			"正解率", Constants.TypedTextColor)

		helper.DrawFillRect(Renderer, correctRateTextBaseColor(),
			area.FromXYWH(Constants.Margin+5, 510,
				int(Acc*250), 3))

		Text := fmt.Sprintf("%05.1f%%", Acc*100)
		TextColor := Constants.RedColor.Multiply(Acc)

		helper.DrawText(Renderer,
			pos.FromXY(Constants.Margin+5, 430),
			helper.LeftAlign, helper.BigFont,
			Text, TextColor)
	}
}
