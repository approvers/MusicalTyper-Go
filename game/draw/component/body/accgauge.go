package body

import (
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/pos"
)

// AccGauge draws accuracy guage and player rank
func AccGauge(c *DrawComponent.DrawContext) {
	CurrentSentence := c.GameState.Beatmap.Notes[c.GameState.CurrentSentenceIndex].Sentence
	RankPosX := int(Constants.WindowWidth * c.GameState.GetAchievementRate(true))

	//正解率ゲージ　100%なら赤色
	Acc := CurrentSentence.GetAccuracy()
	GaugeArea := area.FromXYWH(0, 60, int(Acc), 3)
	if Acc == 1 {
		DrawHelper.DrawFillRect(c.Renderer, Constants.RedColor, GaugeArea)
	} else {
		DrawHelper.DrawFillRect(c.Renderer, Constants.GreenThickColor, GaugeArea)
	}
	//正解率ゲージの上に出るランク
	DrawHelper.DrawText(c.Renderer,
		pos.FromXY(RankPosX, 168),
		DrawHelper.RightAlign, DrawHelper.SystemFont,
		c.GameState.GetRank().Text(), Constants.TypedTextColor)
}
