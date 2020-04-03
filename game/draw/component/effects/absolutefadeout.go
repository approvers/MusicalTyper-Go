package effects

import (
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

// NewAbsoluteFadeout generates effect that draws text with fading out
func NewAbsoluteFadeout(Text string, Color sdl.Color, FontSize DrawHelper.FontSize, Base pos.Pos, Movement int) DrawComponent.DrawableEffect {
	return func(ctx *DrawComponent.EffectDrawContext) {
		Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)
		Color.A = uint8(256 - 255*Ratio)

		DrawHelper.DrawText(ctx.Renderer,
			pos.FromXY(Base.X(), Base.Y()-int(float64(Movement)*Ratio)),
			DrawHelper.LeftAlign, FontSize,
			Text,
			&Color)
	}
}
