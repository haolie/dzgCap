package PageViewCenter

import (
	"fmt"

	"dzgCap/model"
)

var pvMap map[model.PViewEnum]model.IPageView

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
//		@mainView:
func GoToMainView() (mainView model.IPageView) {
	for !pvMap[model.PViewEnum_Main].IsInView() {
		GoBack()
	}

	return
}

// GoBack
// @description: 返回操作
// parameter:
// return:
//		@bool: 是否成功
func GoBack() bool {
	return true
}
