package main

import (
	Game "musicaltyper-go/game"
	Beatmap "musicaltyper-go/game/beatmap"
	Logger "musicaltyper-go/game/logger"
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
