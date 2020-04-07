package mainview

import (
	"musicaltyper-go/game/beatmap"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/view"
	"musicaltyper-go/game/view/game/component"
	Body "musicaltyper-go/game/view/game/component/body"
	Keyboard "musicaltyper-go/game/view/game/component/keyboard"
	RealTimeInfo "musicaltyper-go/game/view/game/component/realtimeinfo"
	Top "musicaltyper-go/game/view/game/component/top"
	"time"

	"github.com/veandco/go-sdl2/mix"

	"github.com/veandco/go-sdl2/sdl"
)

type gameView struct {
	frameCount         int
	printingNextLyrics bool
	musicStartTime     *time.Time
	state              *GameState
}

func NewMainView(beatmap *beatmap.Beatmap) view.View {
	Music, _ := mix.LoadMUS(beatmap.Properties["song_data"])
	Music.Play(1)
	MusicStartTime := time.Now()

	result := gameView{
		frameCount:         0,
		printingNextLyrics: false,
		musicStartTime:     &MusicStartTime,
		state:              NewGameState(beatmap),
	}
	return &result
}

func (view *gameView) GetName() string {
	return "GameView"
}

func (view *gameView) PollEvent() view.Event {
	return nil
}

func (view *gameView) HandleSDLEvent(renderer *sdl.Renderer, event sdl.Event) bool {
	switch e := event.(type) {
	case *sdl.KeyboardEvent:
		key := e.Keysym.Sym
		if e.Type == sdl.KEYDOWN {
			switch key {
			case sdl.K_ESCAPE:
				return false

			case sdl.K_LSHIFT, sdl.K_RSHIFT:
				view.printingNextLyrics = !view.printingNextLyrics
				return true

			default:
				view.state.ParseKeyInput(renderer, key, view.printingNextLyrics)
			}
		}
	}

	return true
}

func (view *gameView) Draw(Renderer *sdl.Renderer) {
	Beatmap := view.state.Beatmap

	view.frameCount = (view.frameCount + 1) % constants.FrameRate
	view.state.Update(float64(time.Now().Sub(*view.musicStartTime).Milliseconds()) / 1000.0)

	var (
		NormalizedRemainingTime float64 = 0
		Properties                      = Beatmap.Properties
		CurrentSentenceIndex            = view.state.CurrentSentenceIndex
		CurrentSentence                 = *Beatmap.Notes[CurrentSentenceIndex].Sentence
		NextLyrics                      = view.state.Beatmap.Notes[CurrentSentenceIndex+1 : CurrentSentenceIndex+4]
		FrameCount                      = view.frameCount
		Combo                           = view.state.Combo
		Point                           = view.state.Point
		IsKeyboardDisabled              = view.printingNextLyrics
		Rank                            = view.state.GetRank()
		Accuracy                        = view.state.GetAccuracy()
		TypingSpeed                     = view.state.GetKeyTypePerSecond()
		AchievementRate                 = view.state.GetAchievementRate(false)
		DrawBeginTime                   = time.Now()
	)

	if len(view.state.Beatmap.Notes) <= view.state.CurrentSentenceIndex+1 {
		NormalizedRemainingTime = 1
	} else {
		var (
			CurrentSentenceStartTime = view.state.Beatmap.Notes[CurrentSentenceIndex].Time
			NextSentenceStartTime    = view.state.Beatmap.Notes[CurrentSentenceIndex+1].Time
			CurrentSentenceDuration  = NextSentenceStartTime - CurrentSentenceStartTime
		)
		NormalizedRemainingTime = (CurrentSentenceDuration - view.state.CurrentTime + CurrentSentenceStartTime) / CurrentSentenceDuration
	}

	Renderer.SetDrawColor(255, 243, 224, 0)
	Renderer.Clear()

	backgroundComponents := []component.Drawable{
		Top.SongInfo(Properties),
		Top.Score(Point, FrameCount),
		Body.TimeGauge(NormalizedRemainingTime),
	}
	backgroundEffectors = drawComponents(Renderer, backgroundComponents, backgroundEffectors)

	foregroundComponents := []component.Drawable{
		Body.TypeText(CurrentSentence),
		Body.ComboText(Combo),
		Body.AccGauge(CurrentSentence, AchievementRate, Rank),
		Body.AchievementGauge(AchievementRate),
		Keyboard.Keyboard(IsKeyboardDisabled, CurrentSentence),
		Keyboard.NextLyrics(!IsKeyboardDisabled, NextLyrics),
		RealTimeInfo.SpeedGauge(TypingSpeed, FrameCount),
		RealTimeInfo.CorrectRateText(Accuracy),
		RealTimeInfo.AchievementRate(AchievementRate),
	}
	foregroundEffectors = drawComponents(Renderer, foregroundComponents, foregroundEffectors)

	Top.Drawtime(&DrawBeginTime, FrameCount, len(foregroundEffectors), len(backgroundEffectors))(Renderer)

	Renderer.Present()
}

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
