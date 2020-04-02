package constants

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// WindowTitle is title of game window
	WindowTitle = "Musical Typer Go"
	// WindowWidth is width of game window
	WindowWidth = 640
	// WindowHeight is height of game window
	WindowHeight = 530
	// FrameRate is update rate of game window frame
	FrameRate = 60

	// Margin is margin of any element on user interface
	Margin = 15

	// PrintTextureLog is flag whether to log about texture
	PrintTextureLog = false
)

var (
	// TextColor is any text color
	TextColor = &sdl.Color{R: 56, G: 56, B: 62, A: 255}
	// LyricTextColor is text color of japanese lyric
	LyricTextColor = &sdl.Color{R: 86, G: 86, B: 92, A: 255}
	// ComboTextColor is indication text color on occur combo
	ComboTextColor = &sdl.Color{R: 106, G: 106, B: 112, A: 255}
	// TypedTextColor is typed text color of roman
	TypedTextColor = &sdl.Color{R: 156, G: 156, B: 162, A: 255}
	// RemainingTextColor is text color of roman to be inputted
	RemainingTextColor = TextColor
	// GreenThinColor is light green
	GreenThinColor = &sdl.Color{R: 178, G: 255, B: 89, A: 255}
	// GreenThickColor is muddy green
	GreenThickColor = &sdl.Color{R: 0, G: 77, B: 64, A: 255}
	// BlueThickColor is muddy blue
	BlueThickColor = &sdl.Color{R: 63, G: 81, B: 181, A: 255}
	// RedColor is muddy red
	RedColor = &sdl.Color{R: 250, G: 119, B: 109, A: 255}

	// BackgroundColor is background color of window
	BackgroundColor = &sdl.Color{R: 255, G: 243, B: 224, A: 0}

	// OneCharPoint is point per typed correct
	OneCharPoint = 10
	// PerfectPoint is extra point when no mistakes
	PerfectPoint = 100
	// SectionPerfectPoint is extra point when no mistakes while the section
	SectionPerfectPoint = 300
	// SpecialPoint has no meanings
	SpecialPoint = 50
	// ClearPoint is extra point when player cleared
	ClearPoint = 50
	// MissPoint is added if user mistook
	MissPoint = -30
	// CouldntTypeCount is point per one character missed to type roman, added when timeout
	CouldntTypeCount = -2
	// IdealTypeSpeed is provisional typing speed used for predication of perfect score
	IdealTypeSpeed = 3
)
