package model

type Rect struct {
	X int
	Y int
	W int
	H int
}

func NewRect(x, y, w, h int) *Rect {
	return &Rect{X: x, Y: y, W: w, H: h}
}
