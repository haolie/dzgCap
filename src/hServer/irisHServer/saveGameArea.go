package irisHServer

import (
	"fmt"
	"image/color"

	"github.com/kataras/iris/v12"

	"dzgCap/src/gameAreaModel"
	"dzgCap/src/gameCenter"
	"dzgCap/src/hServer/common"
	"dzgCap/src/imageTool"
)

func init() {
	RegisterHSv("saveArea", saveArea)
}

func saveArea(ctx iris.Context) {
	name := "saveArea"

	sKey := ctx.URLParam(common.Con_Params_AreaKey)
	if len(sKey) == 0 {
		err := fmt.Errorf("need areaKey")
		common.CommonHandler(name, ctx, err)
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
		err := fmt.Errorf("not find game in screen")
		common.CommonHandler(name, ctx, err)
		return
	}

	gameCenter.ScanArea()

	common.CommonHandler(name, ctx, gameAreaModel.SaveAreaModel(sKey, r))

	return
}
