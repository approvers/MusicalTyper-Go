package body

import (
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/area"

	"github.com/veandco/go-sdl2/sdl"
)

// AchievementGauge draws achievement gauge
func AchievementGauge(achievementRate float64) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		//達成率ゲージ
		if achievementRate > 0 {
			helper.DrawFillRect(Renderer, constants.RedColor, area.FromXYWH(0, 187, int(constants.WindowWidth*achievementRate), 3))
		}

		Color := constants.BlueThickColor
		if achievementRate < 0.8 {
			Color = constants.GreenThickColor
		}

		helper.DrawFillRect(Renderer, Color, area.FromXYWH(0, 187, int(constants.WindowWidth*achievementRate), 3))
	}
}
