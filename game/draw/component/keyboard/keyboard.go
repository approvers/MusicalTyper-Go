package keyboard

import (
	"musicaltyper-go/game/beatmap"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

// Keyboard  draws virtual keyboard
func Keyboard(isDisabled bool, currentSentence beatmap.Sentence) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		if isDisabled {
			return
			// drawDisabledKeyboard(Renderer, "", color.FromRGB(192, 192, 192))
		}
		drawKeyboard(Renderer, DrawHelper.Substring(currentSentence.GetRemainingRoma(), 0, 1))
		//キーボードの下の線
		DrawHelper.DrawThickLine(Renderer,
			pos.FromXY(0, 375), pos.FromXY(Constants.WindowWidth, 375), Constants.TypedTextColor, 2)
	}
}
