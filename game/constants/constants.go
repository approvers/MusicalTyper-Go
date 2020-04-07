package constants

import (
	"musicaltyper-go/game/draw/color"
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
	PrintTextureLog     = false
	PrintRomaJudgeCheck = false
	Print

	// AudioChannelNum is the number of will be allocated sound channels.
	AudioChannelNum = 32

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

var (
	// TextColor is any text color
	TextColor = color.FromRGB(56, 56, 62)
	// LyricTextColor is text color of japanese lyric
	LyricTextColor = color.FromRGB(86, 86, 92)
	// ComboTextColor is indication text color on occur combo
	ComboTextColor = color.FromRGB(106, 106, 112)
	// TypedTextColor is typed text color of roman
	TypedTextColor = color.FromRGB(156, 156, 162)
	// RemainingTextColor is text color of roman to be inputted
	RemainingTextColor = TextColor
	// GreenThinColor is light green
	GreenThinColor = color.FromRGB(178, 255, 89)
	// GreenThickColor is muddy green
	GreenThickColor = color.FromRGB(0, 77, 64)
	// BlueThickColor is muddy blue
	BlueThickColor = color.FromRGB(63, 81, 181)
	// RedColor is muddy red
	RedColor = color.FromRGB(250, 119, 109)

	// BackgroundColor is background color of window
	BackgroundColor = color.FromRGBA(255, 243, 224, 0)
)
