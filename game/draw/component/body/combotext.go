package body

import (
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/color"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"strconv"
)

func comboTextColor() color.Color {
	return color.FromRGB(126, 126, 132)
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
