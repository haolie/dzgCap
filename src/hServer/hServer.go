package hServer

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/ConfigManger"
)

var (
	svMap = make(map[string]func(ctx iris.Context), 8)
)

func RegisterHSv(key string, fn func(ctx iris.Context)) {
	if _, exists := svMap[key]; exists {
		panic("register again Hsv:" + key)
	}

	svMap[key] = fn
}

func StartHServer() error {
	app := iris.New()

	for key, fn := range svMap {
		app.Get("/w/"+key, fn)
	}

	return app.Run(iris.Addr(fmt.Sprintf(":%d", ConfigManger.GetHSPort())))
}
