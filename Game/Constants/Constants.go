package Constants

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowTitle  = "Musical Typer Go"
	WindowWidth  = 640
	WindowHeight = 530
	FrameRate    = 60

	Margin = 15

	IsFontCacheEnabled = true
)

var (
	TextColor           = &sdl.Color{56, 56, 62, 255}
	LyricTextColor      = &sdl.Color{86, 86, 92, 255}
	ComboTextColor      = &sdl.Color{106, 106, 112, 255}
	ComboChainTextColor = &sdl.Color{126, 126, 132, 255}
	TypedTextColor      = &sdl.Color{156, 156, 162, 255}
	RemainingTextColor  = TextColor
	GreenThickColor     = &sdl.Color{0, 77, 64, 255}
	BlueThickColor      = &sdl.Color{63, 81, 181, 255}
	RedColor            = &sdl.Color{250, 119, 109, 255}

	BackgroundColor = &sdl.Color{255, 243, 224, 0}

	RankPoints = []float64{200, 150, 125, 100, 99.50, 99, 98, 97, 94, 90, 80, 60, 40, 20, 10, 0}
	RankTexts  = []string{"Wow", "Unexpected", "Very God", "God", "Pro", "Genius", "Geki-tsuyo", "tsuyotusyo", "AAA", "AA", "A", "B", "C", "D", "E", "F"}
)
