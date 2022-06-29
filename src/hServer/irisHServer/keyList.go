// ************************************
// @package: hServer
// @description: 验证区域是否匹配
// @author:
// @revision history:
// @create date: 2022-04-22 11:33:44
// ************************************
package irisHServer

import (
	"github.com/kataras/iris/v12"

	"dzgCap/src/hServer/common"
	"dzgCap/src/model"
)

func init() {
	RegisterHSv("keyList", keyList)
}

func keyList(ctx iris.Context) {
	//name := "keyList"

	dataMap := common.CreateSuccessHSResponse("")
	dataMap["keyList"] = model.GetKeyMap()
	ctx.JSON(common.CreateSuccessHSResponse(dataMap))
}
