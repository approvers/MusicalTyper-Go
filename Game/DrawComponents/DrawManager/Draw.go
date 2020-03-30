package DrawManager

import (
	"MusicalTyper-Go/Game/DrawComponents"
	"MusicalTyper-Go/Game/DrawComponents/KeyboardArea"
	"MusicalTyper-Go/Game/DrawComponents/MainArea"
	"MusicalTyper-Go/Game/DrawComponents/RealTimeInfoArea"
	"MusicalTyper-Go/Game/DrawComponents/TopArea"
)

type EffectorPos uint8

const (
	FOREGROUND EffectorPos = iota
	BACKGROUND
)

type effectorEntry struct {
	Drawer     DrawComponents.DrawableEffect
	FrameCount int
	Duration   int
}

var (
	ForegroundEffectors = make([]*effectorEntry, 0)
	BackgroundEffectors = make([]*effectorEntry, 0)

	BackgroundComponents = []DrawComponents.Drawable{
		TopArea.SongInfo{},
		TopArea.Score{},
		MainArea.TimeGauge{},
	}

	ForegroundComponents = []DrawComponents.Drawable{
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
	BackgroundEffectors = drawEffectors(ctx, BackgroundComponents, BackgroundEffectors)
	ForegroundEffectors = drawEffectors(ctx, ForegroundComponents, ForegroundEffectors)
}

func AddEffector(Pos EffectorPos, Duration int, Effector DrawComponents.DrawableEffect) {
	NewEntry := new(effectorEntry)
	NewEntry.Drawer = Effector
	NewEntry.Duration = Duration

	switch Pos {
	case FOREGROUND:
		ForegroundEffectors = append(ForegroundEffectors, NewEntry)

	case BACKGROUND:
		BackgroundEffectors = append(BackgroundEffectors, NewEntry)
	}
}

//コンポーネントとエフェクトを描画して残ったエフェクトを返す
func drawEffectors(ctx *DrawComponents.DrawContext, components []DrawComponents.Drawable, effectors []*effectorEntry) []*effectorEntry {
	for _, v := range components {
		v.Draw(ctx)
	}

	EffectorContext := new(DrawComponents.EffectDrawContext)
	EffectorContext.Renderer = ctx.Renderer
	EffectorContext.Window = ctx.Window

	RemainEffectors := make([]*effectorEntry, 0, len(effectors))
	for _, v := range effectors {
		EffectorContext.FrameCount = v.FrameCount
		EffectorContext.Duration = v.Duration

		v.Drawer.Draw(EffectorContext)
		if v.FrameCount < v.Duration {
			v.FrameCount++
			RemainEffectors = append(RemainEffectors, v)
		}
	}
	return RemainEffectors
}
