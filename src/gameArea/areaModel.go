package gameArea

import (
	"context"
	"sync"

	"dzgCap/src/gameAreaModel"
	. "dzgCap/src/model"
	"dzgCap/src/task/taskCenter"
)

var _ IGameArea = (*GameArea)(nil)

type GameArea struct {
	gameModelKey string
	rect         Rect
	task         ITask
	gTable       IGameTable
	startOnce    *sync.Once
}

func (ga *GameArea) GetKey() string {
	return ga.gameModelKey
}

func (ga *GameArea) ClickRectKey(key string) {

}

func (ga *GameArea) ClickPointKey(key string) {

}

func (ga *GameArea) ClickPoint(x, y int) {

}

func (ga *GameArea) VerifyRect(key string) bool {
	return gameAreaModel.VerifyRect(ga.gameModelKey, int32(ga.task.GetTaskType()), key)
}

func (ga *GameArea) StartTask(ctx context.Context, taskType TaskEnum) error {
	ga.Stop()

	ga.task = taskCenter.CreateTask(taskType, ga)
	ga.task.Start(ctx, GetMinTime(), GetMaxTime())

	return nil
}

func (ga *GameArea) Stop() {
	if ga.task != nil {
		ga.task.Stop()
		ga.task = nil
	}

}

func (ga *GameArea) GetStatus() (status TaskStatusEnum, taskType TaskEnum) {
	if ga.task == nil {
		status = TaskStatusEnum_Unstart
		return
	}

	return ga.task.GetStatus(), ga.task.GetTaskType()
}
