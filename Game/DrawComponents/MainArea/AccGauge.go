package MainArea

import (
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/DrawComponents"
	"MusicalTyper-Go/Game/DrawHelper"
)

type AccGauge struct{}

func (s AccGauge) Draw(c *DrawComponents.DrawContext) {
	CurrentSentence := c.GameState.Beatmap.Notes[c.GameState.CurrentSentenceIndex].Sentence
	RankPosX := int(Constants.WindowWidth * c.GameState.GetAchievementRate(true))

	//正解率ゲージ　100%なら赤色
	if Acc := CurrentSentence.GetAccuracy(); Acc == 1 {
		DrawHelper.DrawFillRect(c.Renderer, Constants.RedColor, 0, 60, int(Acc), 3)
	} else {
		DrawHelper.DrawFillRect(c.Renderer, Constants.GreenThickColor, 0, 60, int(Acc), 3)
	}
	//正解率ゲージの上に出るランク
	DrawHelper.DrawText(c.Renderer,
		RankPosX, 168,
		DrawHelper.RightAlign, DrawHelper.SystemFont,
		Constants.RankTexts[c.GameState.GetRank()], Constants.TypedTextColor)
}
