package body

import (
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"
)

// AchievementGauge presents achivement guage
type AchievementGauge struct{}

// Draw draws achivement guage
func (s AchievementGauge) Draw(c *DrawComponent.DrawContext) {
	//達成率ゲージ
	if Rate := c.GameState.GetAchievementRate(false); Rate > 0 {
		DrawHelper.DrawFillRect(c.Renderer, Constants.RedColor, 0, 187, int(Constants.WindowWidth*Rate), 3)
	}

	Color := Constants.BlueThickColor
	if c.GameState.GetAchievementRate(false) < 0.8 {
		Color = Constants.GreenThickColor
	}

	DrawHelper.DrawFillRect(c.Renderer, Color, 0, 187, int(Constants.WindowWidth*c.GameState.GetAchievementRate(true)), 3)
}
