package imageTool

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"os"

	"github.com/go-vgo/robotgo"

	"dzgCap/src/model"
)

const (
	con_compare_num = 30
)

func LoadImage(name string) (img image.Image, err error) {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err = image.Decode(f)

	return
}

func SaveImage(img image.Image, name string) error {
	outFile, err := os.Create(name)
	defer outFile.Close()
	if err != nil {
		return err
	}
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, img)
	if err != nil {
		return err
	}
	return b.Flush()
}

func CompareImage(imgA, imgB image.Image) bool {
	if imgA.Bounds() != imgB.Bounds() {
		return false
	}

	s := imgA.Bounds().Size()
	x := con_compare_num
	if s.X < x {
		x = s.X
	}
	xp := s.X / x

	y := con_compare_num
	if s.Y < y {
		y = s.Y
	}
	yp := s.Y / y

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			p := &image.Point{X: i * xp, Y: j * yp}
			if !CompareColor(imgA.At(p.X, p.Y), imgB.At(p.X, p.Y)) {
				return false
			}
		}
	}

	return true
}

func CapScreen(rect model.Rect) image.Image {
	bitmap := robotgo.CaptureScreen(rect.X, rect.Y, rect.W, rect.H)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bitmap)

	return robotgo.ToImage(bitmap)
}

func CapFullScreen() image.Image {
	bitmap := robotgo.CaptureScreen()
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bitmap)

	return robotgo.ToImage(bitmap)
}

func CompareColor(c1 color.Color, c2 color.Color) (isEqual bool) {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	isEqual = r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
	return
}
