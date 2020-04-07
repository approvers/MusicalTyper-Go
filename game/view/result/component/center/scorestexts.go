package center

import (
	"fmt"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/rank"
	"musicaltyper-go/game/view/result/component"

	"github.com/veandco/go-sdl2/sdl"
)

func ScoreText(Point int, Accuracy, AchievementRate float64, r rank.Rank) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		var (
			PointText = fmt.Sprintf("%08d", Point)

			AccuracyText        = fmt.Sprintf("%06.2f", Accuracy*100)
			AccuracyNumberColor = constants.RedColor.Multiply(Accuracy)
			AccuracyTextColor   = constants.TextColor.Brighter(50)

			AchievementRateText  = fmt.Sprintf("%06.2f%%", AchievementRate*100)
			AchievementRateColor = constants.RedColor.Multiply(Accuracy)

			LineColor = constants.TextColor.Brighter(100)

			NextRank, NextRankExists = r.GetNextRank()
		)

		if NextRankExists {
			var (
				NextRankText   = NextRank.Text() + "まで"
				ToNextRankText = fmt.Sprintf("%06.2f%%", NextRank.BorderRate()-(AchievementRate*100))
			)
			helper.DrawText(Renderer, pos.FromXY(constants.Margin+200, 158), helper.LeftAlign, helper.SystemFont, NextRankText, constants.BlueThickColor)
			helper.DrawText(Renderer, pos.FromXY(constants.Margin+200, 168), helper.LeftAlign, helper.AlphabetFont, ToNextRankText, constants.BlueThickColor)
		}

		helper.DrawText(Renderer, pos.FromXY(constants.Margin, 150), helper.LeftAlign, helper.JapaneseFont, AchievementRateText, AchievementRateColor)
		helper.DrawText(Renderer, pos.FromXY(constants.WindowWidth-15, 150), helper.RightAlign, helper.JapaneseFont, PointText, constants.TextColor)

		helper.DrawText(Renderer, pos.FromXY(constants.Margin, 240), helper.LeftAlign, helper.SystemFont, "正解率", AccuracyTextColor)
		helper.DrawText(Renderer, pos.FromXY(constants.Margin+10, 247), helper.LeftAlign, helper.JapaneseFont, AccuracyText, AccuracyNumberColor)

		helper.DrawThickLine(Renderer, pos.FromXY(0, 320), pos.FromXY(constants.WindowWidth, 320), LineColor, 2)
	}
}
