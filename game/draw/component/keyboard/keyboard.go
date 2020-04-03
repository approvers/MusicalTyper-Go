package keyboard

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

// Keyboard presents virtual keyboard area
type Keyboard struct{}

// Draw draws virtual keyboard area
func (s Keyboard) Draw(c *DrawComponent.DrawContext) {
	if c.PrintNextLyrics {
		for i := 0; i < 3; i++ {
			Index := i + c.GameState.CurrentSentenceIndex + 1
			if Index >= len(c.GameState.Beatmap.Notes) {
				break
			}
			Note := c.GameState.Beatmap.Notes[Index]

			DrawHelper.DrawText(c.Renderer, pos.FromXY(5, 193+60*i), DrawHelper.LeftAlign, DrawHelper.SystemFont, fmt.Sprintf("[%d]", Index), Constants.TextColor)
			DrawHelper.DrawText(c.Renderer, pos.FromXY(5, 210+60*i), DrawHelper.LeftAlign, DrawHelper.FullFont, Note.Sentence.HiraganaSentence, Constants.TextColor)
			DrawHelper.DrawText(c.Renderer, pos.FromXY(5, 230+60*i), DrawHelper.LeftAlign, DrawHelper.SystemFont, Note.Sentence.GetRoma(), Constants.TextColor)
		}
	} else {
		if c.GameState.IsInputDisabled {
			DrawHelper.DrawKeyboard(c.Renderer, "", &sdl.Color{R: 192, G: 192, B: 192, A: 255})
		} else {
			CurrentSentence := c.GameState.Beatmap.Notes[c.GameState.CurrentSentenceIndex].Sentence
			DrawHelper.DrawKeyboard(c.Renderer, DrawHelper.Substring(CurrentSentence.GetRemainingRoma(), 0, 1), nil)
		}
	}
	//キーボードの下の線
	DrawHelper.DrawThickLine(c.Renderer, 0, 375, Constants.WindowWidth, 375, Constants.TypedTextColor, 2)
}
