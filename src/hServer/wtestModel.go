// ************************************
// @package: hServer
// @description: 测试接口
// @author:
// @revision history:
// @create date: 2022-04-22 11:34:53
// ************************************

package hServer

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"

	"dzgCap/src/PageView/PageViewCenter"
	"dzgCap/src/ScreenModel"
	"dzgCap/src/imageTool"
	"dzgCap/src/model"
)

func init() {
	RegisterHSv("wtest", wtest)
}

func wtest(ctx iris.Context) {

	gameRect, exists := ScreenModel.GetCurrentScreenArea().GetRect(0, model.Sys_Key_Rect_Game)
	if !exists {
		ctx.WriteString("not find Sys_Key_Rect_Game")
		return
	}

	//p, exists := ScreenModel.GetCurrentScreenArea().GetPoint(0, model.Syc_Key_Point_Icon_Line)
	//if !exists {
	//	ctx.WriteString("not find Syc_Key_Point_Icon_Line")
	//	return
	//}

	padding := 2
	iconW := (gameRect.W - padding*2) / 6
	for i := 0; i < 6; i++ {
		iconR := model.Rect{
			X: gameRect.X + iconW*i + padding,
			Y: gameRect.Y,
			W: iconW,
			H: gameRect.H,
		}

		img := imageTool.CapScreen(iconR)
		imageTool.SaveImage(img, fmt.Sprintf("./config/test/%d.png", i))

		clickIcon(iconR.X+iconR.W/2, 1156)
		time.Sleep(time.Second)
	}

	ctx.WriteString("success")

}

func clickIcon(x, y int) {

	for i := 0; i < 5; i++ {
		if PageViewCenter.IsMainView() {
			break
		}

		PageViewCenter.GoBack()
	}
	//
	//if !PageViewCenter.IsMainView() {
	//	return
	//}

	ScreenModel.GetCurrentScreenArea().ClickPoint(x, y)
	time.Sleep(time.Millisecond * 1500)
	if isMeeting() {
		ScreenModel.GetCurrentScreenArea().ClickPointKey(1, model.Sys_Key_Point_Meeting_GuestNumReward)
		time.Sleep(time.Second * 2)
		ScreenModel.GetCurrentScreenArea().ClickPointKey(1, model.Sys_Key_Point_Meeting_DrawNumReward)
		time.Sleep(time.Second * 1)
	}
	PageViewCenter.GoBack()
}

func isMeeting() bool {
	sm, err := ScreenModel.GetCurrentScreenArea().CompareRectToCash(1, model.Sys_key_rect_Meeting_List)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if sm {
		ScreenModel.GetCurrentScreenArea().ClickPoint(2479, 1115)
		time.Sleep(time.Second)
	}

	sm, err = ScreenModel.GetCurrentScreenArea().CompareRectToCash(1, model.Sys_key_rect_Meeting_SiWangYan)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if sm {
		return true
	}

	sm, err = ScreenModel.GetCurrentScreenArea().CompareRectToCash(1, model.Sys_key_rect_Meeting_QinWangYan)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if sm {
		return true
	}

	sm, err = ScreenModel.GetCurrentScreenArea().CompareRectToCash(1, model.Sys_key_rect_Meeting_JunWangYan)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(sm)

	if sm {
		return true
	}

	return false

}
