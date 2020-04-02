package MainArea

import (
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/DrawComponent"
	"MusicalTyper-Go/Game/DrawHelper"
)

type AchievementGauge struct{}

func (s AchievementGauge) Draw(c *DrawComponent.DrawContext) {
	//達成率ゲージ
	if GotRank := c.GameState.GetRank(); GotRank > 0 {
		Rate := Constants.RankPoints[GotRank-1] / 100
		DrawHelper.DrawFillRect(c.Renderer, Constants.RedColor, 0, 187, int(Constants.WindowWidth*Rate), 3)
	}

	Color := Constants.BlueThickColor
	if c.GameState.GetAchievementRate(false) < 0.8 {
		Color = Constants.GreenThickColor
	}

	DrawHelper.DrawFillRect(c.Renderer, Color, 0, 187, int(Constants.WindowWidth*c.GameState.GetAchievementRate(true)), 3)
}