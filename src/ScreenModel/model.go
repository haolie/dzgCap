package ScreenModel

import (
	"dzgCap/src/model"
)

// 点模型
type PointModel struct {
	Key string // 键值
	X   int
	Y   int
}

// 区域模型
type RectModel struct {
	Key string // 键
	model.Rect
}

// 任务存储模型
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
