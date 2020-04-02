package RealTimeInfoArea

import (
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/DrawComponent"
	"MusicalTyper-Go/Game/DrawHelper"
	"fmt"
)

type SpeedGauge struct{}

var (
	fastSpeedGaugeAnimateColor      = DrawHelper.GetMoreBlackishColor(Constants.RedColor, 30)
	normalSpeedGaugeForegroundColor = DrawHelper.GetMoreBlackishColor(Constants.GreenThinColor, 50)
)

func (s SpeedGauge) Draw(c *DrawComponent.DrawContext) {
	DrawHelper.DrawText(c.Renderer,
		Constants.Margin, 382,
		DrawHelper.LeftAlign, DrawHelper.SystemFont,
		"タイピング速度", Constants.TypedTextColor)

	if c.GameState.GetKeyTypePerSecond() > 4 {
		//4key/secを超えていたら、赤色でアニメーション
		Color := Constants.RedColor
		if !(c.FrameCount%10 < 5) {
			Color = fastSpeedGaugeAnimateColor
		}
		DrawHelper.DrawFillRect(c.Renderer, Color,
			Constants.Margin, 405,
			Constants.WindowWidth-Constants.Margin*2, 20)
	} else {
		//そうでなければ普通に描画。
		DrawHelper.DrawFillRect(c.Renderer, Constants.GreenThinColor,
			Constants.Margin, 405,
			Constants.WindowWidth-Constants.Margin*2, 20)

		GaugeWidth := c.GameState.GetKeyTypePerSecond() / 4 * (Constants.WindowWidth * 2)
		DrawHelper.DrawFillRect(c.Renderer, normalSpeedGaugeForegroundColor,
			Constants.Margin, 405,
			int(GaugeWidth), 20)
	}
	Text := fmt.Sprintf("%2d Char/sec", c.GameState.GetKeyTypePerSecond())
	DrawHelper.DrawText(c.Renderer,
		Constants.WindowWidth/2, 402,
		DrawHelper.Center, DrawHelper.SystemFont,
		Text, Constants.TextColor)
}
