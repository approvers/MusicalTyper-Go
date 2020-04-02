package DrawComponent

import (
	"github.com/veandco/go-sdl2/sdl"
)

type DrawableEffect interface {
	Draw(*EffectDrawContext)
}

type EffectDrawContext struct {
	Renderer   *sdl.Renderer
	FrameCount int
	Duration   int
}
