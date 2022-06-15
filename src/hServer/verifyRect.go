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

	"dzgCap/src/gameCenter"
	"dzgCap/src/model"
)

func init() {
	RegisterHSv("verifyRect", verifyRect)
}

func verifyRect(ctx iris.Context) {
	name := "verifyRect"

	taskId32, err := ctx.URLParamInt(con_Params_TaskType)
	if err != nil {
		commonHandler(name, ctx, fmt.Errorf("need taskId"))
		return
	}

	taskId := int32(taskId32)

	rk := ctx.URLParam(con_Params_Key)
	if len(rk) == 0 {
		commonHandler(name, ctx, fmt.Errorf("need rect"))
		return
	}

	success := gameCenter.VerifyRect(model.TaskEnum(taskId), rk)
	if !success {
		err = fmt.Errorf("verify Fail")
	}

	commonHandler(name, ctx, err)
}
