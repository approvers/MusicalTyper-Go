package keyboard

import (
	"fmt"
	"musicaltyper-go/game/beatmap"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/view/game/component"

	"github.com/veandco/go-sdl2/sdl"
)

// NextLyrics draws lyrics will be typed
func NextLyrics(isDisabled bool, nextLyrics []*beatmap.Note) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		if isDisabled {
			return
		}
		for i, v := range nextLyrics {
			helper.DrawText(Renderer, pos.FromXY(5, 193+60*i), helper.LeftAlign, helper.SystemFont, fmt.Sprintf("[%d]", i), constants.TextColor)
			helper.DrawText(Renderer, pos.FromXY(5, 210+60*i), helper.LeftAlign, helper.FullFont, v.Sentence.HiraganaSentence, constants.TextColor)
			helper.DrawText(Renderer, pos.FromXY(5, 230+60*i), helper.LeftAlign, helper.SystemFont, v.Sentence.GetRoma(), constants.TextColor)
		}
	}
}
