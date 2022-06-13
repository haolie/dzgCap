package gameAreaModel

import (
	"image"

	. "dzgCap/src/model"
)

func GetModelKeyWithRect(r Rect) (modelKey string, exists bool) {
	for k, mod := range modelMap {
		if _, exists := mod.TaskMap[0]; !exists {
			continue
		}

		if mod.TaskMap[0].RectMap[Sys_Key_Rect_Game] == r {
			return k, true
		}
	}

	return
}

func GetRect(areaModel string, taskType int32, key string) (r Rect, exists bool) {
	areaTaskModelObj, exists := getAreaTaskModel(areaModel, taskType)
	if !exists {
		return
	}

	r, exists = areaTaskModelObj.RectMap[key]

	return
}

func GetPoint(areaModel string, taskType int32, key string) (p Point, exists bool) {
	areaTaskModelObj, exists := getAreaTaskModel(areaModel, taskType)
	if !exists {
		return
	}

	p, exists = areaTaskModelObj.PointMap[key]

	return
}

func GetImage(areaModel string, taskType int32, key string) (img image.Image, exists bool) {
	areaTaskModelObj, exists := getAreaTaskModel(areaModel, taskType)
	if !exists {
		return
	}

	if areaTaskModelObj.imgMap == nil {
		return
	}

	img, exists = areaTaskModelObj.imgMap[key]

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
