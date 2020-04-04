package game

import (
	"fmt"
	"musicaltyper-go/game/beatmap"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/component"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/draw/view/mainview"
	"musicaltyper-go/game/logger"
	"musicaltyper-go/game/state"
	"time"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Run runs game with beatmap
func Run(beatmap *beatmap.Beatmap) {
	Logger := logger.NewLogger("GameRun")

	Logger.CheckError(sdl.Init(sdl.INIT_VIDEO))
	Logger.CheckError(sdl.Init(sdl.INIT_AUDIO))
	defer sdl.Quit()

	Logger.CheckError(mix.Init(mix.INIT_OGG))
	defer mix.Quit()

	Logger.CheckError(ttf.Init())
	defer ttf.Quit()

	defer helper.Quit()

	Logger.CheckError(mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 1024))
	defer mix.CloseAudio()
	mix.AllocateChannels(constants.AudioChannelNum)

	mix.VolumeMusic(mix.MAX_VOLUME / 10)
	Music, Error := mix.LoadMUS(beatmap.Properties["song_data"])
	Logger.CheckError(Error)

	Logger.CheckError(Music.Play(1))
	MusicStartTime := time.Now()
	defer Music.Free()

	Window, Error := sdl.CreateWindow(
		constants.WindowTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		constants.WindowWidth,
		constants.WindowHeight,
		sdl.WINDOW_OPENGL)
	Logger.CheckError(Error)
	defer Window.Destroy()

	Renderer, Error := sdl.CreateRenderer(Window, -1, sdl.RENDERER_ACCELERATED)
	Logger.CheckError(Error)
	defer Renderer.Destroy()

	var (
		Running                  = true
		FrameCount               = 0
		gameState                = state.NewGameState(beatmap)
		isContNextLyricsPrinting = false
		//DrawBegin    time.Time
		//DrawFinish time.Time
	)
	fmt.Println("DrawStart")
	for Running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				Running = false

			case *sdl.KeyboardEvent:
				key := e.Keysym.Sym
				if e.Type == sdl.KEYDOWN {
					switch key {
					case sdl.K_ESCAPE:
						Running = false

					case sdl.K_LSHIFT, sdl.K_RSHIFT:
						isContNextLyricsPrinting = !isContNextLyricsPrinting

					default:
						gameState.ParseKeyInput(Renderer, key, isContNextLyricsPrinting)
					}
				}
			}
		}

		FrameCount = (FrameCount + 1) % constants.FrameRate
		gameState.Update(float64(time.Now().Sub(MusicStartTime).Milliseconds()) / 1000.0)

		var NormalizedRemainingTime float64
		if len(gameState.Beatmap.Notes) <= gameState.CurrentSentenceIndex+1 {
			NormalizedRemainingTime = 1
		} else {
			CurrentSentenceStartTime := gameState.Beatmap.Notes[gameState.CurrentSentenceIndex].Time
			NextSentenceStartTime := gameState.Beatmap.Notes[gameState.CurrentSentenceIndex+1].Time
			CurrentSentenceDuration := NextSentenceStartTime - CurrentSentenceStartTime
			CurrentTimeInCurrentSentence := CurrentSentenceDuration - gameState.CurrentTime + CurrentSentenceStartTime
			NormalizedRemainingTime = CurrentTimeInCurrentSentence / CurrentSentenceDuration
		}

		nextSentenceIndex := gameState.CurrentSentenceIndex + 1
		Time := time.Now()
		Context := component.DrawContext{
			Renderer:                Renderer,
			Properties:              beatmap.Properties,
			CurrentSentence:         *beatmap.Notes[gameState.CurrentSentenceIndex].Sentence,
			Rank:                    gameState.GetRank(),
			NormalizedRemainingTime: NormalizedRemainingTime,
			AchievementRate:         gameState.GetAchievementRate(false),
			Point:                   gameState.Point,
			Combo:                   gameState.Combo,
			NextLyrics:              beatmap.Notes[nextSentenceIndex : nextSentenceIndex+3],
			Accuracy:                gameState.GetAccuracy(),
			TypingSpeed:             gameState.GetKeyTypePerSecond(),
			IsKeyboardDisabled:      isContNextLyricsPrinting,
			FrameCount:              FrameCount,
			DrawBeginTime:           &Time,
		}

		Renderer.SetDrawColor(255, 243, 224, 0)
		Renderer.Clear()

		mainview.Draw(&Context)

		Renderer.Present()
		sdl.Delay(1000 / constants.FrameRate)
	}
}
