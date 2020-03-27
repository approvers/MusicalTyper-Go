package TopArea

import (
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/DrawComponents"
	"MusicalTyper-Go/Game/DrawHelper"
)

type SongInfo struct{}

func (s SongInfo) Draw(c *DrawComponents.DrawContext) {
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
