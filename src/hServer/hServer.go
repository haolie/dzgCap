package hServer

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/ConfigManger"
)

var (
	svMap = make(map[string]func(ctx iris.Context), 8)
)

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
