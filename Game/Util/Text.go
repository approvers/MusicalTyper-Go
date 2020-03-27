package Util

import (
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/Logger"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"time"
)

type FontSize uint8
type TextureCache struct {
	Width, Height int32
	Texture       *sdl.Texture
}

var (
	fontCache    = map[FontSize]*ttf.Font{}
	textureCache = map[FontSize]map[string]TextureCache{}
)

const (
	BigFont      FontSize = 72
	JapaneseFont          = 48
	AlphabetFont          = 32
	FullFont              = 24
	RankFont              = 20
	SystemFont            = 16
)

func DrawText(Renderer *sdl.Renderer, x, y int, alignment AlignmentType, Size FontSize, Text string, Color *sdl.Color) (int, int) {
	logger := Logger.NewLogger("DrawText")
	if Text == "" {
		return 0, 0
	}

	Texture, TextureExists := textureCache[Size][Text]
	if !TextureExists {
		Begin := time.Now()

		Font, FontExists := fontCache[Size]
		if !FontExists {
			LoadedFont, Error := ttf.OpenFont("/home/kawak/Documents/Github/MusicalTyper-Go/mplus-1m-medium.ttf", int(Size))
			logger.CheckError(Error)
			fontCache[Size] = LoadedFont
			Font = LoadedFont
		}

		RenderedText, Error := Font.RenderUTF8Blended(Text, *Color)
		logger.CheckError(Error)
		defer RenderedText.Free()

		TextureFromSurface, Error := Renderer.CreateTextureFromSurface(RenderedText)
		logger.CheckError(Error)
		Result := TextureCache{RenderedText.W, RenderedText.H, TextureFromSurface}

		if _, MapExists := textureCache[Size]; !MapExists {
			textureCache[Size] = map[string]TextureCache{}
		}

		textureCache[Size][Text] = Result
		Texture = Result
		fmt.Printf("Created \"%s\" texture. Size:%d. Took %dμs\n", Text, Size, time.Now().Sub(Begin).Microseconds())
	}

	var Rect sdl.Rect
	if alignment == LeftAlign {
		Rect = sdl.Rect{
			X: int32(x),
			Y: int32(y),
			W: Texture.Width,
			H: Texture.Height,
		}
	} else {
		Rect = sdl.Rect{
			X: int32(x) - Texture.Width,
			Y: int32(y),
			W: Texture.Width,
			H: Texture.Height,
		}
	}

	Renderer.Copy(Texture.Texture, nil, &Rect)
	return int(Texture.Width), int(Texture.Height)
}

func DrawTypingText(Renderer *sdl.Renderer, x, y int, Font FontSize, TypedText, RemainingText string) {
	DrawText(Renderer, x, y, RightAlign, Font, TypedText, Constants.TypedTextColor)
	DrawText(Renderer, x, y, LeftAlign, Font, RemainingText, Constants.RemainingTextColor)
}

func DrawRect(Renderer *sdl.Renderer, Color *sdl.Color, x, y, width, height int32) {
	Renderer.SetDrawColor(Color.R, Color.G, Color.B, Color.A)
	Renderer.FillRect(&sdl.Rect{x, y, width, height})
}

func Quit() {
	for _, v := range textureCache {
		for _, t := range v {
			t.Texture.Destroy()
		}
	}
	textureCache = map[FontSize]map[string]TextureCache{}

	for _, v := range fontCache {
		v.Close()
	}
}
