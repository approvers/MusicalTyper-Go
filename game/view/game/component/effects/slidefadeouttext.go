package effects

import (
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/view/game/component"
)

// NewSlideFadeoutText makes text renderer with fading out and sliding
func NewSlideFadeoutText(Text string, Color color.Color, FontSize helper.FontSize, Offset pos.Pos, Movement int) component.DrawableEffect {
	return func(ctx *component.EffectDrawContext) {
		Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)

		Color = Color.WithTransparency(Ratio)
		TextSize := helper.GetTextSize(ctx.Renderer, FontSize, Text, Color)
		X := (constants.WindowWidth-TextSize.W())/2 + Offset.X()
		Y := (constants.WindowHeight-TextSize.H())/2 + Offset.Y() - int(float64(Movement)*Ratio)

		helper.DrawText(ctx.Renderer,
			pos.FromXY(X, Y),
			helper.LeftAlign, FontSize,
			Text, Color)
	}
}
