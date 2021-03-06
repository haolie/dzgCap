package main

import (
	"dzgCap/ConfigManger"
	"fmt"

	"github.com/go-vgo/robotgo"

	_ "dzgCap/src"
	"dzgCap/src/ScreenModel"
	"dzgCap/src/hServer"
	"dzgCap/src/task/taskCenter"
)

func main() {

	err := ConfigManger.Load("")
	if err != nil {
		panic(err)
	}

	ScreenModel.LoadScreenArea(ConfigManger.GetScreenKey())

	go hServer.StartHServer()

	//err = taskCenter.StartTask(model.TaskEnum_Meeting, 1)
	//if err != nil {
	//	panic(err)
	//}

	//RegisterKey()
	select {}
}

func RegisterKey() {
	go registerKeyPause()
}

func registerKeyPause() {
	ok := robotgo.AddEvents("p", "ctrl")
	if ok {
		task, exists := taskCenter.CurrentTask()
		if exists {
			task.Pause()
		}

		fmt.Println("key Pause")
	}

	go registerKeyGo()

}

func registerKeyGo() {
	ok := robotgo.AddEvents("g", "ctrl")
	if ok {
		err := ScreenModel.GetCurrentScreenArea().FreshArea()
		if err != nil {
			panic(err)
		}

		task, exists := taskCenter.CurrentTask()
		if exists {
			task.Start(50)
		}

		fmt.Println("key Go on")
	}

	go registerKeyPause()
}
