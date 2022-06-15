// ************************************
// @package: hServer
// @description: 验证区域是否匹配
// @author:
// @revision history:
// @create date: 2022-04-22 11:33:44
// ************************************
package hServer

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/src/gameAreaModel"
)

func init() {
	RegisterHSv("verifyKey", verifyKey)
}

func verifyKey(ctx iris.Context) {
	name := "taskCtl"
	taskId32, err := ctx.URLParamInt(con_Params_TaskType)
	if err != nil {
		ctx.WriteString("need taskId")
		return
	}

	taskId := int32(taskId32)

	sKey := ctx.URLParam(con_Params_AreaKey)
	if len(sKey) == 0 {
		commonHandler(name, ctx, fmt.Errorf("need screen key"))
		return
	}

	rk := ctx.URLParam(con_Params_Key)
	if len(rk) == 0 {
		commonHandler(name, ctx, fmt.Errorf("need rect key"))
		return
	}

	r, exists := gameAreaModel.GetRect(sKey, taskId, rk)
	if exists {
		ctx.JSON(createSuccessHSResponse(fmt.Sprintf("key:%s, rect: %v", rk, r)))
		return
	}

	p, exists := gameAreaModel.GetPoint(sKey, taskId, rk)
	if exists {
		ctx.JSON(createSuccessHSResponse(fmt.Sprintf("point:%s, rect: %v", rk, p)))
		return
	}

	commonHandler(name, ctx, fmt.Errorf("key:%s not find", rk))
}
