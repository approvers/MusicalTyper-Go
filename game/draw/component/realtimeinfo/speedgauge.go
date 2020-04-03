package realtimeinfo

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

func fastSpeedGaugeAnimateColor() color.Color {
	return Constants.RedColor.Darker(30)
}
func normalSpeedGaugeForegroundColor() color.Color {
	return Constants.GreenThinColor.Darker(50)
}

// SpeedGauge draws players's typing speed by text and color
func SpeedGauge(typingSpeed int, FrameCount int) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		helper.DrawText(Renderer,
			pos.FromXY(Constants.Margin, 382),
			helper.LeftAlign, helper.SystemFont,
			"タイピング速度", Constants.TypedTextColor)

		Area := area.FromXYWH(Constants.Margin, 405, Constants.WindowWidth-Constants.Margin*2, 20)

		if typingSpeed > 4 {
			//4key/secを超えていたら、赤色でアニメーション
			Color := Constants.RedColor
			if !(FrameCount%10 < 5) {
				Color = fastSpeedGaugeAnimateColor()
			}
			helper.DrawFillRect(Renderer, Color, Area)
		} else {
			//そうでなければ普通に描画。
			helper.DrawFillRect(Renderer, Constants.GreenThinColor, Area)

			GaugeWidth := typingSpeed / 4 * (Constants.WindowWidth * 2)
			helper.DrawFillRect(Renderer, normalSpeedGaugeForegroundColor(),
				area.FromXYWH(Constants.Margin, 405,
					int(GaugeWidth), 20))
		}
		Text := fmt.Sprintf("%2d Char/sec", typingSpeed)
		helper.DrawText(Renderer,
			pos.FromXY(Constants.WindowWidth/2, 402),
			helper.Center, helper.SystemFont,
			Text, Constants.TextColor)
	}
}
