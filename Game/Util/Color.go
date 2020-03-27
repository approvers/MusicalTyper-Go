package Util

import (
	"github.com/veandco/go-sdl2/sdl"
)

func GetMoreWhitishColor(base *sdl.Color, delta uint8) *sdl.Color {
	return &sdl.Color{
		R: MinUInt8(base.R+delta, 255),
		G: MinUInt8(base.G+delta, 255),
		B: MinUInt8(base.B+delta, 255),
		A: base.A,
	}
}

func GetMoreBlackishColor(base *sdl.Color, delta uint8) *sdl.Color {
	return &sdl.Color{
		R: Max(base.R-delta, 0),
		G: Max(base.G-delta, 0),
		B: Max(base.B-delta, 0),
		A: base.A,
	}
}

func GetInvertColor(base *sdl.Color) *sdl.Color {
	return &sdl.Color{
		R: 255 - base.R,
		G: 255 - base.G,
		B: 255 - base.B,
		A: 255 - base.A,
	}
}
