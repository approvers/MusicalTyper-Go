package realtimeinfo

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
)

func correctRateTextBaseColor() color.Color {
	return Constants.RedColor.Darker(50)
}

// CorrectRateText draws correctness rate by percent text
func CorrectRateText(c *DrawComponent.DrawContext) {
	DrawHelper.DrawText(c.Renderer,
		pos.FromXY(Constants.Margin, 430),
		DrawHelper.LeftAlign, DrawHelper.SystemFont,
		"正解率", Constants.TypedTextColor)

	Acc := c.GameState.GetAccuracy()
	DrawHelper.DrawFillRect(c.Renderer, correctRateTextBaseColor(),
		area.FromXYWH(Constants.Margin+5, 510,
			int(Acc*250), 3))

	Text := fmt.Sprintf("%05.1f%%", Acc*100)
	TextColor := Constants.RedColor.Multiply(Acc)

	DrawHelper.DrawText(c.Renderer,
		pos.FromXY(Constants.Margin+5, 430),
		DrawHelper.LeftAlign, DrawHelper.BigFont,
		Text, TextColor)
}
