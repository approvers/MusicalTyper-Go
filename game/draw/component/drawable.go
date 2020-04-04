package component

import (
	"github.com/veandco/go-sdl2/sdl"
	"musicaltyper-go/game/beatmap"
	"musicaltyper-go/game/rank"
	"time"
)

// Drawable is renderer with DrawContext
type Drawable func(*sdl.Renderer)

// DrawContext is whole of state to present
type DrawContext struct {
	Renderer                *sdl.Renderer
	Properties              map[string]string
	CurrentSentence         beatmap.Sentence
	Rank                    rank.Rank
	NormalizedRemainingTime float64
	AchievementRate         float64
	Point                   int
	Combo                   int
	NextLyrics              []*beatmap.Note
	Accuracy                float64
	TypingSpeed             int
	IsKeyboardDisabled      bool
	FrameCount              int
	DrawBeginTime           *time.Time
}
