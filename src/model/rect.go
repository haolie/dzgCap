package model

type Rect struct {
	X int
	Y int
	W int
	H int

	// 相对位置
	rx int
	ry int
}

func NewRect(x, y, w, h int) *Rect {
	return &Rect{X: x, Y: y, W: w, H: h}
}

// Relative
// @description: 返回相对位置区域
// parameter:
//		@receiver r:
//		@rx:
//		@ry:
// return:
//		@Rect:
func (r *Rect) Relative(rx, ry int) Rect {
	return Rect{
		X:  r.X + r.rx - rx,
		Y:  r.Y + r.ry - ry,
		W:  r.W,
		H:  r.H,
		rx: rx,
		ry: ry,
	}
}

// 返回绝对区域
func (r *Rect) DRect() Rect {
	return r.Relative(0, 0)
}

// Move
// @description: 移动
// parameter:
//		@receiver r:
//		@x:
//		@y:
// return:
//		@Rect:
func (r *Rect) Move(x, y int) Rect {
	return Rect{
		X: r.X + x,
		Y: r.Y + y,
		H: r.H,
		W: r.W,
	}
}
