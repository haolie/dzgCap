// ************************************
// @package: hServer
// @description: 验证区域是否匹配
// @author:
// @revision history:
// @create date: 2022-04-22 11:33:44
// ************************************
package hServer

import (
	"github.com/kataras/iris/v12"

	"dzgCap/src/model"
)

func init() {
	RegisterHSv("keyList", keyList)
}

func keyList(ctx iris.Context) {
	//name := "keyList"

	dataMap := createSuccessHSResponse("")
	dataMap["keyList"] = model.GetKeyMap()
	ctx.JSON(createSuccessHSResponse(dataMap))
}
