package ginHServer

import (
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"dzgCap/dzgCap"
)

var (
	startOnce  = new(sync.Once)
	rootMap    = make(map[string]func(c *gin.Context), 8)
	moduleName = "ginHServer"
)

func init() {
	dzgCap.RegisterStart(moduleName, start)
}

func RegisterRoot(key string, handler func(c *gin.Context)) {
	rootMap[key] = handler
}

func start(ctx context.Context) (errList []error) {
	startOnce.Do(func() {
		engine := gin.Default()
		engine.StaticFS("/view", http.Dir("./view"))
		engine.StaticFS("/static", http.Dir("./view/static"))

		for key, handler := range rootMap {
			str := "/express/w/" + key
			engine.GET(str, handler)
		}

		engine.Run(":9090")
	})

	return
}
