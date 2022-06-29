package hModel

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"dzgCap/src/gameAreaModel"
	"dzgCap/src/gameCenter"
	"dzgCap/src/hServer/common"
	"dzgCap/src/hServer/ginHServer/ginHServer"
	"dzgCap/src/model"
)

func init() {
	ginHServer.RegisterRoot("verifyKey", verifyKey)
}

func verifyKey(c *gin.Context) {
	name := "verifyKey"

	sKey, exists := c.GetQuery(common.Con_Params_AreaKey)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need areaKey"))
		return
	}

	rKey, exists := c.GetQuery(common.Con_Params_Key)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need rectKey"))
		return
	}

	taskId, exists := getQueryInt32(c, common.Con_Params_TaskType)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need taskId"))
		return
	}

	r, exists := gameAreaModel.GetRect(sKey, taskId, rKey)
	if exists {
		dataMap := common.CreateSuccessHSResponse(fmt.Sprintf("key:%s, rect: %v", rKey, r))
		success := gameCenter.VerifyRect(model.TaskEnum(taskId), rKey)
		if success {
			dataMap["VerifyRect"] = "Success"
		} else {
			dataMap["VerifyRect"] = "Fail"
		}

		c.JSON(http.StatusOK, dataMap)

		return
	}

	p, exists := gameAreaModel.GetPoint(sKey, taskId, rKey)
	if exists {
		c.JSON(http.StatusOK, common.CreateSuccessHSResponse(fmt.Sprintf("point:%s, rect: %v", rKey, p)))
		return
	}

	commonHandler(name, c, fmt.Errorf("key:%s not find", rKey))
}
