package ScreenModel

import (
	"dzgCap/src/model"
)

type PointModel struct {
	Key string
	X   int
	Y   int
}

type RectModel struct {
	Key string
	model.Rect
}

type TaskSaveModel struct {
	PointList []PointModel
	RectList  []RectModel
}

func NewTaskSaveModel() *TaskSaveModel {
	return &TaskSaveModel{
		PointList: nil,
		RectList:  nil,
	}
}
