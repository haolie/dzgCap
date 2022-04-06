package taskCenter

import (
	"fmt"

	"dzgCap/src/ScreenModel"
	"dzgCap/src/model"
)

var (
	taskMap     = make(map[model.TaskEnum]model.ITask, 4)
	currentTask model.ITask
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

func CurrentTask() (task model.ITask, exists bool) {
	return currentTask, currentTask != nil
}

func StartTask(taskType model.TaskEnum) error {
	task, exists := taskMap[taskType]
	if !exists {
		return fmt.Errorf("not find Task")
	}

	_, errList := ScreenModel.BaseVerify(ScreenModel.GetCurrentModelKey())
	if len(errList) > 0 {
		fmt.Println(errList)
		return fmt.Errorf(" please config ")
	}

	_, errList = ScreenModel.VerifyTask(ScreenModel.GetCurrentModelKey(), int32(task.GetKey()))
	if len(errList) > 0 {
		fmt.Println(errList)
		return fmt.Errorf(" please config ")
	}

	go task.Start()
	currentTask = task

	return nil
}
