package gameCenter

import (
	"context"
	"fmt"
	"sync"

	"dzgCap/Loger"
	"dzgCap/dzgCap"
	"dzgCap/src/gameArea"
	"dzgCap/src/gameAreaModel"
	"dzgCap/src/imageTool"
	"dzgCap/src/model"
)

var (
	areaList   []model.IGameArea
	cancelFun  context.CancelFunc
	gTable     model.IGameTable = newGTable()
	startOnce  sync.Once
	moduleName = "gameCenter"
)

func init() {
	dzgCap.RegisterStart(moduleName, startHandler)
}

func startHandler(ctx context.Context) (errList []error) {
	err := StartWork()
	if err != nil {
		errList = append(errList, err)
	}

	return
}

func ScanArea() error {
	areaList = make([]model.IGameArea, 0, 4)

	img := imageTool.CapFullScreen()
	rectList, exists := imageTool.FindRect(img, model.GetLDColor())
	if !exists {
		return nil
	}

	var modelKey string
	var areaFactory func(r model.Rect) model.IGameArea

	for _, r := range rectList {
		k, exists := gameAreaModel.GetModelKeyWithRect(r)
		if !exists {
			continue
		}

		if modelKey != k {
			modelKey = k
			areaFactory = gameArea.CreateFactory(k, gTable)
		}

		areaList = append(areaList, areaFactory(r))
	}

	Loger.LogInfo(fmt.Sprintf("find %d ScanArea", len(areaList)))

	return nil
}

func StartTask(taskType model.TaskEnum) {
	ScanArea()

	if len(areaList) == 0 {
		return
	}

	Stop()

	var ctx context.Context
	ctx, cancelFun = context.WithCancel(context.Background())

	for _, item := range areaList {
		err := item.StartTask(ctx, taskType)
		if err != nil {
			Loger.LogErr(fmt.Sprintf("任务开始失败 err:%v", err))
		}
	}
}

func StartWork() error {
	(&startOnce).Do(func() {
		ScanArea()
	})

	return nil
}

func Stop() {
	if cancelFun != nil {
		cancelFun()
		cancelFun = nil
	}
}

func VerifyRect(taskType model.TaskEnum, key string) bool {
	if len(areaList) == 0 {
		return false
	}

	return areaList[0].VerifyRect(taskType, key)
}

func IsMainView() bool {
	if len(areaList) == 0 {
		return false
	}

	return areaList[0].IsHome()
}
