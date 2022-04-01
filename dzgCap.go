package main

import (
	"dzgCap/model"
	"dzgCap/task/taskCenter"
)

func main() {
	//fmt.Println(2345)
	//robotgo.Move(2382, 1084)
	//robotgo.Click("left", true)
	//ShowMessage2("","")

	task, exists := taskCenter.GetTask(model.TaskEnum_Meeting)
	if exists {
		go task.Start()
		select {}
	}
}
