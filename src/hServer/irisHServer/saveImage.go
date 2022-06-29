package irisHServer

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/src/gameAreaModel"
	"dzgCap/src/hServer/common"
)

func init() {
	RegisterHSv("saveImage", saveImage)
}

func saveImage(ctx iris.Context) {
	name := "saveImage"

	sKey := ctx.URLParam(common.Con_Params_AreaKey)
	if len(sKey) == 0 {
		err := fmt.Errorf("need areaKey")
		common.CommonHandler(name, ctx, err)
		return
	}

	rKey := ctx.URLParam(common.Con_Params_Key)
	if len(rKey) == 0 {
		err := fmt.Errorf("need rect key")
		common.CommonHandler(name, ctx, err)
		return
	}

	taskId32, err := ctx.URLParamInt(common.Con_Params_TaskType)
	if err != nil {
		err := fmt.Errorf("need taskType")
		common.CommonHandler(name, ctx, err)
		return
	}

	taskId := int32(taskId32)

	err = gameAreaModel.SaveImage(sKey, taskId, rKey)
	common.CommonHandler(name, ctx, err)

	return
}
