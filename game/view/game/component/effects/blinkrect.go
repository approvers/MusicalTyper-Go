package effects

import (
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/view/game/component"

	"github.com/veandco/go-sdl2/sdl"
)

// NewBlinkRect generates effect that draws colored rect renderer with blinking
func NewBlinkRect(Color color.Color, Area area.Area) component.DrawableEffect {
	return func(ctx *component.EffectDrawContext) {
		Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)
		Color := Color
		Color = Color.WithTransparency(Ratio)

		Color.ApplyColor(ctx.Renderer)
		ctx.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		ctx.Renderer.FillRect(Area.ToRect())
		ctx.Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	}
}
