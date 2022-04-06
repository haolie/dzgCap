package ScreenModel

import (
	"fmt"
	"testing"

	"dzgCap/src/model"
)

func TestSaveMode(t *testing.T) {
	basePath = "../../"

	modeObj := NewTaskSaveModel()

	//err := capMager.SaveRectImg(model.Rect{2497, 1356, 35, 7}, model.Sys_Key_Rect_Meeting_Join_Btn)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// 2497,1356,  35 7
	modeObj.AddRect(RectModel{model.Sys_Key_Rect_Meeting_Join_Btn, 2497, 1356, 35, 7})

	err := SaveTaskModel("miniScreen", 1, modeObj)
	fmt.Println(err)
}

func TestSaveMain(t *testing.T) {
	basePath = "../../"

	bm, exists := GetTaskModel("miniScreen", 0)
	if !exists {
		bm = NewTaskSaveModel()
	}

	//err := capMager.SaveRectImg(model.Rect{2383, 1086, 15, 15}, model.Sys_Key_Rect_Main_Check)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	bm.AddRect(RectModel{model.Sys_Key_Rect_Main_Check, 2383, 1086, 15, 15})
	bm.AddPoint(PointModel{model.Sys_Key_Point_Back, 2400, 1082})

	err := SaveTaskModel("miniScreen", 0, bm)
	fmt.Println(err)
}

func TestGetModel(t *testing.T) {
	d, exists := GetTaskModel("Base", 1)
	if exists {
		fmt.Println(*d)
	} else {
		fmt.Println("load flaid")
	}

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
