package gameAreaModel

import (
	"image"

	. "dzgCap/src/model"
)

func GetModelKeyWithRect(r Rect) (modelKey string, exists bool) {
	return
}

func GetRect(areaModel string, taskType int32, key string) (r Rect, exists bool) {
	return
}

func GetPoint(areaModel string, taskType int32, key string) (p Point, exists bool) {
	return
}

func GetImage(areaModel string, taskType int32, key string) (img image.Image, exists bool) {
	return
}

func SaveAreaModel(areaModelKey string, r Rect) {

}

func SaveRect(areaModel string, taskType int32, key string, r Rect) {

}

func SavePoint(areaModel string, taskType int32, key string, p Point) {

}

func SaveImage(areaModel string, taskType int32, key string) {

}

func VerifyRect(areaModel string, taskType int32, key string, mx, my int) bool {

	return false
}
