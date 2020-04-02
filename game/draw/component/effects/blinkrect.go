package effects

import (
	DrawComponent "musicaltyper-go/game/draw/component"

	"github.com/veandco/go-sdl2/sdl"
)

// NewBlinkRect generates effect that draws colored rect renderer with blinking
func NewBlinkRect(Color sdl.Color, Rect *sdl.Rect) DrawComponent.DrawableEffect {
	return func(ctx *DrawComponent.EffectDrawContext) {
		Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)
		Color := Color
		Color.A = uint8(256 - 255*Ratio)

		ctx.Renderer.SetDrawColor(Color.R, Color.G, Color.B, Color.A)
		ctx.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		ctx.Renderer.FillRect(Rect)
		ctx.Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	}
}
