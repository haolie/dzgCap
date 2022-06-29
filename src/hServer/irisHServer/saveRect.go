package irisHServer

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/src/gameAreaModel"
	"dzgCap/src/hServer/common"
)

func init() {
	RegisterHSv("saveRect", saveRect)
}

func saveRect(ctx iris.Context) {
	name := "saveRect"

	sKey := ctx.URLParam(common.Con_Params_AreaKey)
	if len(sKey) == 0 {
		err := fmt.Errorf("need areaKey")
		common.CommonHandler(name, ctx, err)
		return
	}

	rKey, taskId, r, err := getRectParam(ctx)
	if err != nil {
		return
	}

	err = gameAreaModel.SaveRect(sKey, taskId, rKey, r)
	if err != nil {
		common.CommonHandler(name, ctx, err)
		return
	}

	saveImg, err := ctx.URLParamBool(common.Con_Params_SaveImg)
	if err != nil || !saveImg {
		common.CommonHandler(name, ctx, err)
		return
	}

	err = gameAreaModel.SaveImage(sKey, taskId, rKey)
	common.CommonHandler(name, ctx, err)
}
