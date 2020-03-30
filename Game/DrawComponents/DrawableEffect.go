package DrawComponents

import (
	"github.com/veandco/go-sdl2/sdl"
)

type DrawableEffect interface {
	Draw(*EffectDrawContext)
}

type EffectDrawContext struct {
	Renderer   *sdl.Renderer
	Window     *sdl.Window
	FrameCount int
	Duration   int
}
