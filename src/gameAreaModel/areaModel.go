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

type areaModel struct {
	TaskMap map[int32]areaTaskModel
}
