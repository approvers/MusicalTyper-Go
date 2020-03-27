package TopArea

import (
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/DrawComponents"
	"MusicalTyper-Go/Game/DrawHelper"
	"fmt"
)

type Score struct{}

func (s Score) Draw(c *DrawComponents.DrawContext) {
	Text := fmt.Sprintf("%08d", c.GameState.Point)
	if c.GameState.Point < 0 {
		ScoreColor := Constants.BlueThickColor
		if c.FrameCount%20 < 10 {
			ScoreColor = Constants.RedColor
		}
		DrawHelper.DrawText(c.Renderer, 5, 20, DrawHelper.LeftAlign, DrawHelper.AlphabetFont, Text, ScoreColor)
	} else {
		DrawHelper.DrawText(c.Renderer, 5, 20, DrawHelper.LeftAlign, DrawHelper.AlphabetFont, Text, Constants.BlueThickColor)
	}
}
