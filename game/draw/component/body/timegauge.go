package body

import (
	"math"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"github.com/veandco/go-sdl2/sdl"
)

func foregroundColor() color.Color {
	return Constants.BackgroundColor.Darker(50)
}
func backgroundColor() color.Color {
	return Constants.BackgroundColor.Darker(25)
}

// TimeGauge draws remainings time gauge
func TimeGauge(normalizedRemainingTime float64) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		/*var Ratio float64
		if len(c.GameState.Beatmap.Notes) <= c.GameState.CurrentSentenceIndex+1 {
			Ratio = 1
		} else {
			CurrentSentenceStartTime := c.GameState.Beatmap.Notes[c.GameState.CurrentSentenceIndex].Time
			NextSentenceStartTime := c.GameState.Beatmap.Notes[c.GameState.CurrentSentenceIndex+1].Time
			CurrentSentenceDuration := NextSentenceStartTime - CurrentSentenceStartTime
			CurrentTimeInCurrentSentence := CurrentSentenceDuration - c.GameState.CurrentTime + CurrentSentenceStartTime
			Ratio = CurrentTimeInCurrentSentence / CurrentSentenceDuration
		}*/

		RemainingTimeGaugeWidth := int(math.Floor(normalizedRemainingTime * Constants.WindowWidth))
		DrawHelper.DrawFillRect(Renderer, backgroundColor(), area.FromXYWH(0, 60, Constants.WindowWidth, 130))
		DrawHelper.DrawFillRect(Renderer, foregroundColor(), area.FromXYWH(0, 60, RemainingTimeGaugeWidth, 130))
	}
}
