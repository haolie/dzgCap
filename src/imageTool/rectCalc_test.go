package imageTool

import (
	"fmt"
	"image/color"
	"testing"
)

func TestDisLineL(t *testing.T) {
	img := CapFullScreen()
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
	img := CapFullScreen()
	rect, exists := FindMinRect(img, color.RGBA{
		R: 20,
		G: 24,
		B: 31,
		A: 255,
	})

	if exists {
		fmt.Println(rect)
		rectImg := CapScreen(rect)
		SaveImage(rectImg, fmt.Sprintf("rect%d.png", 999))

	} else {
		fmt.Println("not find rect")
	}

}
