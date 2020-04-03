package body

import (
	"musicaltyper-go/game/beatmap"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"

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
		helper.DrawText(Renderer, hiraganaPos(), helper.RightAlign, helper.JapaneseFont, CurrentSentence.GetTypedText(), Constants.TypedTextColor)
		helper.DrawText(Renderer, hiraganaPos(), helper.LeftAlign, helper.JapaneseFont, CurrentSentence.GetRemainingText(), Constants.RemainingTextColor)

		//ローマ字
		helper.DrawText(Renderer, romaPos(), helper.RightAlign, helper.FullFont, CurrentSentence.GetTypedRoma(), Constants.TypedTextColor)
		helper.DrawText(Renderer, romaPos(), helper.LeftAlign, helper.FullFont, CurrentSentence.GetRemainingRoma(), Constants.RemainingTextColor)

		//歌詞
		helper.DrawText(Renderer, lyricPos(), helper.LeftAlign, helper.FullFont, CurrentSentence.OriginalSentence, Constants.LyricTextColor)
	}
}
