package imageTool

import (
	"fmt"
	"path"
	"testing"

	"dzgCap/src/model"
)

func TestImgCapAndSave(t *testing.T) {
	capImg := CapScreen(model.Rect{X: 2414, Y: 1085, W: 8, H: 10})
	err := SaveImage(capImg, path.Join("../../config", "555.png"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
