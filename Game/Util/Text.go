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
	textureCache = map[FontSize]map[string]*TextureCache{}
)

const (
	BigFont      FontSize = 72
	JapaneseFont          = 48
	AlphabetFont          = 32
	FullFont              = 24
	RankFont              = 20
	SystemFont            = 16
)

func getTextureCacheKey(Text string, Color *sdl.Color) string {
	return fmt.Sprintf("%s,%d,%d,%d,%d", Text, Color.R, Color.B, Color.B, Color.A)
}

func makeTexture(Renderer *sdl.Renderer, Size FontSize, Text string, Color *sdl.Color) *TextureCache {
	logger := Logger.NewLogger("makeTexture")
	CacheKey := getTextureCacheKey(Text, Color)

	Texture, TextureExists := textureCache[Size][CacheKey]
	if !TextureExists {
		Begin := time.Now()

		Font, FontExists := fontCache[Size]
		if !FontExists {
			LoadedFont, Error := ttf.OpenFont("./mplus-1m-medium.ttf", int(Size))
			logger.CheckError(Error)
			fontCache[Size] = LoadedFont
			Font = LoadedFont
		}

		RenderedText, Error := Font.RenderUTF8Blended(Text, *Color)
		logger.CheckError(Error)
		defer RenderedText.Free()

		TextureFromSurface, Error := Renderer.CreateTextureFromSurface(RenderedText)
		logger.CheckError(Error)
		Result := &TextureCache{RenderedText.W, RenderedText.H, TextureFromSurface}

		if _, MapExists := textureCache[Size]; !MapExists {
			textureCache[Size] = map[string]*TextureCache{}
		}

		textureCache[Size][CacheKey] = Result
		Texture = Result
		fmt.Printf("Created \"%s\" texture. Key: %s Size:%d. Took %dμs\n", Text, CacheKey, Size, time.Now().Sub(Begin).Microseconds())
	}
	return Texture
}

func DrawText(Renderer *sdl.Renderer, x, y int, alignment AlignmentType, Size FontSize, Text string, Color *sdl.Color) (int, int) {
	logger := Logger.NewLogger("DrawText")
	if Text == "" {
		return 0, 0
	}

	Texture := makeTexture(Renderer, Size, Text, Color)

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

	Error := Renderer.Copy(Texture.Texture, nil, &Rect)
	logger.CheckError(Error)
	return int(Texture.Width), int(Texture.Height)
}

func GetTextSize(Renderer *sdl.Renderer, Size FontSize, Text string, Color *sdl.Color) (int, int) {
	Texture := makeTexture(Renderer, Size, Text, Color)
	return int(Texture.Width), int(Texture.Height)
}

func DrawTypingText(Renderer *sdl.Renderer, x, y int, Font FontSize, TypedText, RemainingText string) {
	DrawText(Renderer, x, y, RightAlign, Font, TypedText, Constants.TypedTextColor)
	DrawText(Renderer, x, y, LeftAlign, Font, RemainingText, Constants.RemainingTextColor)
}

func DrawFillRect(Renderer *sdl.Renderer, Color *sdl.Color, x, y, width, height int) {
	Renderer.SetDrawColor(Color.R, Color.G, Color.B, Color.A)
	Renderer.FillRect(&sdl.Rect{int32(x), int32(y), int32(width), int32(height)})
}

func DrawLineRect(Renderer *sdl.Renderer, Color *sdl.Color, x, y, width, height, thickness int) {
	Renderer.SetDrawColor(Color.R, Color.G, Color.B, Color.A)
	X, Y, Width, Height, Thickness := int32(x), int32(y), int32(width), int32(height), int32(thickness)

	Rects := []sdl.Rect{
		{X, Y, Width, Thickness},
		{X, Y, Thickness, Height},
		{X + Width - Thickness, Y, Thickness, Height},
		{X, Y + Height - Thickness, Width, Thickness}}
	Renderer.DrawRects(Rects)
}

func Quit() {
	for _, v := range textureCache {
		for _, t := range v {
			t.Texture.Destroy()
		}
	}
	textureCache = map[FontSize]map[string]*TextureCache{}

	for _, v := range fontCache {
		v.Close()
	}
}
