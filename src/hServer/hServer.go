package hServer

import (
	"context"
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/ConfigManger"
	"dzgCap/dzgCap"
)

var (
	svMap      = make(map[string]func(ctx iris.Context), 8)
	moduleName = "hServer"
)

func init() {
	dzgCap.RegisterStart(moduleName, startHandler)
}

func startHandler(ctx context.Context) (errList []error) {
	err := StartHServer()
	if err != nil {
		errList = append(errList, err)
	}

	return
}

// 注册http接口
func RegisterHSv(key string, fn func(ctx iris.Context)) {
	if _, exists := svMap[key]; exists {
		panic("register again Hsv:" + key)
	}

	svMap[key] = fn
}

// StartHServer
// @description: 开始http服务
// parameter:
// return:
//		@error:
func StartHServer() error {
	app := iris.New()

	for key, fn := range svMap {
		app.Get("/w/"+key, fn)
	}

	return app.Run(iris.Addr(fmt.Sprintf(":%d", ConfigManger.GetHSPort())))
}
