package top

import (
	"fmt"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/view/game/component"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Drawtime draws draw stats.(μs, effect count)

var (
	drawTimeAvg   = make([]float64, constants.FrameRate)
	drawTimeIndex = 0
)

func Drawtime(DrawBeginTime *time.Time, Framecount, FGEffectCount, BGEffectCount int) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		DrawTimeMS := float64(time.Now().Sub(*DrawBeginTime).Microseconds()) / 1000.0
		drawTimeAvg[drawTimeIndex] = DrawTimeMS
		drawTimeIndex++
		if drawTimeIndex == constants.FrameRate {
			drawTimeIndex = 0
		}

		var SumTime float64
		for _, v := range drawTimeAvg {
			SumTime += v
		}

		Text := fmt.Sprintf(
			"%5.2ffps, %3d, %2d, %2d",
			1000/(SumTime/constants.FrameRate),
			Framecount,
			FGEffectCount,
			BGEffectCount,
		)

		helper.DrawTextWithoutCache(Renderer, pos.FromXY(3, -3), helper.LeftAlign, helper.SystemFont, Text, constants.TextColor)
	}
}
