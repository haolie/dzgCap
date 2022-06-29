// ************************************
// @package: hServer
// @description: 任务控制
// @author:
// @revision history:
// @create date: 2022-04-22 11:33:44
// ************************************
package irisHServer

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"

	"dzgCap/src/gameCenter"
	"dzgCap/src/hServer/common"
	"dzgCap/src/model"
)

func init() {
	RegisterHSv("taskCtl", taskCtl)
}

func taskCtl(ctx iris.Context) {
	name := "taskCtl"
	taskId32, err := ctx.URLParamInt(common.Con_Params_TaskType)
	if err != nil {
		ctx.WriteString("need taskId")
		return
	}

	taskId := int32(taskId32)

	cmd := ctx.URLParam(common.Con_Params_Cmd)
	if len(cmd) == 0 {
		ctx.WriteString("need cmd")
		return
	}

	if cmd == "start" {
		gameCenter.StartTask(model.TaskEnum(taskId))
		common.CommonHandler(name, ctx, nil)
	} else {
		gameCenter.Stop()
		common.CommonHandler(name, ctx, nil)
	}

	ctx.WriteString(fmt.Sprintf("%v, success", time.Now()))
}
