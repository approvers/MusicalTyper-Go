package DrawComponents

import (
	"MusicalTyper-Go/Game/State"
	"github.com/veandco/go-sdl2/sdl"
)

type Drawable interface {
	Draw(*DrawContext)
}

type DrawContext struct {
	Renderer        *sdl.Renderer
	GameState       *Struct.GameState
	PrintNextLyrics bool
	FrameCount      int
}
