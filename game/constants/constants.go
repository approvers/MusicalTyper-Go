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
	TextColor = &sdl.Color{56, 56, 62, 255}
	// LyricTextColor is text color of japanese lyric
	LyricTextColor = &sdl.Color{86, 86, 92, 255}
	// ComboTextColor is indication text color on occur combo
	ComboTextColor = &sdl.Color{106, 106, 112, 255}
	// TypedTextColor is typed text color of roman
	TypedTextColor = &sdl.Color{156, 156, 162, 255}
	// RemainingTextColor is text color of roman to be inputted
	RemainingTextColor = TextColor
	// GreenThinColor is light green
	GreenThinColor = &sdl.Color{178, 255, 89, 255}
	// GreenThickColor is muddy green
	GreenThickColor = &sdl.Color{0, 77, 64, 255}
	// BlueThickColor is muddy blue
	BlueThickColor = &sdl.Color{63, 81, 181, 255}
	// RedColor is muddy red
	RedColor = &sdl.Color{250, 119, 109, 255}

	// BackgroundColor is background color of window
	BackgroundColor = &sdl.Color{255, 243, 224, 0}

	// RankPoints are waypoints to decide rank
	RankPoints = [...]float64{200, 150, 125, 100, 99.50, 99, 98, 97, 94, 90, 80, 60, 40, 20, 10, 0}
	// RankTexts are expressions of rank
	RankTexts = [...]string{"Wow", "Unexpected", "Very God", "God", "Pro", "Genius", "Geki-tsuyo", "tsuyotusyo", "AAA", "AA", "A", "B", "C", "D", "E", "F"}

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
