package model

type TaskStatusEnum int32

const (
	// 未开始
	TaskStatusEnum_Unstart TaskStatusEnum = iota
	// 正在执行
	TaskStatusEnum_Runing
	// 暂停中
	TaskStatusEnum_Pause
)
