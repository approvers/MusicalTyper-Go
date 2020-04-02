package DrawManager

import (
	"MusicalTyper-Go/Game/DrawComponent"
	"MusicalTyper-Go/Game/DrawComponent/KeyboardArea"
	"MusicalTyper-Go/Game/DrawComponent/MainArea"
	"MusicalTyper-Go/Game/DrawComponent/RealTimeInfoArea"
	"MusicalTyper-Go/Game/DrawComponent/TopArea"
)

type EffectorPos uint8

const (
	FOREGROUND EffectorPos = iota
	BACKGROUND
)

type effectorEntry struct {
	Drawer     DrawComponent.DrawableEffect
	FrameCount int
	Duration   int
}

var (
	ForegroundEffectors = make([]*effectorEntry, 0)
	BackgroundEffectors = make([]*effectorEntry, 0)

	BackgroundComponents = []DrawComponent.Drawable{
		TopArea.SongInfo{},
		TopArea.Score{},
		MainArea.TimeGauge{},
	}

	ForegroundComponents = []DrawComponent.Drawable{
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

func Draw(ctx *DrawComponent.DrawContext) {
	BackgroundEffectors = drawComponents(ctx, BackgroundComponents, BackgroundEffectors)
	ForegroundEffectors = drawComponents(ctx, ForegroundComponents, ForegroundEffectors)
}

func AddEffector(Pos EffectorPos, Duration int, Effector DrawComponent.DrawableEffect) {
	NewEntry := new(effectorEntry)
	NewEntry.Drawer = Effector
	NewEntry.Duration = Duration
	NewEntry.FrameCount = -1

	switch Pos {
	case FOREGROUND:
		ForegroundEffectors = append(ForegroundEffectors, NewEntry)

	case BACKGROUND:
		BackgroundEffectors = append(BackgroundEffectors, NewEntry)
	}
}

func EffectorCount(Pos EffectorPos) int {
	switch Pos {
	case FOREGROUND:
		return len(ForegroundEffectors)

	case BACKGROUND:
		return len(BackgroundEffectors)

	default:
		panic("Unknown effector pos has passed to DrawManager.EffectorCount()")
	}
}

//コンポーネントとエフェクトを描画して残ったエフェクトを返す
func drawComponents(ctx *DrawComponent.DrawContext, components []DrawComponent.Drawable, effectors []*effectorEntry) []*effectorEntry {
	for _, v := range components {
		v.Draw(ctx)
	}

	EffectorContext := new(DrawComponent.EffectDrawContext)
	EffectorContext.Renderer = ctx.Renderer

	RemainEffectors := make([]*effectorEntry, 0, len(effectors))
	for _, v := range effectors {
		v.FrameCount++
		EffectorContext.FrameCount = v.FrameCount
		EffectorContext.Duration = v.Duration

		v.Drawer.Draw(EffectorContext)
		if v.FrameCount < v.Duration {
			RemainEffectors = append(RemainEffectors, v)
		}
	}
	return RemainEffectors
}
