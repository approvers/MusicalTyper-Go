package manager

import (
	DrawComponent "musicaltyper-go/game/draw/component"
	Body "musicaltyper-go/game/draw/component/body"
	Keyboard "musicaltyper-go/game/draw/component/keyboard"
	RealTimeInfo "musicaltyper-go/game/draw/component/realtimeinfo"
	Top "musicaltyper-go/game/draw/component/top"
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
		Top.SongInfo{},
		Top.Score{},
		Body.TimeGauge{},
	}

	ForegroundComponents = []DrawComponent.Drawable{
		Body.TypeText{},
		Body.ComboText{},
		Body.AccGauge{},
		Body.AchievementGauge{},
		Keyboard.Keyboard{},
		RealTimeInfo.SpeedGauge{},
		RealTimeInfo.CorrectRateText{},
		RealTimeInfo.AchievementRate{},
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
