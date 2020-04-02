package DrawHelper

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
	//fmt.Printf("\"%s\" was %t \n", CacheKey, TextureExists)

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
		if Constants.PrintTextureLog {
			fmt.Printf("Created \"%s\" texture. Key: %s Size:%d. Took %dμs\n", Text, CacheKey, Size, time.Now().Sub(Begin).Microseconds())
		}
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
	switch alignment {
	case LeftAlign:
		Rect = sdl.Rect{
			X: int32(x),
			Y: int32(y),
			W: Texture.Width,
			H: Texture.Height,
		}
	case RightAlign:
		Rect = sdl.Rect{
			X: int32(x) - Texture.Width,
			Y: int32(y),
			W: Texture.Width,
			H: Texture.Height,
		}
	case Center:
		Rect = sdl.Rect{
			X: int32(x) - Texture.Width/2,
			Y: int32(y),
			W: Texture.Width,
			H: Texture.Height,
		}
	}

	Error := Renderer.Copy(Texture.Texture, nil, &Rect)
	logger.CheckError(Error)
	return int(Texture.Width), int(Texture.Height)
}

//fixme: 計算ガバガバなので斜めの線とか引くと多分バグる
func DrawThickLine(Renderer *sdl.Renderer, x, y, dstx, dsty int, Color *sdl.Color, Thickness int) {
	Renderer.SetDrawColor(Color.R, Color.G, Color.B, Color.A)
	X, Y, DistX, DistY := int32(x), int32(y), int32(dstx), int32(dsty)
	Renderer.DrawRect(&sdl.Rect{X, Y, DistX - X, DistY - Y})
}

func DrawLine(Renderer *sdl.Renderer, x, y, dstx, dsty int, Color *sdl.Color) {
	DrawThickLine(Renderer, x, y, dstx, dsty, Color, 1)
}

func GetTextSize(Renderer *sdl.Renderer, Size FontSize, Text string, Color *sdl.Color) (int, int) {
	Texture := makeTexture(Renderer, Size, Text, Color)
	return int(Texture.Width), int(Texture.Height)
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
