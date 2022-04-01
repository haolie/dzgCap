package taskCenter

import (
	"fmt"

	"dzgCap/model"
)

var (
	taskMap = make(map[model.TaskEnum]model.ITask, 4)
)

func RegisterTask(task model.ITask) {
	if _, exists := taskMap[task.GetKey()]; exists {
		panic(fmt.Sprintf("can not register again :%v", task.GetKey()))
	}

	taskMap[task.GetKey()] = task
}

func GetTask(taskType model.TaskEnum) (task model.ITask, exists bool) {
	task, exists = taskMap[taskType]
	return
}
