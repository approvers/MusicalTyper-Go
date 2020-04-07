package keyboard

import (
	"musicaltyper-go/game/beatmap"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/view/game/component"

	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

// Keyboard draws virtual keyboard
func Keyboard(isDisabled, isInputDisabled bool, currentSentence beatmap.Sentence) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		if isDisabled {
			return
		}

		if isInputDisabled {
			drawDisabledKeyboard(Renderer, "", color.FromRGB(192, 192, 192))
		} else {
			drawKeyboard(Renderer, helper.Substring(currentSentence.GetRemainingRoma(), 0, 1))
		}
		//キーボードの下の区切り線
		helper.DrawThickLine(Renderer,
			pos.FromXY(0, 375),
			pos.FromXY(constants.WindowWidth, 375),
			constants.TypedTextColor,
			2,
		)
	}
}
