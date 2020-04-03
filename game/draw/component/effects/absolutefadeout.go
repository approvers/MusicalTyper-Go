package effects

import (
	"musicaltyper-go/game/draw/color"
	DrawComponent "musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
)

// NewAbsoluteFadeout generates effect that draws text with fading out
func NewAbsoluteFadeout(Text string, Color color.Color, FontSize helper.FontSize, Base pos.Pos, Movement int) DrawComponent.DrawableEffect {
	return func(ctx *DrawComponent.EffectDrawContext) {
		Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)
		Color = Color.WithTransparency(Ratio)

		helper.DrawText(ctx.Renderer,
			pos.FromXY(Base.X(), Base.Y()-int(float64(Movement)*Ratio)),
			helper.LeftAlign, FontSize,
			Text,
			Color)
	}
}
