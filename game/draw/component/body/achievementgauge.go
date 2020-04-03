package body

import (
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/area"

	"github.com/veandco/go-sdl2/sdl"
)

// AchievementGauge draws achivement guage
func AchievementGauge(achievementRate float64) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		//達成率ゲージ
		if achievementRate > 0 {
			DrawHelper.DrawFillRect(Renderer, Constants.RedColor, area.FromXYWH(0, 187, int(Constants.WindowWidth*achievementRate), 3))
		}

		Color := Constants.BlueThickColor
		if achievementRate < 0.8 {
			Color = Constants.GreenThickColor
		}

		DrawHelper.DrawFillRect(Renderer, Color, area.FromXYWH(0, 187, int(Constants.WindowWidth*achievementRate), 3))
	}
}
