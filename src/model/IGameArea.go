package model

import (
	"context"
)

type IGameArea interface {
	// GetKey
	// @description: getKey
	// parameter:
	// return:
	//		@string: gameAreaKey
	GetKey() string

	// ClickRectKey
	// @description: 点击区域（key）
	// parameter:
	//		@key: 区域key
	// return:
	ClickRectKey(key string)

	// ClickPointKey
	// @description: 点击
	// parameter:
	//		@key: key
	// return:
	ClickPointKey(key string)

	// ClickPoint
	// @description: 点击指定位置
	// parameter:
	//		@x: x
	//		@y: y
	// return:
	ClickPoint(x, y int)

	// VerifyRect
	// @description: 界面区域验证
	// parameter:
	//		@key: rect key
	// return:
	//		@bool: 验证是否通过
	VerifyRect(key string) bool

	// StartTask
	// @description:
	// parameter:
	//		@ctx:
	//		@taskType:
	// return:
	//		@error:
	StartTask(ctx context.Context, taskType TaskEnum) error

	// Stop
	// @description:
	// parameter:
	// return:
	//		@error:
	Stop()

	// GetStatus
	// @description:
	// parameter:
	// return:
	//		@status:
	//		@taskType:
	GetStatus() (status TaskStatusEnum, taskType TaskEnum)
}
