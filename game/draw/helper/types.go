package helper

// AlignmentType means kind of alignment
type AlignmentType uint8

const (
	// RightAlign means to align to right
	RightAlign AlignmentType = iota
	// LeftAlign means to align to left
	LeftAlign
	// Center means to align to center
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
