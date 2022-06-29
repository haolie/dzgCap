package irisHServer

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/src/gameAreaModel"
	"dzgCap/src/hServer/common"
	"dzgCap/src/model"
)

func init() {
	RegisterHSv("savePoint", savePoint)
}

func savePoint(ctx iris.Context) {
	name := "savePoint"
	sKey := ctx.URLParam(common.Con_Params_AreaKey)
	if len(sKey) == 0 {
		err := fmt.Errorf("need areaKey")
		common.CommonHandler(name, ctx, err)
		return
	}

	taskId32, err := ctx.URLParamInt(common.Con_Params_TaskType)
	if err != nil {
		err := fmt.Errorf("need taskType")
		common.CommonHandler(name, ctx, err)
		return
	}

	pkey := ctx.URLParam(common.Con_Params_Key)
	if len(pkey) == 0 {
		err := fmt.Errorf("need key")
		common.CommonHandler(name, ctx, err)
		return
	}

	taskId := int32(taskId32)

	x, err := ctx.URLParamInt(common.Con_Params_X)
	if err != nil {
		err := fmt.Errorf("need x")
		common.CommonHandler(name, ctx, err)
		return
	}

	y, err := ctx.URLParamInt(common.Con_Params_Y)
	if err != nil {
		err := fmt.Errorf("need y")
		common.CommonHandler(name, ctx, err)
		return
	}

	err = gameAreaModel.SavePoint(sKey, taskId, pkey, model.Point{X: x, Y: y})
	common.CommonHandler(name, ctx, err)

	return
}
