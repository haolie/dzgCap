package capMager

import (
	"image"
	"path"

	"github.com/go-vgo/robotgo"

	"dzgCap/src/imageTool"
	"dzgCap/src/model"
)

const con_pic_dir = "./pic"

var (
	imageMap = make(map[string]image.Image, 8)
)

func GetCashImg(name string) (img image.Image, err error) {
	if img, exists := imageMap[name]; exists {
		return img, nil
	}

	fileName := path.Join(con_pic_dir, name+".png")
	img, err = imageTool.LoadImage(fileName)
	if err != nil {
		return nil, err
	}

	imageMap[name] = img

	return
}

func SaveRectImg(rect model.Rect, name string) error {
	img := imageTool.CapScreen(rect)
	return SaveImg(img, name)
}

func SaveImg(img image.Image, name string) error {
	fileName := path.Join(con_pic_dir, name+".png")
	err := imageTool.SaveImage(img, fileName)
	if err != nil {
		return err
	}

	imageMap[name] = img

	return nil
}

func CompareToCash(img image.Image, name string) (isSame bool, err error) {
	cashImg, err := GetCashImg(name)
	if err != nil {
		return
	}

	return imageTool.CompareImage(img, cashImg), nil
}

func CompareRectToCash(rect model.Rect, name string) (isSame bool, err error) {
	img := imageTool.CapScreen(rect)
	return CompareToCash(img, name)
}

func ClickPoint(x, y int) {
	robotgo.MoveClick(x, y)
}

func ClickRect(rect model.Rect) {
	robotgo.MoveClick(rect.X+rect.W/2, rect.Y+rect.H/2)
}
