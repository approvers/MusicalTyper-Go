package body

import (
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"
	"strconv"

	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

func comboTextColor() *sdl.Color {
	return &sdl.Color{R: 126, G: 126, B: 132, A: 255}
}

// ComboText draws combo indication text
func ComboText(c *DrawComponent.DrawContext) {
	ComboTextWidth, _ :=
		DrawHelper.DrawText(c.Renderer,
			pos.FromXY(Constants.Margin-12, 157),
			DrawHelper.LeftAlign, DrawHelper.FullFont,
			strconv.Itoa(c.GameState.Combo), Constants.ComboTextColor)

	DrawHelper.DrawText(c.Renderer,
		pos.FromXY(ComboTextWidth+5, 165),
		DrawHelper.LeftAlign, DrawHelper.SystemFont,
		"chain", comboTextColor())
}
