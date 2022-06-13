package main

import (
	"context"
	"fmt"

	"dzgCap/dzgCap"
	_ "dzgCap/src"
)

type RegisterFun func(ctx context.Context) []error

var (
	loadMap  = make(map[string]RegisterFun, 4)
	startMap = make(map[string]RegisterFun, 4)
	existsFn func()
	mainCtx  context.Context
)

func RegisterLoad(key string, fn RegisterFun) {
	if _, exists := loadMap[key]; exists {
		panic(fmt.Sprintf("register Again with key:%s", key))
	}

	loadMap[key] = fn
}

func RegisterStart(key string, fn RegisterFun) {
	if _, exists := startMap[key]; exists {
		panic(fmt.Sprintf("register Again with key:%s", key))
	}

	startMap[key] = fn
}

func main() {
	dzgCap.Start()
}

//
//func RegisterKey() {
//	go registerKeyPause()
//}
//
//func registerKeyPause() {
//	ok := robotgo.AddEvents("p", "ctrl")
//	if ok {
//		task, exists := taskCenter.CurrentTask()
//		if exists {
//			task.Pause()
//		}
//
//		fmt.Println("key Pause")
//	}
//
//	go registerKeyGo()
//
//}
//
//func registerKeyGo() {
//	ok := robotgo.AddEvents("g", "ctrl")
//	if ok {
//		err := ScreenModel.GetCurrentScreenArea().FreshArea()
//		if err != nil {
//			panic(err)
//		}
//
//		task, exists := taskCenter.CurrentTask()
//		if exists {
//			task.Start(50)
//		}
//
//		fmt.Println("key Go on")
//	}
//
//	go registerKeyPause()
//}
