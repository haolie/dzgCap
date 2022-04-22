package model

import (
	"time"
)

type ITask interface {
	// GetKey
    // @description: 返回任务Id
    // parameter:
    // return:
    //		@TaskEnum:
	GetKey() TaskEnum

	// GetStatus
    // @description: 返回任务状态
    // parameter:
    // return:
    //		@TaskStatusEnum:
	GetStatus() TaskStatusEnum

	// GetStartTime
    // @description: 返回任务开始时间
    // parameter:
    // return:
    //		@time.Time:
	GetStartTime() time.Time

	// GetEndTime
    // @description: 返回任务结束时间
    // parameter:
    // return:
    //		@time.Time:
	GetEndTime() time.Time

	// Start
    // @description: 开始任务
    // parameter:
    // return:
	Start(param interface{})

	// Stop
    // @description: 结束任务
    // parameter:
    // return:
	Stop()

	// Pause
    // @description: 暂停任务
    // parameter:
    // return:
	Pause()

	// Release
    // @description: 重置任务
    // parameter:
    // return:
	Release()
}
