package keyboard

import (
	"fmt"
	"musicaltyper-go/game/beatmap"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

// NextLyrics draws lyrics will be typed
func NextLyrics(isDisabled bool, nextLyrics []*beatmap.Note) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		if isDisabled {
			return
		}
		for i := range nextLyrics {
			if i >= len(nextLyrics) {
				break
			}
			Note := nextLyrics[i]

			helper.DrawText(Renderer, pos.FromXY(5, 193+60*i), helper.LeftAlign, helper.SystemFont, fmt.Sprintf("[%d]", i), Constants.TextColor)
			helper.DrawText(Renderer, pos.FromXY(5, 210+60*i), helper.LeftAlign, helper.FullFont, Note.Sentence.HiraganaSentence, Constants.TextColor)
			helper.DrawText(Renderer, pos.FromXY(5, 230+60*i), helper.LeftAlign, helper.SystemFont, Note.Sentence.GetRoma(), Constants.TextColor)
		}
	}
}
