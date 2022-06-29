package hModel

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dzgCap/src/hServer/common"
	"dzgCap/src/hServer/ginHServer/ginHServer"
	"dzgCap/src/model"
)

func init() {
	ginHServer.RegisterRoot("keyList", keyList)
}

func keyList(c *gin.Context) {
	dataMap := common.CreateSuccessHSResponse("")
	dataMap["keyList"] = model.GetKeyMap()
	c.JSON(http.StatusOK, common.CreateSuccessHSResponse(dataMap))
}
