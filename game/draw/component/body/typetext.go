package body

import (
	"musicaltyper-go/game/beatmap"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

func hiraganaPos() pos.Pos { return pos.FromXY(Constants.WindowWidth/2, 80) }

func romaPos() pos.Pos { return pos.FromXY(Constants.WindowWidth/2, 130) }

func lyricPos() pos.Pos { return pos.FromXY(Constants.Margin-12, 60) }

// TypeText draws hiragana, roman, and japanese lyrics text
func TypeText(CurrentSentence beatmap.Sentence) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		//ひらがな
		DrawHelper.DrawText(Renderer, hiraganaPos(), DrawHelper.RightAlign, DrawHelper.JapaneseFont, CurrentSentence.GetTypedText(), Constants.TypedTextColor)
		DrawHelper.DrawText(Renderer, hiraganaPos(), DrawHelper.LeftAlign, DrawHelper.JapaneseFont, CurrentSentence.GetRemainingText(), Constants.RemainingTextColor)

		//ローマ字
		DrawHelper.DrawText(Renderer, romaPos(), DrawHelper.RightAlign, DrawHelper.FullFont, CurrentSentence.GetTypedRoma(), Constants.TypedTextColor)
		DrawHelper.DrawText(Renderer, romaPos(), DrawHelper.LeftAlign, DrawHelper.FullFont, CurrentSentence.GetRemainingRoma(), Constants.RemainingTextColor)

		//歌詞
		DrawHelper.DrawText(Renderer, lyricPos(), DrawHelper.LeftAlign, DrawHelper.FullFont, CurrentSentence.OriginalSentence, Constants.LyricTextColor)
	}
}
