package realtimeinfo

import (
	"fmt"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	correctRateTextBaseColor = constants.RedColor.Darker(50)
)

// CorrectRateText draws correctness rate by percent text
func CorrectRateText(Acc float64) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		helper.DrawText(Renderer,
			pos.FromXY(constants.Margin, 430),
			helper.LeftAlign, helper.SystemFont,
			"正解率", constants.TypedTextColor)

		helper.DrawFillRect(Renderer, correctRateTextBaseColor,
			area.FromXYWH(constants.Margin+5, 510,
				int(Acc*250), 3))

		Text := fmt.Sprintf("%05.1f%%", Acc*100)
		TextColor := constants.RedColor.Multiply(Acc)

		helper.DrawText(Renderer,
			pos.FromXY(constants.Margin+5, 430),
			helper.LeftAlign, helper.BigFont,
			Text, TextColor)
	}
}
