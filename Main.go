package main

import (
	"MusicalTyper-Go/Game"
	"MusicalTyper-Go/Game/Beatmap"
	Logger2 "MusicalTyper-Go/Game/Logger"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

func InitMap() *Beatmap.Beatmap {
	Logger := Logger2.NewLogger("Main")

	if len(os.Args) < 2 {
		Logger.FatalError("Song file is not specified.")
	}
	BeatMapPath := os.Args[1]
	Stat, Err := os.Stat(BeatMapPath)
	Logger.CheckError(Err)

	if !Stat.Mode().IsRegular() {
		Logger.FatalError("Specified path isn't file or doesn't exists.")
	}

	return Beatmap.LoadMap(BeatMapPath)
}

func main() {
	Map := InitMap()

	sdl.Main(func() {
		Game.Run(Map)
	})
}
