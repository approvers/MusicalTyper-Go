package MainArea

import (
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/DrawComponents"
	"MusicalTyper-Go/Game/DrawHelper"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type ComboText struct{}

var (
	comboTextColor = &sdl.Color{126, 126, 132, 255}
)

func (s ComboText) Draw(c *DrawComponents.DrawContext) {
	ComboTextWidth, _ :=
		DrawHelper.DrawText(c.Renderer,
			Constants.Margin-12, 157,
			DrawHelper.LeftAlign, DrawHelper.FullFont,
			strconv.Itoa(c.GameState.Combo), Constants.ComboTextColor)

	DrawHelper.DrawText(c.Renderer,
		ComboTextWidth+5, 165,
		DrawHelper.LeftAlign, DrawHelper.SystemFont,
		"chain", comboTextColor)
}
