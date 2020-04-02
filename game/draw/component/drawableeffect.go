package component

import (
	"github.com/veandco/go-sdl2/sdl"
)

// DrawableEffect is renderer with EffectDrawContext
type DrawableEffect interface {
	Draw(*EffectDrawContext)
}

// EffectDrawContext is required state to present effect
type EffectDrawContext struct {
	Renderer   *sdl.Renderer
	FrameCount int
	Duration   int
}
