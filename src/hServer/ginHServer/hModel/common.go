package hModel

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"dzgCap/src/hServer/common"
)

func getQueryInt32(c *gin.Context, key string) (q int32, exists bool) {
	qs, exists := c.GetQuery(key)
	if !exists {
		return
	}

	temp, err := strconv.Atoi(qs)
	if err != nil {
		exists = false
		return
	}

	return int32(temp), true
}

func getQueryInt(c *gin.Context, key string) (q int, exists bool) {
	qs, exists := c.GetQuery(key)
	if !exists {
		return
	}

	temp, err := strconv.Atoi(qs)
	if err != nil {
		exists = false
		return
	}

	return temp, true
}

func getQueryBool(c *gin.Context, key string) (q bool, exists bool) {
	qs, exists := c.GetQuery(key)
	if !exists {
		return
	}

	q = qs == "true"

	return
}

func commonHandler(name string, c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusOK, common.CreateErrHSResponse(fmt.Sprintf("%s Fail:%s", name, err)))
		return
	} else {
		c.JSON(http.StatusOK, common.CreateSuccessHSResponse(name+"success"))
	}
}
