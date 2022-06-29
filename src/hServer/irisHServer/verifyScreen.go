// ************************************
// @package: hServer
// @description: 验证游戏区域模型
// @author:
// @revision history:
// @create date: 2022-04-22 11:34:22
// ************************************
package irisHServer

import (
	"github.com/kataras/iris/v12"
)

func init() {
	RegisterHSv("verifyScreen", verifyScreen)
}

func verifyScreen(ctx iris.Context) {
	//sKey := ctx.URLParam("screen")//
	//if len(sKey) == 0 {
	//	ctx.WriteString("need key")
	//	return
	//}
	//
	//taskId32, err := ctx.URLParamInt("taskId")
	//if err != nil {
	//	ctx.WriteString("need taskId")
	//	return
	//}
	//
	//taskId := int32(taskId32)
	//
	//success, errList := ScreenModel.VerifyTask(sKey, taskId)
	//if len(errList) > 0 {
	//	ctx.WriteString(fmt.Sprintf("%v", errList))
	//	return
	//}
	//
	//ctx.WriteString(fmt.Sprintf("%v", success))

}
