package hModel

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dzgCap/src/gameAreaModel"
	"dzgCap/src/hServer/common"
	"dzgCap/src/hServer/ginHServer/ginHServer"
	"dzgCap/src/model"
)

func init() {
	ginHServer.RegisterRoot("savePoint", savePoint)
}

func savePoint(c *gin.Context) {
	name := "savePoint"

	sKey, exists := c.GetQuery(common.Con_Params_AreaKey)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need areaKey"))
		return
	}

	pKey, exists := c.GetQuery(common.Con_Params_Key)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need rectKey"))
		return
	}

	taskId, exists := getQueryInt32(c, common.Con_Params_TaskType)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need taskId"))
		return
	}

	x, exists := getQueryInt(c, common.Con_Params_X)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need x"))
		return
	}

	y, exists := getQueryInt(c, common.Con_Params_Y)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need y"))
		return
	}

	err := gameAreaModel.SavePoint(sKey, taskId, pKey, model.Point{X: x, Y: y})
	commonHandler(name, c, err)
}
