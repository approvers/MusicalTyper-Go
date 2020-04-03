package size

// Size expresses size on screen
type Size struct {
	width  int
	height int
}

// FromWH makes Size from width and height amount
func FromWH(w, h int) Size {
	return Size{width: w, height: h}
}

// W returns width of Size
func (s Size) W() int {
	return s.width
}

// H returns height of Size
func (s Size) H() int {
	return s.height
}
