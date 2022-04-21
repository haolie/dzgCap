package hServer

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"

	"dzgCap/src/ScreenModel"
)

func init() {
	RegisterHSv("saveRectImage", saveRectImage)
}

func saveRectImage(ctx iris.Context) {
	sKey := ctx.URLParam("screen")
	if len(sKey) == 0 {
		ctx.WriteString("need key")
		return
	}

	rkey, taskId, r, err := getRectParam(ctx)
	if err != nil {
		return
	}

	sr, exists := ScreenModel.GetScreenArea(sKey)
	if !exists {
		ctx.WriteString("not find screen:" + sKey)
		return
	}

	sr.AddRect(taskId, rkey, r)

	err = sr.SaveRectImg(taskId, rkey)
	if err != nil {
		ctx.WriteString(fmt.Sprintf("%s", err))
		return
	}

	err = ScreenModel.SaveScreenModel(sr)
	if err != nil {
		ctx.WriteString(fmt.Sprintf("%s", err))
		return
	}

	ctx.WriteString(successStr("saveRectImage"))
}

func successStr(name string) string {
	return fmt.Sprintf("%v:%s", time.Now(), name)
}
