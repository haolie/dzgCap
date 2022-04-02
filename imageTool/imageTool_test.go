package imageTool

import (
	"fmt"
	"path"
	"testing"

	"dzgCap/model"
)

func TestImgCapAndSave(t *testing.T) {
	capImg := CapScreen(model.Rect{0, 0, 2560, 1408})
	err := SaveImage(capImg, path.Join("../config", "test.png"))
	if err != nil {
		fmt.Println(err)
		return
	}

	getImg, err := LoadImage(path.Join("../config", "test.png"))
	if err != nil {
		fmt.Println(err)
		return
	}

	if CompareImage(capImg, getImg) {
		fmt.Println("Success")
	}

}
