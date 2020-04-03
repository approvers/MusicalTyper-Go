package top

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"
)

// Score draws current score by colored text
func Score(c *DrawComponent.DrawContext) {
	Text := fmt.Sprintf("%08d", c.GameState.Point)
	if c.GameState.Point < 0 {
		ScoreColor := Constants.BlueThickColor
		if c.FrameCount%20 < 10 {
			ScoreColor = Constants.RedColor
		}
		DrawHelper.DrawText(c.Renderer, pos.FromXY(5, 20), DrawHelper.LeftAlign, DrawHelper.AlphabetFont, Text, ScoreColor)
	} else {
		DrawHelper.DrawText(c.Renderer, pos.FromXY(5, 20), DrawHelper.LeftAlign, DrawHelper.AlphabetFont, Text, Constants.BlueThickColor)
	}
}
