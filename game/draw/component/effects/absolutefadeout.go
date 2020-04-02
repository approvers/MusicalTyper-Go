package effects

import (
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"github.com/veandco/go-sdl2/sdl"
)

type absoluteFadeout struct {
	Text     string
	Color    *sdl.Color
	FontSize DrawHelper.FontSize

	BaseX    int
	BaseY    int
	Movement int
}

// NewAbsoluteFadeout makes text renderer with fading out
func NewAbsoluteFadeout(Text string, Color sdl.Color, Size DrawHelper.FontSize, BaseX, BaseY, Movement int) *absoluteFadeout {
	Result := absoluteFadeout{
		Text:     Text,
		Color:    &Color,
		FontSize: Size,
		BaseX:    BaseX,
		BaseY:    BaseY,
		Movement: Movement,
	}
	return &Result
}

// Draw draws text with fading out
func (Self absoluteFadeout) Draw(ctx *DrawComponent.EffectDrawContext) {
	Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)
	Color := Self.Color
	Color.A = uint8(256 - 255*Ratio)

	DrawHelper.DrawText(ctx.Renderer,
		Self.BaseX, Self.BaseY-int(float64(Self.Movement)*Ratio),
		DrawHelper.LeftAlign, Self.FontSize,
		Self.Text,
		Color)
}
