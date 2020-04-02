package effects

import (
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"github.com/veandco/go-sdl2/sdl"
)

type slideFadeoutText struct {
	Text     string
	Color    *sdl.Color
	FontSize DrawHelper.FontSize

	OffsetX  int
	OffsetY  int
	Movement int
}

// NewSlideFadeoutText makes text renderer with fading out and sliding
func NewSlideFadeoutText(Text string, Color sdl.Color, Size DrawHelper.FontSize, OffsetX, OffsetY, Movement int) *slideFadeoutText {
	Result := slideFadeoutText{
		Text:     Text,
		Color:    &Color,
		FontSize: Size,
		OffsetX:  OffsetX,
		OffsetY:  OffsetY,
		Movement: Movement,
	}
	return &Result
}

// Draw draws text with fading out and sliding
func (Self slideFadeoutText) Draw(ctx *DrawComponent.EffectDrawContext) {
	Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)

	Color := Self.Color
	Color.A = uint8(256 - 255*Ratio)
	TextWidth, TextHeight := DrawHelper.GetTextSize(ctx.Renderer, Self.FontSize, Self.Text, Self.Color)
	X := (Constants.WindowWidth-TextWidth)/2 + Self.OffsetX
	Y := (Constants.WindowHeight-TextHeight)/2 + Self.OffsetY - int(float64(Self.Movement)*Ratio)

	DrawHelper.DrawText(ctx.Renderer,
		X, Y,
		DrawHelper.LeftAlign, Self.FontSize,
		Self.Text, Color)
}
