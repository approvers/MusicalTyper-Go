package top

import (
	Constants "musicaltyper-go/game/constants"
	DrawComponent "musicaltyper-go/game/draw/component"
	DrawHelper "musicaltyper-go/game/draw/helper"
)

type SongInfo struct{}

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

	DrawHelper.DrawText(c.Renderer, Constants.WindowWidth-2, 0, DrawHelper.RightAlign, DrawHelper.AlphabetFont, Title, Constants.TextColor)
	DrawHelper.DrawText(c.Renderer, Constants.WindowWidth-5, 33, DrawHelper.RightAlign, DrawHelper.SystemFont, AuthorText, Constants.TypedTextColor)
}
