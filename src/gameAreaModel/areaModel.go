package gameAreaModel

import (
	"image"

	"dzgCap/src/model"
)

type areaTaskModel struct {
	RectMap  map[string]model.Rect
	PointMap map[string]model.Point
	imgMap   map[string]image.Image
}

func createAreaTaskModel() *areaTaskModel {
	return &areaTaskModel{
		RectMap:  make(map[string]model.Rect, 4),
		PointMap: make(map[string]model.Point, 4),
		imgMap:   make(map[string]image.Image, 4),
	}
}

type areaModel struct {
	Key     string
	TaskMap map[int32]*areaTaskModel
}

func createNewArea(key string) *areaModel {
	return &areaModel{
		Key:     key,
		TaskMap: make(map[int32]*areaTaskModel, 2),
	}
}
