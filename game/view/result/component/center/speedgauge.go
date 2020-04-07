package center

import (
	"fmt"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/view/result/component"

	"github.com/veandco/go-sdl2/sdl"
)

func SpeedGauge(typespeed float64) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		if typespeed > 4 {
			helper.DrawFillRect(Renderer, constants.RedColor.Darker(30), area.FromXYWH(
				constants.Margin, 210,
				constants.WindowWidth-constants.Margin*2, 20,
			))
		} else {
			helper.DrawFillRect(Renderer, constants.GreenThinColor, area.FromXYWH(
				constants.Margin, 210,
				constants.WindowWidth-constants.Margin*2, 20,
			))

			helper.DrawFillRect(Renderer, constants.GreenThinColor.Darker(50), area.FromXYWH(
				constants.Margin, 210,
				int(typespeed/4.0*float64(constants.WindowWidth-constants.Margin*2)), 20,
			))
		}
		Text := fmt.Sprintf("%4.2f Char/sec", typespeed)
		helper.DrawText(Renderer, pos.FromXY(constants.WindowWidth/2, 208), helper.Center, helper.SystemFont, Text, constants.TextColor)
	}
}
