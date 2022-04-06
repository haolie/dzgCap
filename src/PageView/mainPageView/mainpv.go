package mainPageView

import (
	"fmt"

	"dzgCap/src/PageView/PageViewCenter"
	"dzgCap/src/ScreenModel"
	"dzgCap/src/capMager"
	. "dzgCap/src/model"
)

func init() {
	PageViewCenter.RegisterPView(PViewEnum_Main, new(mainPv))
}

type mainPv struct {
}

func (mp *mainPv) GetId() int32 {
	return int32(PViewEnum_Main)
}

func (mp *mainPv) GetName() string {
	name, _ := GetPVName(PViewEnum_Main)
	return name
}

func (mp *mainPv) GetEnterPointList() {

}

func (mp *mainPv) GoToView(key string) bool {
	return false
}

func (mp *mainPv) IsInView() bool {
	r, exists := ScreenModel.GetRectModel(0, Sys_Key_Rect_Main_Check)
	if !exists {
		return exists
	}

	isMain, err := capMager.CompareRectToCash(r, Sys_Key_Rect_Main_Check)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return isMain
}
