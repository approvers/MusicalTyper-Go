package effects

import (
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/view/game/component"
)

// NewAbsoluteFadeout generates effect that draws text with fading out
func NewAbsoluteFadeout(Text string, Color color.Color, FontSize helper.FontSize, Base pos.Pos, Movement int) component.DrawableEffect {
	return func(ctx *component.EffectDrawContext) {
		Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)
		Color = Color.WithTransparency(Ratio)

		helper.DrawText(ctx.Renderer,
			pos.FromXY(Base.X(), Base.Y()-int(float64(Movement)*Ratio)),
			helper.LeftAlign, FontSize,
			Text,
			Color)
	}
}
