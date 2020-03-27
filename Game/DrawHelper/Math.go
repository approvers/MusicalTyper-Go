package DrawHelper

import "math"

func Max(a, b uint8) uint8 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinUInt8(a, b uint8) uint8 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Floor(a int32) int32 {
	return int32(math.Floor(float64(a)))
}
