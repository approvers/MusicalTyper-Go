package keyboard

import (
	"fmt"
	"musicaltyper-go/game/beatmap"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"
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

			DrawHelper.DrawText(Renderer, pos.FromXY(5, 193+60*i), DrawHelper.LeftAlign, DrawHelper.SystemFont, fmt.Sprintf("[%d]", i), Constants.TextColor)
			DrawHelper.DrawText(Renderer, pos.FromXY(5, 210+60*i), DrawHelper.LeftAlign, DrawHelper.FullFont, Note.Sentence.HiraganaSentence, Constants.TextColor)
			DrawHelper.DrawText(Renderer, pos.FromXY(5, 230+60*i), DrawHelper.LeftAlign, DrawHelper.SystemFont, Note.Sentence.GetRoma(), Constants.TextColor)
		}
	}
}
