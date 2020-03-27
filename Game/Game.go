package Game

import (
	"MusicalTyper-Go/Game/Beatmap"
	"MusicalTyper-Go/Game/Constants"
	"MusicalTyper-Go/Game/Logger"
	"MusicalTyper-Go/Game/Util"
	"fmt"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"math"
	"strconv"
	"sync"
	"time"
)

var runningMutex sync.Mutex

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

		defer Util.Quit()

		Logger.CheckError(mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096))
		defer mix.CloseAudio()

		mix.VolumeMusic(mix.MAX_VOLUME / 3)
		Music, Error := mix.LoadMUS("/home/kawak/Documents/Github/MusicalTyper-Go/kkiminochikara-edited.mp3")
		Logger.CheckError(Error)
		//Logger.CheckError(Music.Play(1))
		MusicStartTime := time.Now()
		defer Music.Free()

		Window, Error := sdl.CreateWindow(
			Constants.WindowTitle,
			sdl.WINDOWPOS_CENTERED,
			sdl.WINDOWPOS_CENTERED,
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
			AuthorText               = fmt.Sprintf("%s/%s", beatmap.Properties["song_author"], beatmap.Properties["singer"])
			GameState                = NewGameState(beatmap)
			isTmpNextLyricsPrinting  = false //fixme: test
			isContNextLyricsPrinting = false
			//DrawBegin    time.Time
			//DrawFinish time.Time
		)
		fmt.Println("DrawStart")
		for Running {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch e := event.(type) {
				case *sdl.QuitEvent:
					runningMutex.Lock()
					Running = false
					runningMutex.Unlock()
				case *sdl.TextInputEvent:
					ProcessKeyEvent(e.GetText())
				}
			}

			var (
				DrawBegin               = time.Now()
				RemainingTimeGaugeWidth = int(math.Floor(GameState.GetSentenceTimeRemainRatio() * Constants.WindowWidth))
				CurrentSentence         = GameState.Beatmap.Notes[GameState.CurrentSentenceIndex].Sentence
				RankPosX                = int(Constants.WindowWidth * GameState.GetAchievementRate(true))
			)
			FrameCount = (FrameCount + 1) % Constants.FrameRate
			if FrameCount == 0 {
				//Util.PlaySE(Util.TleSE)
			}

			GameState.Update(float64(time.Now().Sub(MusicStartTime).Milliseconds()) / 1000.0)

			//DrawBegin = time.Now()
			Renderer.SetDrawColor(255, 243, 224, 0)
			Renderer.Clear()

			//曲のタイトルと作曲者名
			Util.DrawText(Renderer, Constants.WindowWidth-2, 0, Util.RightAlign, Util.AlphabetFont, beatmap.Properties["title"], Constants.TextColor)
			Util.DrawText(Renderer, Constants.WindowWidth-5, 33, Util.RightAlign, Util.SystemFont, AuthorText, Constants.TypedTextColor)

			//残り時間ゲージ
			Util.DrawFillRect(Renderer, Util.GetMoreBlackishColor(Constants.BackgroundColor, 25), 0, 60, Constants.WindowWidth, 130)
			Util.DrawFillRect(Renderer, Util.GetMoreBlackishColor(Constants.BackgroundColor, 50), 0, 60, RemainingTimeGaugeWidth, 130)

			//タイプする文字(ひらがな), ローマ字, 歌詞
			Util.DrawTypingText(Renderer, Constants.WindowWidth/2, 80, Util.JapaneseFont, CurrentSentence.GetTypedText(), CurrentSentence.GetRemainingText())
			Util.DrawTypingText(Renderer, Constants.WindowWidth/2, 130, Util.FullFont, CurrentSentence.GetTypedRoma(), CurrentSentence.GetRemainingRoma())
			Util.DrawText(Renderer, Constants.Margin-12, 60, Util.LeftAlign, Util.FullFont, CurrentSentence.OriginalSentence, Constants.LyricTextColor)

			//コンボ (数字と"chain")
			ComboTextWidth, _ :=
				Util.DrawText(Renderer, Constants.Margin-12, 157, Util.LeftAlign, Util.FullFont, strconv.Itoa(GameState.Combo), Constants.ComboTextColor)
			Util.DrawText(Renderer, ComboTextWidth+5, 165, Util.LeftAlign, Util.SystemFont, "chain", Constants.ComboChainTextColor)

			//正解率ゲージ　100%なら赤色
			if Acc := CurrentSentence.GetAccuracy(); Acc == 1 {
				Util.DrawFillRect(Renderer, Constants.RedColor, 0, 60, int(Acc), 3)
			} else {
				Util.DrawFillRect(Renderer, Constants.GreenThickColor, 0, 60, int(Acc), 3)
			}
			//正解率ゲージの上に出るランク
			Util.DrawText(Renderer, RankPosX, 168, Util.RightAlign, Util.SystemFont, Constants.RankTexts[GameState.GetRank()], Constants.TypedTextColor)

			//達成率ゲージ
			if GotRank := GameState.GetRank(); GotRank > 0 {
				Rate := Constants.RankPoints[GotRank-1] / 100
				Util.DrawFillRect(Renderer, Constants.RedColor, 0, 187, int(Constants.WindowWidth*Rate), 3)
			}
			if GameState.GetAchievementRate(false) < 0.8 {
				Util.DrawFillRect(Renderer, Constants.GreenThickColor, 0, 187, int(Constants.WindowWidth*GameState.GetAchievementRate(true)), 3)
			} else {
				Util.DrawFillRect(Renderer, Constants.BlueThickColor, 0, 187, int(Constants.WindowWidth*GameState.GetAchievementRate(true)), 3)
			}

			//キーボード
			if isTmpNextLyricsPrinting || isContNextLyricsPrinting {
				for i := 0; i < 3; i++ {
					Index := i + GameState.CurrentSentenceIndex + 1
					Note := GameState.Beatmap.Notes[Index]
					if Index >= len(GameState.Beatmap.Notes) {
						break
					}
					Util.DrawText(Renderer, 5, 193+60*i, Util.LeftAlign, Util.SystemFont, fmt.Sprintf("[%d]", Index), Constants.TextColor)
					//if game_info.score.score[lyrics_index][1][:1] != "/": -> :thinking:
					Util.DrawText(Renderer, 5, 210+60*i, Util.LeftAlign, Util.FullFont, Note.Sentence.HiraganaSentence, Constants.TextColor)
					Util.DrawText(Renderer, 5, 230+60*i, Util.LeftAlign, Util.SystemFont, Note.Sentence.GetRoma(), Constants.TextColor)
				}
			} else {
				if GameState.IsInputDisabled {
					Util.DrawKeyboard(Renderer, "", nil)
				} else {
					Util.DrawKeyboard(Renderer, Util.Substring(CurrentSentence.GetRemainingRoma(), 0, 1), nil)
				}
			}

			//スコア表示
			Text := fmt.Sprintf("%08d", GameState.Point)
			if GameState.Point < 0 {
				ScoreColor := Constants.BlueThickColor
				if FrameCount%20 < 10 {
					ScoreColor = Constants.RedColor
				}
				Util.DrawText(Renderer, 5, 20, Util.LeftAlign, Util.AlphabetFont, Text, ScoreColor)
			} else {
				Util.DrawText(Renderer, 5, 20, Util.LeftAlign, Util.AlphabetFont, Text, Constants.BlueThickColor)
			}

			//キーボードの下の線
			Util.DrawThickLine(Renderer, 0, 375, Constants.WindowWidth, 375, Constants.TypedTextColor, 2)

			//タイプ速度の表示とバー
			Util.DrawText(Renderer, Constants.Margin, 382, Util.LeftAlign, Util.SystemFont, "タイピング速度", Constants.TypedTextColor)
			if GameState.GetKeyTypePerSecond() > 4 {
				Color := Constants.RedColor
				if !(FrameCount%10 < 5) {
					Color = Util.GetMoreBlackishColor(Color, 30)
				}
				Util.DrawFillRect(Renderer, Color, Constants.Margin, 405, Constants.WindowWidth-Constants.Margin*2, 20)
			} else {
				Util.DrawFillRect(Renderer, Constants.GreenThinColor, Constants.Margin, 405, Constants.WindowWidth-Constants.Margin*2, 20)
				Color := Util.GetMoreBlackishColor(Constants.GreenThinColor, 50)
				Util.DrawFillRect(Renderer, Color, Constants.Margin, 405, int(GameState.GetKeyTypePerSecond()/4*(Constants.WindowWidth*2)), 20)
			}
			Text = fmt.Sprintf("%4.2f Char/sec", GameState.GetKeyTypePerSecond())
			Util.DrawText(Renderer, Constants.WindowWidth/2, 402, Util.Center, Util.SystemFont, Text, Constants.TextColor)

			//正解率
			Util.DrawText(Renderer, Constants.Margin, 430, Util.LeftAlign, Util.SystemFont, "正解率", Constants.TypedTextColor)
			Color := Util.GetMoreBlackishColor(Constants.RedColor, 50)
			Util.DrawFillRect(Renderer, Color, Constants.Margin+5, 510, int(GameState.GetAccuracy()*250), 3)
			Acc := GameState.GetAccuracy()
			Text = fmt.Sprintf("%05.1f%%", Acc*100)
			Red := Constants.RedColor
			Color = &sdl.Color{uint8(Acc) * Red.R, uint8(Acc) * Red.G, uint8(Acc) * Red.B, 255}
			Util.DrawText(Renderer, Constants.Margin+5, 430, Util.LeftAlign, Util.BigFont, Text, Color)

			//ランク
			Util.DrawText(Renderer, Constants.Margin+320, 430, Util.LeftAlign, Util.SystemFont, "達成率", Constants.TypedTextColor)
			Text = fmt.Sprintf("%05.1f%%", GameState.GetAchievementRate(false))
			Util.DrawText(Renderer, Constants.Margin+330, 430, Util.LeftAlign, Util.BigFont, Text, Constants.BlueThickColor)

			//DrawTime
			DrawTimeStr := fmt.Sprintf("%4d μs", int(time.Now().Sub(DrawBegin).Microseconds()))
			Util.DrawText(Renderer, 3, -3, Util.LeftAlign, Util.SystemFont, DrawTimeStr, Constants.TextColor)

			Renderer.Present()
			//fmt.Println("Drawtime:", DrawFinish.Sub(DrawBegin).Microseconds(), "μs")
			sdl.Delay(1000 / Constants.FrameRate)
		}
	})
	return 0
}

func ProcessKeyEvent(Input string) {

}
