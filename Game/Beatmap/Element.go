package Beatmap

type NoteType int

const (
	NORMAL NoteType = iota
	CAPTION
	BLANK
)

type Note struct {
	Type     NoteType
	Time     float64
	Sentence Sentence
	isEnd    bool
	Caption  string
}

type Section struct {
	Time float64
	ID   string
}

type Zone struct {
	Time    float64
	Name    string
	isBegin bool
}

func newNote(Sec float64, Lyric, Pron string) Note {
	Result := Note{}
	Result.Time = Sec
	Result.Sentence = NewSentence(Lyric, Pron)
	Result.Type = NORMAL
	return Result
}

func newBlankNote(Sec float64) Note {
	Result := Note{}
	Result.Type = BLANK
	Result.Time = Sec
	return Result
}

func newCaptionNote(Sec float64, Caption string) Note {
	Result := Note{}
	Result.Type = CAPTION
	Result.Time = Sec
	Result.Caption = Caption
	return Result
}

func endMap(Sec float64) Note {
	Result := Note{}
	Result.Time = Sec
	Result.isEnd = true
	return Result
}

func newSection(Sec float64, ID string) Section {
	Result := Section{}
	Result.Time = Sec
	Result.ID = ID
	return Result
}

func newZone(Sec float64, isBegin bool, Name string) Zone {
	Result := Zone{}
	Result.Time = Sec
	Result.isBegin = isBegin
	Result.Name = Name
	return Result
}
