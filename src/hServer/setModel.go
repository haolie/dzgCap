package hServer

import (
	"image/color"

	"github.com/kataras/iris"

	"dzgCap/src/ScreenModel"
	"dzgCap/src/capMager"
	"dzgCap/src/imageTool"
	"dzgCap/src/model"
)

func init() {
	RegisterHSv("savePoint", savePoint)
	RegisterHSv("saveRect", saveRect)
	RegisterHSv("saveImage", saveImage)
	RegisterHSv("saveScreenBase", saveScreenBase)
}

func savePoint(ctx iris.Context) {

}

func saveRect(ctx iris.Context) {

}

func saveImage(ctx iris.Context) {

}

func saveScreenBase(ctx iris.Context) {
	img := imageTool.CapFullScreen()
	r, exists := capMager.FindMinRect(img, color.RGBA{
		R: 20,
		G: 24,
		B: 31,
		A: 255,
	})

	if !exists {
		return
	}

	sm, exists := ScreenModel.GetTaskModel(ScreenModel.GetCurrentModelKey(), 0)
	if !exists {
		sm.AddRect(ScreenModel.RectModel{model.Sys_Key_Rect_Game, r.X, r.Y, r.W, r.H})
		ScreenModel.SaveTaskModel(ScreenModel.GetCurrentModelKey(), 0, sm)
	}

}
