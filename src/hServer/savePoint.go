package hServer

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/src/gameAreaModel"
	"dzgCap/src/model"
)

func init() {
	RegisterHSv("savePoint", savePoint)
}

func savePoint(ctx iris.Context) {
	name := "saveRect"
	sKey := ctx.URLParam(con_Params_AreaKey)
	if len(sKey) == 0 {
		err := fmt.Errorf("need areaKey")
		commonHandler(name, ctx, err)
		return
	}

	taskId32, err := ctx.URLParamInt(con_Params_TaskType)
	if err != nil {
		err := fmt.Errorf("need taskType")
		commonHandler(name, ctx, err)
		return
	}

	pkey := ctx.URLParam(con_Params_Key)
	if len(pkey) == 0 {
		err := fmt.Errorf("need key")
		commonHandler(name, ctx, err)
		return
	}

	taskId := int32(taskId32)

	x, err := ctx.URLParamInt(con_Params_X)
	if err != nil {
		err := fmt.Errorf("need x")
		commonHandler(name, ctx, err)
		return
	}

	y, err := ctx.URLParamInt(con_Params_Y)
	if err != nil {
		err := fmt.Errorf("need y")
		commonHandler(name, ctx, err)
		return
	}

	err = gameAreaModel.SavePoint(sKey, taskId, pkey, model.Point{X: x, Y: y})
	commonHandler(name, ctx, err)

	return
}
