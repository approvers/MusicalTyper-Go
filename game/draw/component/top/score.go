package top

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

// Score draws current score by colored text
func Score(Point int, FrameCount int) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		Text := fmt.Sprintf("%08d", Point)
		if Point < 0 {
			ScoreColor := Constants.BlueThickColor
			if FrameCount%20 < 10 {
				ScoreColor = Constants.RedColor
			}
			helper.DrawText(Renderer, pos.FromXY(5, 20), helper.LeftAlign, helper.AlphabetFont, Text, ScoreColor)
		} else {
			helper.DrawText(Renderer, pos.FromXY(5, 20), helper.LeftAlign, helper.AlphabetFont, Text, Constants.BlueThickColor)
		}
	}
}
