package taskCenter

import (
	"fmt"

	. "dzgCap/src/model"
)

var (
	createrMap = make(map[TaskEnum]func(ga IGameArea) ITask, 4)
)

func RegisterTask(taskType TaskEnum, creater func(ga IGameArea) ITask) {
	createrMap[taskType] = creater
}

func CreateTask(taskType TaskEnum, ga IGameArea) (t ITask) {
	c, exists := createrMap[taskType]
	if !exists {
		panic(fmt.Sprintf("not find taskType :%d", taskType))
	}

	return c(ga)
}
