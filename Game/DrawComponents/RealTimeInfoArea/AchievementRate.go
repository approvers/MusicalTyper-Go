package RealTimeInfoArea

import (
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/DrawComponents"
	"MusicalTyper-Go/Game/DrawHelper"
	"fmt"
)

type AchievementRate struct{}

func (s AchievementRate) Draw(c *DrawComponents.DrawContext) {
	DrawHelper.DrawText(c.Renderer,
		Constants.Margin+320, 430,
		DrawHelper.LeftAlign, DrawHelper.SystemFont,
		"達成率", Constants.TypedTextColor)

	Text := fmt.Sprintf("%05.1f%%", c.GameState.GetAchievementRate(false))
	DrawHelper.DrawText(c.Renderer,
		Constants.Margin+330, 430,
		DrawHelper.LeftAlign, DrawHelper.BigFont,
		Text, Constants.BlueThickColor)
}
