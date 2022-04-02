package capMager

import (
	"fmt"
	"image/color"
	"testing"

	"dzgCap/imageTool"
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
	rectList, exists := FindRect(img, color.RGBA{
		R: 20,
		G: 24,
		B: 31,
		A: 255,
	})

	if exists {
		fmt.Println(rectList)
		for i, rectItem := range rectList {
			rectImg := imageTool.CapScreen(rectItem)
			imageTool.SaveImage(rectImg, fmt.Sprintf("rect%d.png", i))
			fmt.Println(rectItem)
		}
	} else {
		fmt.Println("not find rect")
	}

}
