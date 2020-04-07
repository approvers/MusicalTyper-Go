package top

import (
	Constants "musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/view/game/component"

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

		helper.DrawText(Renderer, pos.FromXY(Constants.WindowWidth-2, 0), helper.RightAlign, helper.AlphabetFont, Title, Constants.TextColor)
		helper.DrawText(Renderer, pos.FromXY(Constants.WindowWidth-5, 33), helper.RightAlign, helper.SystemFont, AuthorText, Constants.TypedTextColor)
	}
}
