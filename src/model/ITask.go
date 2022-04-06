package model

import (
	"time"
)

type ITask interface {
	GetKey() TaskEnum
	GetStatus() TaskStatusEnum
	GetStartTime() time.Time
	GetEndTime() time.Time
	Start()
	Stop()
	Pause()
	Release()
}
