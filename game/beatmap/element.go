package beatmap

// NoteType is a kind of Note
type NoteType int

const (
	// NORMAL is a kind of Note, needs for user to type
	NORMAL NoteType = iota
	// CAPTION is a kind of Note, doesn't needs for user to do anything
	CAPTION
	// BLANK is a kind of Note, nothing presents
	BLANK
)

// Note has Sentence and its timing
type Note struct {
	Type     NoteType
	Time     float64
	Sentence *Sentence
	isEnd    bool
	Caption  string
}

// Section expresses time as concrete entity
type Section struct {
	Time float64
	ID   string
}

func newNote(Sec float64, Lyric, Pron string) *Note {
	Result := new(Note)
	Result.Time = Sec
	Result.Sentence = NewSentence(Lyric, Pron)
	Result.Type = NORMAL
	return Result
}

func newBlankNote(Sec float64) *Note {
	Result := new(Note)
	Result.Type = BLANK
	Result.Time = Sec
	return Result
}

func newCaptionNote(Sec float64, Caption string) *Note {
	Result := new(Note)
	Result.Type = CAPTION
	Result.Time = Sec
	Result.Caption = Caption
	Result.Sentence = NewSentence("", "")
	return Result
}

func endMap(Sec float64) *Note {
	Result := new(Note)
	Result.Time = Sec
	Result.isEnd = true
	return Result
}

func newSection(Sec float64, ID string) *Section {
	Result := new(Section)
	Result.Time = Sec
	Result.ID = ID
	return Result
}
