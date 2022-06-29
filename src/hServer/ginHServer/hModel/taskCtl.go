package hModel

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"dzgCap/src/gameCenter"
	"dzgCap/src/hServer/common"
	"dzgCap/src/hServer/ginHServer/ginHServer"
	"dzgCap/src/model"
)

func init() {
	ginHServer.RegisterRoot("taskCtl", taskCtl)
}

func taskCtl(c *gin.Context) {
	name := "taskCtl"

	taskId, exists := getQueryInt32(c, common.Con_Params_TaskType)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need taskId"))
		return
	}

	exists = model.TaskEnumVerify(taskId)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse(fmt.Sprintf("not exists taskId:%d", taskId)))
		return
	}

	cmd, exists := c.GetQuery(common.Con_Params_Cmd)
	if !exists {
		cmd = "start"
	}

	if cmd == "start" {
		gameCenter.StartTask(model.TaskEnum(taskId))
	} else {
		gameCenter.Stop()
	}

	commonHandler(name, c, nil)
}
