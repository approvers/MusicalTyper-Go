package effects

import (
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"github.com/veandco/go-sdl2/sdl"
)

// NewSlideFadeoutText makes text renderer with fading out and sliding
func NewSlideFadeoutText(Text string, Color sdl.Color, FontSize DrawHelper.FontSize, OffsetX, OffsetY, Movement int) DrawComponent.DrawableEffect {
	return func(ctx *DrawComponent.EffectDrawContext) {
		Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)

		Color.A = uint8(256 - 255*Ratio)
		TextWidth, TextHeight := DrawHelper.GetTextSize(ctx.Renderer, FontSize, Text, &Color)
		X := (Constants.WindowWidth-TextWidth)/2 + OffsetX
		Y := (Constants.WindowHeight-TextHeight)/2 + OffsetY - int(float64(Movement)*Ratio)

		DrawHelper.DrawText(ctx.Renderer,
			X, Y,
			DrawHelper.LeftAlign, FontSize,
			Text, &Color)
	}
}
