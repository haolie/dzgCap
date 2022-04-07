package hServer

import (
	"fmt"
	"image"
	"image/color"

	"github.com/kataras/iris/v12"

	"dzgCap/src/ScreenModel"
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
	sKey := ctx.URLParam("screen")
	if len(sKey) == 0 {
		ctx.WriteString("need key")
		return
	}

	taskId32, err := ctx.URLParamInt("taskId")
	if err != nil {
		ctx.WriteString("need taskId")
		return
	}

	pkey := ctx.URLParam("point")
	if len(pkey) == 0 {
		ctx.WriteString("need pointKey")
		return
	}

	taskId := int32(taskId32)

	x, err := ctx.URLParamInt("x")
	if err != nil {
		ctx.WriteString("need x")
		return
	}

	y, err := ctx.URLParamInt("y")
	if err != nil {
		ctx.WriteString("need y")
		return
	}

	sr, exists := ScreenModel.GetScreenArea(sKey)
	if !exists {
		ctx.WriteString("not find screen:" + sKey)
		return
	}

	sr.AddPoint(taskId, pkey, image.Point{X: x, Y: y})
	err = ScreenModel.SaveScreenModel(sr)
	if err != nil {
		ctx.WriteString(fmt.Sprintf("%s", err))
		return
	}

	ctx.WriteString("success")
}

func getRectParam(ctx iris.Context) (key string, taskId int32, r model.Rect, err error) {
	key = ctx.URLParam("rect")
	if len(key) == 0 {
		ctx.WriteString("need key")
		return
	}

	taskId32, err := ctx.URLParamInt("taskId")
	if err != nil {
		ctx.WriteString("need taskId")
		return
	}

	taskId = int32(taskId32)

	r.X, err = ctx.URLParamInt("x")
	if err != nil {
		ctx.WriteString("need x")
		return
	}

	r.Y, err = ctx.URLParamInt("y")
	if err != nil {
		ctx.WriteString("need y")
		return
	}

	r.W, err = ctx.URLParamInt("w")
	if err != nil {
		ctx.WriteString("need w")
		return
	}

	r.H, err = ctx.URLParamInt("h")
	if err != nil {
		ctx.WriteString("need h")
		return
	}

	return
}

func saveRect(ctx iris.Context) {
	sKey := ctx.URLParam("screen")
	if len(sKey) == 0 {
		ctx.WriteString("need key")
		return
	}

	rkey, taskId, r, err := getRectParam(ctx)
	if err != nil {
		return
	}

	sr, exists := ScreenModel.GetScreenArea(sKey)
	if !exists {
		ctx.WriteString("not find screen:" + sKey)
		return
	}

	sr.AddRect(taskId, rkey, r)
	err = ScreenModel.SaveScreenModel(sr)
	if err != nil {
		ctx.WriteString(fmt.Sprintf("%s", err))
		return
	}

	ctx.WriteString("success")
}

func saveImage(ctx iris.Context) {
	sKey := ctx.URLParam("screen")
	if len(sKey) == 0 {
		ctx.WriteString("need key")
		return
	}

	rkey := ctx.URLParam("rect")
	if len(rkey) == 0 {
		ctx.WriteString("need rect")
		return
	}

	taskId32, err := ctx.URLParamInt("taskId")
	if err != nil {
		ctx.WriteString("need taskId")
		return
	}

	taskId := int32(taskId32)

	sr, exists := ScreenModel.GetScreenArea(sKey)
	if !exists {
		ctx.WriteString("not find screen:" + sKey)
		return
	}

	err = sr.SaveRectImg(taskId, rkey)
	if err != nil {
		ctx.WriteString(fmt.Sprintf("%s", err))
		return
	}

	ctx.WriteString("success")
}

func saveScreenBase(ctx iris.Context) {
	sKey := ctx.URLParam("screen")
	if len(sKey) == 0 {
		ctx.WriteString("need key")
		return
	}

	sr, exists := ScreenModel.GetScreenArea(sKey)
	if !exists {
		sr = ScreenModel.NewScreenArea(sKey)
	}

	img := imageTool.CapFullScreen()
	r, exists := imageTool.FindMinRect(img, color.RGBA{
		R: 20,
		G: 24,
		B: 31,
		A: 255,
	})

	if !exists {
		return
	}

	sr.AddRect(0, model.Sys_Key_Rect_Game, r)
	err := ScreenModel.SaveScreenModel(sr)
	if err != nil {
		ctx.WriteString(fmt.Sprintf("%s", err))
		return
	}

	ctx.WriteString("success")
}
