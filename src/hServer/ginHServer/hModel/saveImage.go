package hModel

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dzgCap/src/gameAreaModel"
	"dzgCap/src/hServer/common"
	"dzgCap/src/hServer/ginHServer/ginHServer"
)

func init() {
	ginHServer.RegisterRoot("saveImage", saveImage)
}

func saveImage(c *gin.Context) {
	name := "saveImage"

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

	err := gameAreaModel.SaveImage(sKey, taskId, rKey)
	commonHandler(name, c, err)
}
