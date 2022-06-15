package hServer

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/src/gameAreaModel"
)

func init() {
	RegisterHSv("saveRect", saveRect)
}

func saveRect(ctx iris.Context) {
	name := "saveRect"

	sKey := ctx.URLParam(con_Params_AreaKey)
	if len(sKey) == 0 {
		err := fmt.Errorf("need areaKey")
		commonHandler(name, ctx, err)
		return
	}

	rKey, taskId, r, err := getRectParam(ctx)
	if err != nil {
		return
	}

	err = gameAreaModel.SaveRect(sKey, taskId, rKey, r)
	if err != nil {
		commonHandler(name, ctx, err)
		return
	}

	saveImg, err := ctx.URLParamBool(con_Params_SaveImg)
	if err != nil || !saveImg {
		commonHandler(name, ctx, err)
		return
	}

	err = gameAreaModel.SaveImage(sKey, taskId, rKey)
	commonHandler(name, ctx, err)
}
