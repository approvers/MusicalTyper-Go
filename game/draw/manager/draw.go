package manager

import (
	DrawComponent "musicaltyper-go/game/draw/component"
	Body "musicaltyper-go/game/draw/component/body"
	Keyboard "musicaltyper-go/game/draw/component/keyboard"
	RealTimeInfo "musicaltyper-go/game/draw/component/realtimeinfo"
	Top "musicaltyper-go/game/draw/component/top"
)

// EffectorPos is kind of effector's position
type EffectorPos uint8

const (
	// FOREGROUND means effector in foreground
	FOREGROUND EffectorPos = iota
	// BACKGROUND means effector in background
	BACKGROUND
)

type effectorEntry struct {
	Drawer     DrawComponent.DrawableEffect
	FrameCount int
	Duration   int
}

var (
	foregroundEffectors = make([]*effectorEntry, 0)
	backgroundEffectors = make([]*effectorEntry, 0)

	backgroundComponents = []DrawComponent.Drawable{
		Top.SongInfo{},
		Top.Score{},
		Body.TimeGauge{},
	}

	foregroundComponents = []DrawComponent.Drawable{
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

// Draw draws components and caches them
func Draw(ctx *DrawComponent.DrawContext) {
	backgroundEffectors = drawComponents(ctx, backgroundComponents, backgroundEffectors)
	foregroundEffectors = drawComponents(ctx, foregroundComponents, foregroundEffectors)
}

// AddEffector adds effector with position and duration
func AddEffector(Pos EffectorPos, Duration int, Effector DrawComponent.DrawableEffect) {
	NewEntry := new(effectorEntry)
	NewEntry.Drawer = Effector
	NewEntry.Duration = Duration
	NewEntry.FrameCount = -1

	switch Pos {
	case FOREGROUND:
		foregroundEffectors = append(foregroundEffectors, NewEntry)

	case BACKGROUND:
		backgroundEffectors = append(backgroundEffectors, NewEntry)
	}
}

func EffectorCount(Pos EffectorPos) int {
	switch Pos {
	case FOREGROUND:
		return len(foregroundEffectors)

	case BACKGROUND:
		return len(backgroundEffectors)

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
