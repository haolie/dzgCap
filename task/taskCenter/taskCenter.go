package taskCenter

import (
	"fmt"

	"dzgCap/ScreenModel"
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

func StartTask(taskType model.TaskEnum) {
	task, exists := taskMap[taskType]
	if !exists {
		panic("not find Task")
	}

	_, errList := ScreenModel.BaseVerify(model.Sys_Con_Model_Base)
	if len(errList) > 0 {
		fmt.Println(errList)
		panic(" please config ")
	}

	_, errList = ScreenModel.VerifyTask(model.Sys_Con_Model_Base, int32(task.GetKey()))
	if len(errList) > 0 {
		fmt.Println(errList)
		panic(" please config ")
	}

	task.Start()
}
