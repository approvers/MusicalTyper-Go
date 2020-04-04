package body

import (
	"musicaltyper-go/game/beatmap"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	hiraganaPos = pos.FromXY(constants.WindowWidth/2, 80)
	romaPos     = pos.FromXY(constants.WindowWidth/2, 130)
	lyricPos    = pos.FromXY(constants.Margin-12, 60)
)

// TypeText draws hiragana, roman, and japanese lyrics text
func TypeText(CurrentSentence beatmap.Sentence) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		//ひらがな
		helper.DrawText(Renderer, hiraganaPos, helper.RightAlign, helper.JapaneseFont, CurrentSentence.GetTypedText(), constants.TypedTextColor)
		helper.DrawText(Renderer, hiraganaPos, helper.LeftAlign, helper.JapaneseFont, CurrentSentence.GetRemainingText(), constants.RemainingTextColor)

		//ローマ字
		helper.DrawText(Renderer, romaPos, helper.RightAlign, helper.FullFont, CurrentSentence.GetTypedRoma(), constants.TypedTextColor)
		helper.DrawText(Renderer, romaPos, helper.LeftAlign, helper.FullFont, CurrentSentence.GetRemainingRoma(), constants.RemainingTextColor)

		//歌詞
		helper.DrawText(Renderer, lyricPos, helper.LeftAlign, helper.FullFont, CurrentSentence.OriginalSentence, constants.LyricTextColor)
	}
}
