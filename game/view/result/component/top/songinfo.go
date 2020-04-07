package top

import (
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/view/result/component"

	"musicaltyper-go/game/draw/pos"

	"github.com/veandco/go-sdl2/sdl"
)

// SongInfo draws title, author, and singer
func SongInfo(Properties map[string]string) component.Drawable {
	return func(Renderer *sdl.Renderer) {
		Title := Properties["title"]
		Author, AuthorExists := Properties["song_author"]
		Singer, SingerExists := Properties["singer"]

		var AuthorText string
		if AuthorExists && SingerExists {
			AuthorText = Author + "/" + Singer
		} else if AuthorExists {
			AuthorText = Author
		} else {
			AuthorText = Singer
		}

		helper.DrawText(Renderer, pos.FromXY(Constants.Margin, 0), helper.LeftAlign, helper.JapaneseFont, Title, Constants.TextColor)
		helper.DrawText(Renderer, pos.FromXY(Constants.Margin, 50), helper.LeftAlign, helper.FullFont, AuthorText, Constants.TextColor.Brighter(25))
		helper.DrawThickLine(Renderer, pos.FromXY(0, 90), pos.FromXY(Constants.WindowWidth, 90), Constants.TextColor.Brighter(100), 2)
	}
}
