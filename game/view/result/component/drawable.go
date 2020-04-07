package component

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Drawable is renderer with DrawContext
type Drawable func(*sdl.Renderer)
