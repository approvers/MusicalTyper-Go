package realtimeinfo

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"github.com/veandco/go-sdl2/sdl"
)

// CorrectRateText presents correctness rate
type CorrectRateText struct{}

var (
	correctRateTextBaseColor = DrawHelper.GetMoreBlackishColor(Constants.RedColor, 50)
)

// Draw draws correctness rate by percent text
func (s CorrectRateText) Draw(c *DrawComponent.DrawContext) {
	DrawHelper.DrawText(c.Renderer,
		Constants.Margin, 430,
		DrawHelper.LeftAlign, DrawHelper.SystemFont,
		"正解率", Constants.TypedTextColor)

	Acc := c.GameState.GetAccuracy()
	DrawHelper.DrawFillRect(c.Renderer, correctRateTextBaseColor,
		Constants.Margin+5, 510,
		int(Acc*250), 3)

	Text := fmt.Sprintf("%05.1f%%", Acc*100)
	TextColor := &sdl.Color{
		R: uint8(Acc * float64(Constants.RedColor.R)),
		G: uint8(Acc * float64(Constants.RedColor.G)),
		B: uint8(Acc * float64(Constants.RedColor.B)),
		A: 255}

	DrawHelper.DrawText(c.Renderer,
		Constants.Margin+5, 430,
		DrawHelper.LeftAlign, DrawHelper.BigFont,
		Text, TextColor)
}
