package main

import (
	"MusicalTyper-Go/Game"
	"MusicalTyper-Go/Game/Beatmap"
	Logger "MusicalTyper-Go/Game/Logger"
	"os"
	"runtime"
)

func InitMap() *Beatmap.Beatmap {
	logger := Logger.NewLogger("Main")

	if len(os.Args) < 2 {
		logger.FatalError("Song file is not specified.")
	}
	BeatMapPath := os.Args[1]
	Stat, Err := os.Stat(BeatMapPath)
	logger.CheckError(Err)

	if !Stat.Mode().IsRegular() {
		logger.FatalError("Specified path isn't file or doesn't exists.")
	}

	return Beatmap.LoadMap(BeatMapPath)
}

func main() {
	//Be sure this goroutine to run on main thread.
	runtime.LockOSThread()

	Map := InitMap()
	Game.Run(Map)
}
