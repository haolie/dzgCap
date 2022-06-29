package hModel

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/gin-gonic/gin"

	"dzgCap/src/gameAreaModel"
	"dzgCap/src/gameCenter"
	"dzgCap/src/hServer/common"
	"dzgCap/src/hServer/ginHServer/ginHServer"
	"dzgCap/src/imageTool"
)

func init() {
	ginHServer.RegisterRoot("saveArea", saveArea)
}

func saveArea(c *gin.Context) {
	name := "saveArea"

	sKey, exists := c.GetQuery(common.Con_Params_AreaKey)
	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("need areaKey"))
		return
	}

	img := imageTool.CapFullScreen()
	r, exists := imageTool.FindMinRect(img, color.RGBA{
		R: 20,
		G: 24,
		B: 31,
		A: 255,
	})

	if !exists {
		c.JSON(http.StatusOK, common.CreateErrHSResponse("not find game in screen"))
		return
	}

	gameCenter.ScanArea()

	err := gameAreaModel.SaveAreaModel(sKey, r)
	if err != nil {
		c.JSON(http.StatusOK, common.CreateErrHSResponse(fmt.Sprintf("saveArea err:%s", err)))
		return
	} else {
		c.JSON(http.StatusOK, common.CreateSuccessHSResponse(name+"success"))
	}
}
