package bottom

import (
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/view/result/component"

	"github.com/veandco/go-sdl2/sdl"
)

func KeyText() component.Drawable {
	return func(Renderer *sdl.Renderer) {
		helper.DrawText(Renderer, pos.FromXY(constants.Margin-10, 320), helper.LeftAlign, helper.AlphabetFont, "[R]/リトライ", constants.TextColor)
		helper.DrawText(Renderer, pos.FromXY(constants.Margin+300, 320), helper.LeftAlign, helper.AlphabetFont, "[Esc]/終了", constants.TextColor)
	}
}
