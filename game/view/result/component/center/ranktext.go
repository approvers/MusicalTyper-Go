package center

import (
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/rank"
	"musicaltyper-go/game/view/result/component"

	"github.com/veandco/go-sdl2/sdl"
)

func RankText(r rank.Rank) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		Color := constants.RedColor.Darker(int(150.0 * (r.BorderRate() / 200)))

		helper.DrawText(Renderer, pos.FromXY(constants.Margin, 75), helper.LeftAlign, helper.BigFont, r.Text(), Color)
	}
}
