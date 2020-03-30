package Game

import (
	"MusicalTyper-Go/Game/Beatmap"
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/DrawComponents"
	"MusicalTyper-Go/Game/DrawComponents/DrawManager"
	"MusicalTyper-Go/Game/DrawHelper"
	GameState2 "MusicalTyper-Go/Game/GameState"
	"MusicalTyper-Go/Game/GameSystem"
	"MusicalTyper-Go/Game/Logger"
	"fmt"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"time"
)

func Run(beatmap *Beatmap.Beatmap) int {
	sdl.Do(func() {
		Logger := Logger.NewLogger("GameRun")

		Logger.CheckError(sdl.Init(sdl.INIT_VIDEO))
		Logger.CheckError(sdl.Init(sdl.INIT_AUDIO))
		defer sdl.Quit()

		Logger.CheckError(mix.Init(mix.INIT_MP3))
		defer mix.Quit()

		Logger.CheckError(ttf.Init())
		defer ttf.Quit()

		defer DrawHelper.Quit()

		Logger.CheckError(mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096))
		defer mix.CloseAudio()

		mix.VolumeMusic(mix.MAX_VOLUME / 3)
		Music, Error := mix.LoadMUS(beatmap.Properties["song_data"])
		Logger.CheckError(Error)

		Logger.CheckError(Music.Play(1))
		MusicStartTime := time.Now()
		defer Music.Free()

		Window, Error := sdl.CreateWindow(
			Constants.WindowTitle,
			sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED,
			Constants.WindowWidth,
			Constants.WindowHeight,
			sdl.WINDOW_OPENGL)
		Logger.CheckError(Error)
		defer Window.Destroy()

		Renderer, Error := sdl.CreateRenderer(Window, -1, sdl.RENDERER_ACCELERATED)
		Logger.CheckError(Error)
		defer Renderer.Destroy()

		var (
			Running                  = true
			FrameCount               = 0
			GameState                = GameState2.NewGameState(beatmap)
			isTmpNextLyricsPrinting  = false
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
					switch e.Type {
					case sdl.KEYUP:
						if key == sdl.K_SPACE {
							isTmpNextLyricsPrinting = false
						}

					case sdl.KEYDOWN:
						switch key {
						case sdl.K_ESCAPE:
							Running = false

						case sdl.K_SPACE:
							isTmpNextLyricsPrinting = true

						case sdl.K_LSHIFT, sdl.K_RSHIFT:
							isContNextLyricsPrinting = !isContNextLyricsPrinting

						default:
							GameSystem.ParseKeyInput(Renderer, GameState, key, isTmpNextLyricsPrinting || isContNextLyricsPrinting)
						}
					}
				}
			}

			FrameCount = (FrameCount + 1) % Constants.FrameRate
			GameState.Update(float64(time.Now().Sub(MusicStartTime).Milliseconds()) / 1000.0)

			Context := DrawComponents.DrawContext{
				Renderer:        Renderer,
				GameState:       GameState,
				PrintNextLyrics: isContNextLyricsPrinting || isTmpNextLyricsPrinting,
				FrameCount:      FrameCount,
				Window:          Window,
			}

			Renderer.SetDrawColor(255, 243, 224, 0)
			Renderer.Clear()

			DrawManager.Draw(&Context)

			Renderer.Present()
			sdl.Delay(1000 / Constants.FrameRate)
		}
	})
	return 0
}
