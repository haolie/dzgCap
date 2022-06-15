package hServer

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/src/gameAreaModel"
)

func init() {
	RegisterHSv("saveImage", saveImage)
}

func saveImage(ctx iris.Context) {
	name := "saveImage"

	sKey := ctx.URLParam(con_Params_AreaKey)
	if len(sKey) == 0 {
		err := fmt.Errorf("need areaKey")
		commonHandler(name, ctx, err)
		return
	}

	rKey := ctx.URLParam(con_Params_Key)
	if len(rKey) == 0 {
		err := fmt.Errorf("need rect key")
		commonHandler(name, ctx, err)
		return
	}

	taskId32, err := ctx.URLParamInt(con_Params_TaskType)
	if err != nil {
		err := fmt.Errorf("need taskType")
		commonHandler(name, ctx, err)
		return
	}

	taskId := int32(taskId32)

	err = gameAreaModel.SaveImage(sKey, taskId, rKey)
	commonHandler(name, ctx, err)

	return
}
