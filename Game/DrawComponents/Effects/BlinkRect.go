package Effects

import (
	"MusicalTyper-Go/Game/DrawComponents"
	"github.com/veandco/go-sdl2/sdl"
)

type blinkRect struct {
	Color *sdl.Color
	Rect  *sdl.Rect
}

func NewBlinkRect(Color sdl.Color, Rect *sdl.Rect) *blinkRect {
	Result := blinkRect{
		Color: &Color,
		Rect:  Rect,
	}
	return &Result
}

func (Self blinkRect) Draw(ctx *DrawComponents.EffectDrawContext) {
	Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)
	Color := Self.Color
	Color.A = uint8(255 - 255*Ratio)

	ctx.Renderer.SetDrawColor(Color.R, Color.G, Color.B, Color.A)
	ctx.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	ctx.Renderer.FillRect(Self.Rect)
	ctx.Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
}
