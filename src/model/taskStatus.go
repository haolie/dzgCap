package model

type TaskStatusEnum int32

const (
	TaskStatusEnum_Unstart TaskStatusEnum = iota
	TaskStatusEnum_Runing
	TaskStatusEnum_Pause
)
