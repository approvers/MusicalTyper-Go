package effects

import (
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	DrawComponent "musicaltyper-go/game/draw/component"

	"github.com/veandco/go-sdl2/sdl"
)

// NewBlinkRect generates effect that draws colored rect renderer with blinking
func NewBlinkRect(Color color.Color, Area area.Area) DrawComponent.DrawableEffect {
	return func(ctx *DrawComponent.EffectDrawContext) {
		Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)
		Color := Color
		Color = Color.WithTransparency(Ratio)

		Color.ProxyColor(ctx.Renderer)
		ctx.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		ctx.Renderer.FillRect(Area.ToRect())
		ctx.Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	}
}
