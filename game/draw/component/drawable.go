package component

import (
	"github.com/veandco/go-sdl2/sdl"
	GameState "musicaltyper-go/game/state"
)

// Drawable is renderer with DrawContext
type Drawable interface {
	Draw(*DrawContext)
}

// DrawContext is whole of state to present
type DrawContext struct {
	Renderer        *sdl.Renderer
	Window          *sdl.Window
	GameState       *GameState.GameState
	PrintNextLyrics bool
	FrameCount      int
}
