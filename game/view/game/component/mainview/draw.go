package mainview

import (
	"musicaltyper-go/game/view/game/component"
	Body "musicaltyper-go/game/view/game/component/body"
	Keyboard "musicaltyper-go/game/view/game/component/keyboard"
	RealTimeInfo "musicaltyper-go/game/view/game/component/realtimeinfo"
	Top "musicaltyper-go/game/view/game/component/top"

	"github.com/veandco/go-sdl2/sdl"
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
	Drawer     component.DrawableEffect
	FrameCount int
	Duration   int
}

var (
	foregroundEffectors = make([]*effectorEntry, 0)
	backgroundEffectors = make([]*effectorEntry, 0)
)

// Draw draws components and caches them
func Draw(ctx *component.DrawContext) {
	backgroundComponents := []component.Drawable{
		Top.SongInfo(ctx.Properties),
		Top.Score(ctx.Point, ctx.FrameCount),
		Body.TimeGauge(ctx.NormalizedRemainingTime),
	}
	backgroundEffectors = drawComponents(ctx.Renderer, backgroundComponents, backgroundEffectors)

	foregroundComponents := []component.Drawable{
		Body.TypeText(ctx.CurrentSentence),
		Body.ComboText(ctx.Combo),
		Body.AccGauge(ctx.CurrentSentence, ctx.AchievementRate, ctx.Rank),
		Body.AchievementGauge(ctx.AchievementRate),
		Keyboard.Keyboard(ctx.IsKeyboardDisabled, ctx.CurrentSentence),
		Keyboard.NextLyrics(!ctx.IsKeyboardDisabled, ctx.NextLyrics),
		RealTimeInfo.SpeedGauge(ctx.TypingSpeed, ctx.FrameCount),
		RealTimeInfo.CorrectRateText(ctx.Accuracy),
		RealTimeInfo.AchievementRate(ctx.AchievementRate),
	}
	foregroundEffectors = drawComponents(ctx.Renderer, foregroundComponents, foregroundEffectors)
}

// AddEffector adds effector with position and duration
func AddEffector(Pos EffectorPos, Duration int, Effector component.DrawableEffect) {
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

func effectorCount(Pos EffectorPos) int {
	switch Pos {
	case FOREGROUND:
		return len(foregroundEffectors)

	case BACKGROUND:
		return len(backgroundEffectors)

	default:
		panic("Unknown effector pos has passed to DrawManager.effectorCount()")
	}
}

//コンポーネントとエフェクトを描画して残ったエフェクトを返す
func drawComponents(renderer *sdl.Renderer, components []component.Drawable, effectors []*effectorEntry) []*effectorEntry {
	for _, v := range components {
		v(renderer)
	}

	EffectorContext := new(component.EffectDrawContext)
	EffectorContext.Renderer = renderer

	RemainEffectors := make([]*effectorEntry, 0, len(effectors))
	for _, v := range effectors {
		v.FrameCount++
		EffectorContext.FrameCount = v.FrameCount
		EffectorContext.Duration = v.Duration

		v.Drawer(EffectorContext)
		if v.FrameCount < v.Duration {
			RemainEffectors = append(RemainEffectors, v)
		}
	}
	return RemainEffectors
}
