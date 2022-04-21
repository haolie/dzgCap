package PageViewCenter

import (
	"fmt"
	"time"

	"dzgCap/src/ScreenModel"
	"dzgCap/src/model"
)

var pvMap = make(map[model.PViewEnum]model.IPageView, 4)

// RegisterPView
// @description: 注册PV
// parameter:
//		@enum:
//		@pv:
// return:
//		@error:
func RegisterPView(enum model.PViewEnum, pv model.IPageView) error {
	if _, exists := pvMap[enum]; exists {
		return fmt.Errorf("pageView has register")
	}

	pvMap[enum] = pv

	return nil
}

// GetPageView
// @description: 获取PV
// parameter:
//		@enum:
// return:
//		@pv:
//		@exists:
func GetPageView(enum model.PViewEnum) (pv model.IPageView, exists bool) {
	pv, exists = pvMap[enum]
	return
}

// TryDefineView
// @description: 尝试确认界面
// parameter:
// return:
//		@pv:
//		@success:
func TryDefineView() (pv model.IPageView, success bool) {
	return nil, false
}

// GoToMainView
// @description: 返回主界面
// parameter:
// return:
//		@success:
func GoToMainView() (success bool) {
	mp, exists := pvMap[model.PViewEnum_Main]
	if !exists {
		return
	}

	fn := func(n int) bool {
		for i := 0; i < n; i++ {
			if mp.IsInView() {
				return true
			}

			GoBack()
			time.Sleep(model.Sys_Con_jump_Waite)
		}

		return false
	}

	if fn(5) {
		return true
	}

	err := ScreenModel.GetCurrentScreenArea().FreshArea()
	if err != nil {
		panic(err)
	}

	return fn(5)
}

func IsMainView() bool {
	mp, exists := pvMap[model.PViewEnum_Main]
	if !exists {
		return false
	}

	return mp.IsInView()
}

// GoBack
// @description: 返回操作
// parameter:
// return:
//		@bool: 是否成功
func GoBack() bool {
	ScreenModel.GetCurrentScreenArea().ClickPointKey(0, model.Sys_Key_Point_Back)

	return true
}