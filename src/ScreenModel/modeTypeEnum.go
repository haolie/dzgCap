package ScreenModel

type ModelTypeEnum int32

const (
	// 模型类型-点
	ModelTypeEnum_Point ModelTypeEnum = iota + 1
	// 模型类型-区域
	ModelTypeEnum_Rect
	// 模型类型-图片
	ModelTypeEnum_Image
)
