package result

import (
	"musicaltyper-go/game/rank"
	"musicaltyper-go/game/view"
	"musicaltyper-go/game/view/result/component"
	"musicaltyper-go/game/view/result/component/bottom"
	"musicaltyper-go/game/view/result/component/center"
	"musicaltyper-go/game/view/result/component/top"

	"github.com/veandco/go-sdl2/sdl"
)

type GameResult struct {
	Rank            rank.Rank
	Point           int
	TypeSpeed       float64
	Accuracy        float64
	AchievementRate float64
	MapInfo         map[string]string
}

type resultView struct {
	result *GameResult
}

func NewResultView(result *GameResult) view.View {
	Result := resultView{}
	Result.result = result
	return &Result
}

func (view *resultView) GetName() string {
	return "ResultView"
}

func (view *resultView) HandleSDLEvent(_ *sdl.Renderer, event sdl.Event) bool {
	switch e := event.(type) {
	case *sdl.KeyboardEvent:
		key := e.Keysym.Sym
		if e.Type == sdl.KEYDOWN {
			switch key {
			case sdl.K_ESCAPE:
				return false
			}
		}
	}

	return true
}

func (view *resultView) PollEvent() view.Event {
	return nil
}

func (view *resultView) Draw(Renderer *sdl.Renderer) {
	Renderer.SetDrawColor(255, 243, 224, 0)
	Renderer.Clear()

	Components := []component.Drawable{
		top.SongInfo(view.result.MapInfo),
		center.RankText(view.result.Rank),
		center.ScoreText(view.result.Point, view.result.Accuracy, view.result.AchievementRate, view.result.Rank),
		center.SpeedGauge(view.result.TypeSpeed),
		bottom.KeyText(),
	}

	for _, v := range Components {
		v(Renderer)
	}

	Renderer.Present()
}
