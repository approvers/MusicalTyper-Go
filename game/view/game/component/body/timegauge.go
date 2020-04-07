package body

import (
	"math"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/view/game/component"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	foregroundColor = constants.BackgroundColor.Darker(50)
	backgroundColor = constants.BackgroundColor.Darker(25)
)

// TimeGauge draws remaining time gauge
func TimeGauge(normalizedRemainingTime float64) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		RemainingTimeGaugeWidth := int(math.Floor(normalizedRemainingTime * constants.WindowWidth))
		helper.DrawFillRect(Renderer, backgroundColor, area.FromXYWH(0, 60, constants.WindowWidth, 130))
		helper.DrawFillRect(Renderer, foregroundColor, area.FromXYWH(0, 60, RemainingTimeGaugeWidth, 130))
	}
}
