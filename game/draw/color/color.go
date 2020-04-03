package color

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Color expresses color on secreen
type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

// FromRGB makes opaque Color from RGB
func FromRGB(r, g, b uint8) Color {
	return FromRGBA(r, g, b, 255)
}

// FromRGBA makes Color from RGBA
func FromRGBA(r, g, b, a uint8) Color {
	return Color{
		r, g, b, a,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func clamp(a int) uint8 {
	return uint8(max(0, min(255, a)))
}

// WithTransparency makes new Color with transparency 0.0 (transparent) ~ 1.0 (opaque)
func (c Color) WithTransparency(alpha float64) Color {
	return Color{
		r: c.r,
		g: c.g,
		b: c.b,
		a: clamp(int(256 - 255*alpha)),
	}
}

// Multiply makes new Color with multiplied by factor
func (c Color) Multiply(factor float64) Color {
	return Color{
		r: clamp(int(factor * float64(c.r))),
		g: clamp(int(factor * float64(c.g))),
		b: clamp(int(factor * float64(c.b))),
		a: c.a}
}

// Brighter calculates more brighter color
func (c Color) Brighter(delta int) Color {
	return Color{
		r: clamp(int(c.r) + delta),
		g: clamp(int(c.g) + delta),
		b: clamp(int(c.b) + delta),
		a: c.a,
	}
}

// Darker calculates more darker color
func (c Color) Darker(delta int) Color {
	return Color{
		r: clamp(int(c.r) - delta),
		g: clamp(int(c.g) - delta),
		b: clamp(int(c.b) - delta),
		a: c.a,
	}
}

// Invert calculates inverted color
func (c Color) Invert() Color {
	return Color{
		r: 255 - c.r,
		g: 255 - c.g,
		b: 255 - c.b,
		a: 255 - c.a,
	}
}

// Cast casts Color to sdl.Color
func (c Color) Cast() *sdl.Color {
	return &sdl.Color{
		R: c.r,
		G: c.g,
		B: c.b,
		A: c.a,
	}
}

// ProxyColor sets color to sdl.Renderer
func (c Color) ProxyColor(renderer *sdl.Renderer) {
	renderer.SetDrawColor(c.r, c.g, c.b, c.a)
}

func (c Color) String() string {
	return fmt.Sprintf("%d,%d,%d,%d", c.r, c.g, c.b, c.a)
}
