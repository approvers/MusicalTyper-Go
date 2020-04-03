package realtimeinfo

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/pos"
)

// SpeedGauge presents player's typing speed
type SpeedGauge struct{}

var (
	fastSpeedGaugeAnimateColor      = DrawHelper.GetMoreBlackishColor(Constants.RedColor, 30)
	normalSpeedGaugeForegroundColor = DrawHelper.GetMoreBlackishColor(Constants.GreenThinColor, 50)
)

// Draw draws players's typing speed by text and color
func (s SpeedGauge) Draw(c *DrawComponent.DrawContext) {
	DrawHelper.DrawText(c.Renderer,
		pos.FromXY(Constants.Margin, 382),
		DrawHelper.LeftAlign, DrawHelper.SystemFont,
		"タイピング速度", Constants.TypedTextColor)

	Area := area.FromXYWH(Constants.Margin, 405, Constants.WindowWidth-Constants.Margin*2, 20)

	if c.GameState.GetKeyTypePerSecond() > 4 {
		//4key/secを超えていたら、赤色でアニメーション
		Color := Constants.RedColor
		if !(c.FrameCount%10 < 5) {
			Color = fastSpeedGaugeAnimateColor
		}
		DrawHelper.DrawFillRect(c.Renderer, Color, Area)
	} else {
		//そうでなければ普通に描画。
		DrawHelper.DrawFillRect(c.Renderer, Constants.GreenThinColor, Area)

		GaugeWidth := c.GameState.GetKeyTypePerSecond() / 4 * (Constants.WindowWidth * 2)
		DrawHelper.DrawFillRect(c.Renderer, normalSpeedGaugeForegroundColor,
			area.FromXYWH(Constants.Margin, 405,
				int(GaugeWidth), 20))
	}
	Text := fmt.Sprintf("%2d Char/sec", c.GameState.GetKeyTypePerSecond())
	DrawHelper.DrawText(c.Renderer,
		pos.FromXY(Constants.WindowWidth/2, 402),
		DrawHelper.Center, DrawHelper.SystemFont,
		Text, Constants.TextColor)
}
