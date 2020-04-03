package helper

import (
	"fmt"
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/area"
	"musicaltyper-go/game/draw/color"
	"musicaltyper-go/game/draw/pos"
	"musicaltyper-go/game/draw/size"
	Logger "musicaltyper-go/game/logger"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// FontSize means size of font
type FontSize uint8

type textureWithSize struct {
	Width, Height int32
	Texture       *sdl.Texture
}

var (
	fontCache    = map[FontSize]*ttf.Font{}
	textureCache = map[FontSize]map[string]*textureWithSize{}
)

const (
	// BigFont is bigger font size
	BigFont FontSize = 72
	// JapaneseFont is font size for japanese
	JapaneseFont = 48
	// AlphabetFont is font size for ASCII
	AlphabetFont = 32
	// FullFont is font size for full width
	FullFont = 24
	// RankFont is font size for rank of player
	RankFont = 20
	// SystemFont is default font size
	SystemFont = 16
)

func gettextureWithSizeKey(Text string, Color color.Color) string {
	return fmt.Sprintf("%s,%s", Text, Color.String())
}

func makeTexture(Renderer *sdl.Renderer, Size FontSize, Text string, Color color.Color) *textureWithSize {
	logger := Logger.NewLogger("makeTexture")
	CacheKey := gettextureWithSizeKey(Text, Color)

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

		RenderedText, Error := Font.RenderUTF8Blended(Text, *Color.Cast())
		logger.CheckError(Error)
		defer RenderedText.Free()

		TextureFromSurface, Error := Renderer.CreateTextureFromSurface(RenderedText)
		logger.CheckError(Error)
		Result := &textureWithSize{RenderedText.W, RenderedText.H, TextureFromSurface}

		if _, MapExists := textureCache[Size]; !MapExists {
			textureCache[Size] = map[string]*textureWithSize{}
		}

		textureCache[Size][CacheKey] = Result
		Texture = Result
		if Constants.PrintTextureLog {
			fmt.Printf("Created \"%s\" texture. Key: %s Size:%d. Took %dμs\n", Text, CacheKey, Size, time.Now().Sub(Begin).Microseconds())
		}
	}
	return Texture
}

// DrawText renders text
func DrawText(Renderer *sdl.Renderer, p pos.Pos, alignment AlignmentType, Size FontSize, Text string, Color color.Color) (int, int) {
	logger := Logger.NewLogger("DrawText")
	if Text == "" {
		return 0, 0
	}

	Texture := makeTexture(Renderer, Size, Text, Color)
	x, y := p.X(), p.Y()

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

// DrawThickLine renders thick line
//fixme: 計算ガバガバなので斜めの線とか引くと多分バグる
func DrawThickLine(Renderer *sdl.Renderer, from, to pos.Pos, Color color.Color, Thickness int) {
	Color.ProxyColor(Renderer)
	Renderer.DrawRect(area.FromTwoPos(from, to).ToRect())
}

// DrawLine render line
func DrawLine(Renderer *sdl.Renderer, from, to pos.Pos, Color color.Color) {
	DrawThickLine(Renderer, from, to, Color, 1)
}

// GetTextSize calculates dimension of text by actual rendering
func GetTextSize(Renderer *sdl.Renderer, Size FontSize, Text string, Color color.Color) size.Size {
	Texture := makeTexture(Renderer, Size, Text, Color)
	return size.FromWH(int(Texture.Width), int(Texture.Height))
}

// DrawFillRect renders filled rect
func DrawFillRect(Renderer *sdl.Renderer, Color color.Color, a area.Area) {
	Color.ProxyColor(Renderer)
	Renderer.FillRect(a.ToRect())
}

// DrawLineRect render rect by lines
func DrawLineRect(Renderer *sdl.Renderer, Color color.Color, a area.Area, thickness int) {
	Color.ProxyColor(Renderer)
	X, Y, Width, Height, Thickness := int32(a.X()), int32(a.Y()), int32(a.W()), int32(a.H()), int32(thickness)

	Rects := []sdl.Rect{
		{X: X, Y: Y, W: Width, H: Thickness},
		{X: X, Y: Y, W: Thickness, H: Height},
		{X: X + Width - Thickness, Y: Y, W: Thickness, H: Height},
		{X: X, Y: Y + Height - Thickness, W: Width, H: Thickness}}
	Renderer.DrawRects(Rects)
}

// Quit destroys textures
func Quit() {
	for _, v := range textureCache {
		for _, t := range v {
			t.Texture.Destroy()
		}
	}
	textureCache = map[FontSize]map[string]*textureWithSize{}

	for _, v := range fontCache {
		v.Close()
	}
}
