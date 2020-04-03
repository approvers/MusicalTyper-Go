package component

import (
	GameState "musicaltyper-go/game/state"

	"github.com/veandco/go-sdl2/sdl"
)

// Drawable is renderer with DrawContext
type Drawable func(*DrawContext)

// DrawContext is whole of state to present
type DrawContext struct {
	Renderer        *sdl.Renderer
	Window          *sdl.Window
	GameState       *GameState.GameState
	PrintNextLyrics bool
	FrameCount      int
}
