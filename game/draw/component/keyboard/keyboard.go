package KeyboardArea

import (
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/DrawComponent"
	"MusicalTyper-Go/Game/DrawHelper"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardArea struct{}

func (s KeyboardArea) Draw(c *DrawComponent.DrawContext) {
	if c.PrintNextLyrics {
		for i := 0; i < 3; i++ {
			Index := i + c.GameState.CurrentSentenceIndex + 1
			if Index >= len(c.GameState.Beatmap.Notes) {
				break
			}
			Note := c.GameState.Beatmap.Notes[Index]

			DrawHelper.DrawText(c.Renderer, 5, 193+60*i, DrawHelper.LeftAlign, DrawHelper.SystemFont, fmt.Sprintf("[%d]", Index), Constants.TextColor)
			DrawHelper.DrawText(c.Renderer, 5, 210+60*i, DrawHelper.LeftAlign, DrawHelper.FullFont, Note.Sentence.HiraganaSentence, Constants.TextColor)
			DrawHelper.DrawText(c.Renderer, 5, 230+60*i, DrawHelper.LeftAlign, DrawHelper.SystemFont, Note.Sentence.GetRoma(), Constants.TextColor)
		}
	} else {
		if c.GameState.IsInputDisabled {
			DrawHelper.DrawKeyboard(c.Renderer, "", &sdl.Color{192, 192, 192, 255})
		} else {
			CurrentSentence := c.GameState.Beatmap.Notes[c.GameState.CurrentSentenceIndex].Sentence
			DrawHelper.DrawKeyboard(c.Renderer, DrawHelper.Substring(CurrentSentence.GetRemainingRoma(), 0, 1), nil)
		}
	}
	//キーボードの下の線
	DrawHelper.DrawThickLine(c.Renderer, 0, 375, Constants.WindowWidth, 375, Constants.TypedTextColor, 2)
}
