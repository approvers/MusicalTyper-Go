package sehelper

import (
	Logger "musicaltyper-go/game/logger"

	"github.com/veandco/go-sdl2/mix"
)

// SEType is kind of sound effect
type SEType string

const (
	// SuccessSE rings when succeed
	SuccessSE SEType = "ses/success.wav"
	// SpecialSuccessSE rings epically when succeed
	SpecialSuccessSE SEType = "ses/special.wav"
	// FailedSE rings when failed
	FailedSE SEType = "ses/failed.wav"
	// UnneccesarySE rings when player did unneccesary thing
	UnneccesarySE SEType = "ses/unneccesary.wav"
	// GameoverSE rings when game was over
	GameoverSE SEType = "ses/gameover.wav"
	// AcSE rings when accepted
	AcSE SEType = "ses/ac.wav"
	// WaSE rings when wrong-answered
	WaSE SEType = "ses/wa.wav"
	// FastSE rings when player was faster
	FastSE SEType = "ses/fast.wav"
	// TleSE rings when player failed to type all
	TleSE SEType = "ses/tle.wav"
)

var (
	seCache = map[SEType]*mix.Chunk{}
)

// Play plays sound effect
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
