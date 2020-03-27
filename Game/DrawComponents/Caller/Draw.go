package Caller

import (
	"MusicalTyper-Go/Game/DrawComponents"
	"MusicalTyper-Go/Game/DrawComponents/KeyboardArea"
	"MusicalTyper-Go/Game/DrawComponents/MainArea"
	"MusicalTyper-Go/Game/DrawComponents/RealTimeInfoArea"
	"MusicalTyper-Go/Game/DrawComponents/TopArea"
)

var (
	Components = []DrawComponents.Drawable{
		TopArea.SongInfo{},
		TopArea.Score{},
		MainArea.TimeGauge{},
		MainArea.TypeText{},
		MainArea.ComboText{},
		MainArea.AccGauge{},
		MainArea.AchievementGauge{},
		KeyboardArea.KeyboardArea{},
		RealTimeInfoArea.SpeedGauge{},
		RealTimeInfoArea.CorrectRateText{},
		RealTimeInfoArea.AchievementRate{},
	}
)

func Draw(ctx *DrawComponents.DrawContext) {
	for _, v := range Components {
		v.Draw(ctx)
	}
}
