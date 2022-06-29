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
	ginHServer.RegisterRoot("saveRect", saveRect)
}

func saveRect(c *gin.Context) {
	name := "saveRect"

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

	w, exists := getQueryInt(c, common.Con_Params_W)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need W"))
		return
	}

	h, exists := getQueryInt(c, common.Con_Params_H)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need H"))
		return
	}

	r := model.Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}

	err := gameAreaModel.SaveRect(sKey, taskId, rKey, r)
	if err != nil {
		commonHandler(name, c, err)
	}

	saveImg, exists := getQueryBool(c, common.Con_Params_SaveImg)
	if !exists || !saveImg {
		return
	}

	err = gameAreaModel.SaveImage(sKey, taskId, rKey)
	commonHandler(name, c, err)
}
