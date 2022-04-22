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
	RegisterHSv("verifyRect", verifyRect)
}

func verifyRect(ctx iris.Context) {
	taskId32, err := ctx.URLParamInt("taskId")
	if err != nil {
		ctx.WriteString("need taskId")
		return
	}

	taskId := int32(taskId32)

	rk := ctx.URLParam("rect")
	if len(rk) == 0 {
		ctx.WriteString("need rect")
		return
	}

	success, err := ScreenModel.GetCurrentScreenArea().CompareRectToCash(taskId, rk)
	if err != nil {
		ctx.WriteString(fmt.Sprintf("%v", err))
		return
	}

	ctx.WriteString(fmt.Sprintf("%v", success))
}
