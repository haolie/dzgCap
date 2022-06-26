package dzgCap

import (
	"context"
	"fmt"

	"dzgCap/ConfigManger"
	"dzgCap/Loger"
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

func Start() {


	mainCtx, existsFn = context.WithCancel(context.Background())

	errList :=ConfigManger.LoadHandler(mainCtx)
	if len(errList) > 0 {
		printErrList(errList)
		panic("load fail")
	}

	errList = load(mainCtx)
	if len(errList) > 0 {
		printErrList(errList)
		panic("load fail")
	}

	errList = start(mainCtx)
	if len(errList) > 0 {
		printErrList(errList)
		panic("load fail")
	}

	<-mainCtx.Done()

	//err := ConfigManger.Load("")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = gameCenter.StartWork()
	//if err != nil {
	//	panic(err)
	//}
	//
	//go hServer.StartHServer()
	//
	////err = taskCenter.StartTask(model.TaskEnum_Meeting, 1)
	////if err != nil {
	////	panic(err)
	////}
	//
	////RegisterKey()
	//select {}
}

func start(ctx context.Context) (errList []error) {
	for _, fn := range startMap {
		tempList := fn(ctx)
		errList = append(errList, tempList...)
	}

	return
}

func load(ctx context.Context) (errList []error) {
	for _, fn := range loadMap {
		tempList := fn(ctx)
		errList = append(errList, tempList...)
	}

	return
}

func printErrList(errList []error) {
	for _, err := range errList {
		Loger.LogErr(fmt.Sprintf("%s", err))
	}
}
