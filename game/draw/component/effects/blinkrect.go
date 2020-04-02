package effects

import (
	DrawComponent "musicaltyper-go/game/draw/component"

	"github.com/veandco/go-sdl2/sdl"
)

type blinkRect struct {
	Color *sdl.Color
	Rect  *sdl.Rect
}

// NewBlinkRect makes colored rect renderer with blinking
func NewBlinkRect(Color sdl.Color, Rect *sdl.Rect) *blinkRect {
	Result := blinkRect{
		Color: &Color,
		Rect:  Rect,
	}
	return &Result
}

// Draw draws colored rect with blinking
func (Self blinkRect) Draw(ctx *DrawComponent.EffectDrawContext) {
	Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)
	Color := Self.Color
	Color.A = uint8(256 - 255*Ratio)

	ctx.Renderer.SetDrawColor(Color.R, Color.G, Color.B, Color.A)
	ctx.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	ctx.Renderer.FillRect(Self.Rect)
	ctx.Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
}
