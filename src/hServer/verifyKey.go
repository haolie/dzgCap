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

	"dzgCap/src/ScreenModel"
)

func init() {
	RegisterHSv("verifyKey", verifyKey)
}

func verifyKey(ctx iris.Context) {
	taskId32, err := ctx.URLParamInt("taskId")
	if err != nil {
		ctx.WriteString("need taskId")
		return
	}

	taskId := int32(taskId32)

	rk := ctx.URLParam("key")
	if len(rk) == 0 {
		ctx.WriteString("need key")
		return
	}

	r, exists := ScreenModel.GetCurrentScreenArea().GetRect(taskId, rk)
	if exists {
		ctx.WriteString(fmt.Sprintf("key:%s, rect: %v", rk, r))
		return
	}

	p, exists := ScreenModel.GetCurrentScreenArea().GetPoint(taskId, rk)
	if exists {
		ctx.WriteString(fmt.Sprintf("point:%s, rect: %v", rk, p))
		return
	}

	ctx.WriteString(fmt.Sprintf("key:%s not find !!!", rk))
}
