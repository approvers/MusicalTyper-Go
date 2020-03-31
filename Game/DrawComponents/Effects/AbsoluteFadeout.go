package Effects

import (
	"MusicalTyper-Go/Game/DrawComponents"
	"MusicalTyper-Go/Game/DrawHelper"
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

func (Self absoluteFadeout) Draw(ctx *DrawComponents.EffectDrawContext) {
	Ratio := float64(ctx.FrameCount) / float64(ctx.Duration)
	Color := Self.Color
	Color.A = uint8(255 * Ratio)

	DrawHelper.DrawText(ctx.Renderer,
		Self.BaseX, Self.BaseY-int(float64(Self.Movement)*Ratio),
		DrawHelper.LeftAlign, Self.FontSize,
		Self.Text,
		Color)
}
