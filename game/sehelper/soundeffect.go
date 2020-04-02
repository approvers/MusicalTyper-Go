package sehelper

import (
	Logger "musicaltyper-go/game/logger"

	"github.com/veandco/go-sdl2/mix"
)

type SEType string

const (
	SuccessSE        SEType = "ses/success.wav"
	SpecialSuccessSE SEType = "ses/special.wav"
	FailedSE         SEType = "ses/failed.wav"
	UnneccesarySE    SEType = "ses/unneccesary.wav"
	GameoverSE       SEType = "ses/gameover.wav"
	AcSE             SEType = "ses/ac.wav"
	WaSE             SEType = "ses/wa.wav"
	FastSE           SEType = "ses/fast.wav"
	TleSE            SEType = "ses/tle.wav"
)

var (
	seCache = map[SEType]*mix.Chunk{}
)

func Play(seType SEType) {
	logger := Logger.NewLogger("Play")
	SE, Exists := seCache[seType]
	if !Exists {
		LoadedSE, Err := mix.LoadWAV(string(seType))
		logger.CheckError(Err)
		seCache[seType] = LoadedSE
		SE = LoadedSE
	}
	_, Err := SE.Play(1, 0)
	logger.CheckError(Err)
}
