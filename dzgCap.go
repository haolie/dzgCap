package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"

	"dzgCap/ConfigManger"
	_ "dzgCap/src"
	"dzgCap/src/ScreenModel"
	"dzgCap/src/model"
	"dzgCap/src/task/taskCenter"
)

func main() {
	//fmt.Println(2345)
	//robotgo.Move(2382, 1084)
	//robotgo.Click("left", true)
	//ShowMessage2("","")

	err := ConfigManger.Load("")
	if err != nil {
		panic(err)
	}

	err = ScreenModel.SetScreenModel(ConfigManger.GetScreenKey())
	if err != nil {
		panic(err)
	}

	//err = hServer.StartHServer()
	//if err != nil {
	//	panic(err)
	//}

	err = taskCenter.StartTask(model.TaskEnum_Meeting)
	if err != nil {
		panic(err)
	}

	RegisterKey()
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
		task, exists := taskCenter.CurrentTask()
		if exists {
			task.Start()
		}

		fmt.Println("key Go on")
	}

	go registerKeyPause()
}
