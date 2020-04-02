package helper

// Max returns max of a and b
func Max(a, b uint8) uint8 {
	if a > b {
		return a
	} else {
		return b
	}
}

// MinUInt8 returns min of a and b, then cast it uint8
func MinUInt8(a, b int) uint8 {
	if a < b {
		return uint8(a)
	} else {
		return uint8(b)
	}
}
