package DrawComponents

import (
	"MusicalTyper-Go/Game/GameState"
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
