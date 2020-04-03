package helper

import (
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/pos"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	startY       = 193
	fontSize     = FullFont
	keyMargin    = 5
	keySize      = 40
	keyLineWidth = 2
)

var (
	// KeyboardKeys is set of rows to present to virtual keyboard
	KeyboardKeys = [...]string{"1234567890-\\^", "qwertyuiop@[", "asdfghjkl;:]", "zxcvbnm,./\\"}
)

// DrawKeyboard renders virtual keyboard
func DrawKeyboard(Renderer *sdl.Renderer, HighlightKey string) {
	HighlightKey = strings.ToLower(HighlightKey)
	for Row := 0; Row < 4; Row++ {
		Keys := KeyboardKeys[Row]
		StartPos := (Constants.WindowWidth - ((keySize + keyMargin) * len(Keys))) / 2

		for KeyIndex, Key := range strings.Split(Keys, "") {
			HighlightThisKey := strings.ToLower(HighlightKey) == strings.ToLower(Key)
			RectPosX := StartPos + (keySize+keyMargin)*KeyIndex
			RectPosY := startY + (keySize+keyMargin)*Row
			Area := area.FromXYWH(RectPosX, RectPosY, keySize, keySize)

			if HighlightThisKey {
				DrawFillRect(Renderer, Constants.GreenThickColor, Area)
			}
			DrawLineRect(Renderer, Constants.TextColor, Area, keyLineWidth)

			Color := Constants.TextColor
			if HighlightThisKey {
				Color = Color.Invert()
			} else if Key == "f" || Key == "j" {
				Color = Constants.BlueThickColor
			}

			Key = strings.ToUpper(Key)
			TextSize := GetTextSize(Renderer, fontSize, Key, Color)
			DrawText(Renderer,
				pos.FromXY(StartPos+(keySize+keyMargin)*KeyIndex+keySize/2-TextSize.W()/2,
					startY+(keySize+keyMargin)*Row+keySize/2-TextSize.H()/2),
				LeftAlign, fontSize, Key, Color)
		}
	}
}

// DrawDisabledKeyboard renders disabled virtual keyboard
func DrawDisabledKeyboard(Renderer *sdl.Renderer, HighlightKey string, BackgroundColor color.Color) {
	HighlightKey = strings.ToLower(HighlightKey)
	for Row := 0; Row < 4; Row++ {
		Keys := KeyboardKeys[Row]
		StartPos := (Constants.WindowWidth - ((keySize + keyMargin) * len(Keys))) / 2

		for KeyIndex, Key := range strings.Split(Keys, "") {
			HighlightThisKey := strings.ToLower(HighlightKey) == strings.ToLower(Key)
			RectPosX := StartPos + (keySize+keyMargin)*KeyIndex
			RectPosY := startY + (keySize+keyMargin)*Row
			Area := area.FromXYWH(RectPosX, RectPosY, keySize, keySize)

			if HighlightThisKey {
				DrawFillRect(Renderer, Constants.GreenThickColor, Area)
			} else {
				DrawFillRect(Renderer, BackgroundColor, Area)
			}
			DrawLineRect(Renderer, Constants.TextColor, Area, keyLineWidth)

			Color := Constants.TextColor
			if HighlightThisKey {
				Color = Color.Invert()
			} else if Key == "f" || Key == "j" {
				Color = Constants.BlueThickColor
			}

			Key = strings.ToUpper(Key)
			TextSize := GetTextSize(Renderer, fontSize, Key, Color)
			DrawText(Renderer,
				pos.FromXY(StartPos+(keySize+keyMargin)*KeyIndex+keySize/2-TextSize.W()/2,
					startY+(keySize+keyMargin)*Row+keySize/2-TextSize.H()/2),
				LeftAlign, fontSize, Key, Color)
		}
	}
}

// GetKeyPos calculates position from string of key
func GetKeyPos(key string) pos.Pos {
	Size := keySize + keyMargin
	for i, v := range KeyboardKeys {
		Index := strings.Index(v, key)
		if Index != -1 {
			x := (Constants.WindowWidth-Size*len(v))/2 + Index*Size + keySize/2
			y := startY + i*Size
			return pos.FromXY(x, y)
		}
	}
	return pos.FromXY(0, 0)
}
