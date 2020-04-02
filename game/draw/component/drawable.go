package component

import (
	GameState "musicaltyper-go/game/state"
	"github.com/veandco/go-sdl2/sdl"
)

type Drawable interface {
	Draw(*DrawContext)
}

type DrawContext struct {
	Renderer        *sdl.Renderer
	Window          *sdl.Window
	GameState       *GameState.GameState
	PrintNextLyrics bool
	FrameCount      int
}
