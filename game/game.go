package game

import (
	"fmt"
	"musicaltyper-go/game/beatmap"
	"musicaltyper-go/game/constants"
	"musicaltyper-go/game/draw/helper"
	"musicaltyper-go/game/logger"
	"musicaltyper-go/game/view"
	mainview "musicaltyper-go/game/view/game"

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

	Window, Error := sdl.CreateWindow(
		constants.WindowTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		constants.WindowWidth,
		constants.WindowHeight,
		sdl.WINDOW_OPENGL,
	)

	Logger.CheckError(Error)
	defer Window.Destroy()

	Renderer, Error := sdl.CreateRenderer(Window, -1, sdl.RENDERER_ACCELERATED)
	Logger.CheckError(Error)
	defer Renderer.Destroy()

	fmt.Println("DrawStart")

	var (
		CurrentView = mainview.NewMainView(beatmap)
		Running     = true
	)

	for Running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				Running = false
			default:
				Running = CurrentView.HandleSDLEvent(Renderer, event)
			}
		}

		CurrentView.Draw(Renderer)

		for event := CurrentView.PollEvent(); event != nil; event = CurrentView.PollEvent() {
			switch ev := event.(type) {
			case *view.ChangeViewEvent:
				CurrentView = ev.ToChangeView
				fmt.Println("Changed view to", CurrentView.GetName())
			}
		}

		sdl.Delay(1000 / constants.FrameRate)
	}
}
