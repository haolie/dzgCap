// ************************************
// @package: hServer
// @description: 任务控制
// @author:
// @revision history:
// @create date: 2022-04-22 11:33:44
// ************************************
package hServer

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"

	"dzgCap/src/model"
	"dzgCap/src/task/taskCenter"
)

func init() {
	RegisterHSv("taskCtl", taskCtl)
}

func taskCtl(ctx iris.Context) {
	taskId32, err := ctx.URLParamInt("taskId")
	if err != nil {
		ctx.WriteString("need taskId")
		return
	}

	taskId := int32(taskId32)

	cmd := ctx.URLParam("cmd")
	if len(cmd) == 0 {
		ctx.WriteString("need cmd")
		return
	}

	param:= ctx.URLParam("param")

	if cmd =="start"{
		err = taskCenter.StartTask(model.TaskEnum(taskId),param)
		if err != nil {
			ctx.WriteString(fmt.Sprintf("%s",err))
		}
	}else {
		t,exists := taskCenter.CurrentTask()
		if exists {
			t.Stop()
		}
	}

	ctx.WriteString(fmt.Sprintf("%v, success", time.Now()))
}
