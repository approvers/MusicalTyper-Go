package DrawHelper

import (
	"MusicalTyper-Go/Game/Constants"
	"github.com/veandco/go-sdl2/sdl"
	"strings"
)

const (
	startY       = 193
	fontSize     = FullFont
	keyMargin    = 5
	keySize      = 40
	keyLineWidth = 2 //fixme: name?
)

var (
	KeyboardKeys = [...]string{"1234567890-\\^", "qwertyuiop@[", "asdfghjkl;:]", "zxcvbnm,./\\"}
)

func DrawKeyboard(Renderer *sdl.Renderer, HighlightKey string, BackgroundColor *sdl.Color) {
	HighlightKey = strings.ToLower(HighlightKey)
	for Row := 0; Row < 4; Row++ {
		Keys := KeyboardKeys[Row]
		StartPos := (Constants.WindowWidth - ((keySize + keyMargin) * len(Keys))) / 2

		for KeyIndex, Key := range strings.Split(Keys, "") {
			HighlightThisKey := strings.ToLower(HighlightKey) == strings.ToLower(Key)
			RectPosX := StartPos + (keySize+keyMargin)*KeyIndex
			RectPosY := startY + (keySize+keyMargin)*Row

			if HighlightThisKey {
				DrawFillRect(Renderer, Constants.GreenThickColor, RectPosX, RectPosY, keySize, keySize)
			} else if BackgroundColor != nil {
				DrawFillRect(Renderer, BackgroundColor, RectPosX, RectPosY, keySize, keySize)
			}
			DrawLineRect(Renderer, Constants.TextColor, RectPosX, RectPosY, keySize, keySize, keyLineWidth)

			Color := Constants.TextColor
			if HighlightThisKey {
				Color = GetInvertColor(Color)
			} else if Key == "f" || Key == "j" {
				Color = Constants.BlueThickColor
			}

			Key = strings.ToUpper(Key)
			TextWidth, TextHeight := GetTextSize(Renderer, fontSize, Key, Color)
			DrawText(Renderer, StartPos+(keySize+keyMargin)*KeyIndex+keySize/2-TextWidth/2, startY+(keySize+keyMargin)*Row+keySize/2-TextHeight/2, LeftAlign, fontSize, Key, Color)
		}
	}
}
