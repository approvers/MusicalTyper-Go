package effects

import (
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/color"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
)

// NewSlideFadeoutText makes text renderer with fading out and sliding
func NewSlideFadeoutText(Text string, Color color.Color, FontSize DrawHelper.FontSize, Offset pos.Pos, Movement int) DrawComponent.DrawableEffect {
	return func(ctx *DrawComponent.EffectDrawContext) {
		Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)

		Color = Color.WithTransparency(Ratio)
		TextSize := DrawHelper.GetTextSize(ctx.Renderer, FontSize, Text, Color)
		X := (Constants.WindowWidth-TextSize.W())/2 + Offset.X()
		Y := (Constants.WindowHeight-TextSize.H())/2 + Offset.Y() - int(float64(Movement)*Ratio)

		DrawHelper.DrawText(ctx.Renderer,
			pos.FromXY(X, Y),
			DrawHelper.LeftAlign, FontSize,
			Text, Color)
	}
}
