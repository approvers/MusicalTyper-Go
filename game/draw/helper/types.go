package helper

type AlignmentType uint8

const (
	RightAlign AlignmentType = iota
	LeftAlign
	Center
)

func (t AlignmentType) String() string {
	switch t {
	case RightAlign:
		return "RightAlign"
	case LeftAlign:
		return "LeftAlign"
	case Center:
		return "Center"
	default:
		panic("Unknown alignment type")
	}
}
