package ScreenModel

import (
	"fmt"
	"testing"
)

const con_screen_key = "centerScreen"

func TestSaveMode(t *testing.T) {
	//basePath = "../../"
	//
	//
	//modeObj := NewTaskSaveModel()
	//
	//err := SaveRectImg(model.Rect{1204, 1125, 80, 20}, model.Sys_Key_Rect_Meeting_Join_Btn)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//// 2497,1356,  35 7
	//modeObj.AddRect(RectModel{model.Sys_Key_Rect_Meeting_Join_Btn, 1204, 1125, 80, 20})
	//
	//err = SaveTaskModel("centerScreen", 1, modeObj)
	//fmt.Println(err)
}

func TestSaveMain(t *testing.T) {
	//basePath = "../../"
	//
	//bm, exists := GetTaskModel("centerScreen", 0)
	//if !exists {
	//	bm = NewTaskSaveModel()
	//}
	//
	//err := SaveRectImg(model.Rect{919, 467, 25, 25}, model.Sys_Key_Rect_Main_Check)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//bm.AddRect(RectModel{model.Sys_Key_Rect_Main_Check, 919, 467, 25, 25})
	//bm.AddPoint(PointModel{model.Sys_Key_Point_Back, 986, 473})
	//
	//err = SaveTaskModel("centerScreen", 0, bm)
	//fmt.Println(err)
}

func TestGetModel(t *testing.T) {
	//d, exists := GetTaskModel("Base", 1)
	//if exists {
	//	fmt.Println(*d)
	//} else {
	//	fmt.Println("load flaid")
	//}

}

func TestVerifyModel(t *testing.T) {
	RegisterModelKey(1, int32(ModelTypeEnum_Point), "1")
	RegisterModelKey(1, int32(ModelTypeEnum_Point), "2")

	RegisterModelKey(1, int32(ModelTypeEnum_Rect), "2")
	RegisterModelKey(1, int32(ModelTypeEnum_Rect), "44")

	su, list := VerifyTask("Base", 1)
	if !su {
		fmt.Println(list)
	}
}
