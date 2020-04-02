package body

import (
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"
)

type TypeText struct{}

const (
	hiraganaPosX = Constants.WindowWidth / 2
	hiraganaPosY = 80

	romaPosX = Constants.WindowWidth / 2
	romaPosY = 130

	lyricPosX = Constants.Margin - 12
	lyricPosY = 60
)

func (s TypeText) Draw(c *DrawComponent.DrawContext) {
	CurrentNote := c.GameState.Beatmap.Notes[c.GameState.CurrentSentenceIndex]
	CurrentSentence := CurrentNote.Sentence

	//ひらがな
	DrawHelper.DrawText(c.Renderer, hiraganaPosX, hiraganaPosY, DrawHelper.RightAlign, DrawHelper.JapaneseFont, CurrentSentence.GetTypedText(), Constants.TypedTextColor)
	DrawHelper.DrawText(c.Renderer, hiraganaPosX, hiraganaPosY, DrawHelper.LeftAlign, DrawHelper.JapaneseFont, CurrentSentence.GetRebodyingText(), Constants.RebodyingTextColor)

	//ローマ字
	DrawHelper.DrawText(c.Renderer, romaPosX, romaPosY, DrawHelper.RightAlign, DrawHelper.FullFont, CurrentSentence.GetTypedRoma(), Constants.TypedTextColor)
	DrawHelper.DrawText(c.Renderer, romaPosX, romaPosY, DrawHelper.LeftAlign, DrawHelper.FullFont, CurrentSentence.GetRebodyingRoma(), Constants.RebodyingTextColor)

	//歌詞
	DrawHelper.DrawText(c.Renderer, lyricPosX, lyricPosY, DrawHelper.LeftAlign, DrawHelper.FullFont, CurrentSentence.OriginalSentence, Constants.LyricTextColor)
}
