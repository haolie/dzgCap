package capMager

import (
	"fmt"
	"image/color"
	"testing"

	"dzgCap/src/imageTool"
)

func TestDisLineL(t *testing.T) {
	img := imageTool.CapFullScreen()
	//imageTool.SaveImage(img,"t2.png")
	lines := lineFromImageL(img, color.RGBA{
		R: 20,
		G: 24,
		B: 31,
		A: 255,
	})

	fmt.Println(lines)
}

func TestFindRect(t *testing.T) {
	img := imageTool.CapFullScreen()
	rect, exists := FindMinRect(img, color.RGBA{
		R: 20,
		G: 24,
		B: 31,
		A: 255,
	})

	if exists {
		fmt.Println(rect)
		rectImg := imageTool.CapScreen(rect)
		imageTool.SaveImage(rectImg, fmt.Sprintf("rect%d.png", 999))

	} else {
		fmt.Println("not find rect")
	}

}
