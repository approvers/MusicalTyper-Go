package body

import (
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"
)

func hiraganaPos() pos.Pos { return pos.FromXY(Constants.WindowWidth/2, 80) }

func romaPos() pos.Pos { return pos.FromXY(Constants.WindowWidth/2, 130) }

func lyricPos() pos.Pos { return pos.FromXY(Constants.Margin-12, 60) }

// TypeText draws hiragana, roman, and japanese lyrics text
func TypeText(c *DrawComponent.DrawContext) {
	CurrentNote := c.GameState.Beatmap.Notes[c.GameState.CurrentSentenceIndex]
	CurrentSentence := CurrentNote.Sentence

	//ひらがな
	DrawHelper.DrawText(c.Renderer, hiraganaPos(), DrawHelper.RightAlign, DrawHelper.JapaneseFont, CurrentSentence.GetTypedText(), Constants.TypedTextColor)
	DrawHelper.DrawText(c.Renderer, hiraganaPos(), DrawHelper.LeftAlign, DrawHelper.JapaneseFont, CurrentSentence.GetRemainingText(), Constants.RemainingTextColor)

	//ローマ字
	DrawHelper.DrawText(c.Renderer, romaPos(), DrawHelper.RightAlign, DrawHelper.FullFont, CurrentSentence.GetTypedRoma(), Constants.TypedTextColor)
	DrawHelper.DrawText(c.Renderer, romaPos(), DrawHelper.LeftAlign, DrawHelper.FullFont, CurrentSentence.GetRemainingRoma(), Constants.RemainingTextColor)

	//歌詞
	DrawHelper.DrawText(c.Renderer, lyricPos(), DrawHelper.LeftAlign, DrawHelper.FullFont, CurrentSentence.OriginalSentence, Constants.LyricTextColor)
}
