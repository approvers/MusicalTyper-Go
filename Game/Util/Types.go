package Util

type AlignmentType uint8

const (
	RightAlign AlignmentType = iota
	LeftAlign
)

func (t AlignmentType) String() string {
	switch t {
	case RightAlign:
		return "RightAlign"
	case LeftAlign:
		return "LeftAlign"
	default:
		panic("Unknown alignment type")
	}
}
