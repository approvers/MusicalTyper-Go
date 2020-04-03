package realtimeinfo

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

// AchievementRate draws achievement rate by percent text
func AchievementRate(achievementRate float64) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		helper.DrawText(Renderer,
			pos.FromXY(Constants.Margin+320, 430),
			helper.LeftAlign, helper.SystemFont,
			"達成率", Constants.TypedTextColor)

		Text := fmt.Sprintf("%05.1f%%", achievementRate*100)
		helper.DrawText(Renderer,
			pos.FromXY(Constants.Margin+330, 430),
			helper.LeftAlign, helper.BigFont,
			Text, Constants.BlueThickColor)
	}
}
