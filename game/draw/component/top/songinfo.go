package top

import (
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"

	"musicaltyper-go/game/draw/pos"
)

// SongInfo presents title, author, and singer
type SongInfo struct{}

// Draw draws title, author, and singer
func (s SongInfo) Draw(c *DrawComponent.DrawContext) {
	Title := c.GameState.Beatmap.Properties["title"]
	Author, AuthorExists := c.GameState.Beatmap.Properties["song_author"]
	Singer, SingerExists := c.GameState.Beatmap.Properties["singer"]

	var AuthorText string
	if AuthorExists && SingerExists {
		AuthorText = Author + "/" + Singer
	} else if AuthorExists {
		AuthorText = Author
	} else {
		AuthorText = Singer
	}

	DrawHelper.DrawText(c.Renderer, pos.FromXY(Constants.WindowWidth-2, 0), DrawHelper.RightAlign, DrawHelper.AlphabetFont, Title, Constants.TextColor)
	DrawHelper.DrawText(c.Renderer, pos.FromXY(Constants.WindowWidth-5, 33), DrawHelper.RightAlign, DrawHelper.SystemFont, AuthorText, Constants.TypedTextColor)
}
