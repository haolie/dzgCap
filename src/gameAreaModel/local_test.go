package gameAreaModel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

const con_screen_key = "centerScreen"

//#region test

type S2 struct {
	I int
	S string
}

type S1 struct {
	M1 map[int]S2
	M2 map[string]S2
}

//#endregion

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

	obj := &S1{
		M1: make(map[int]S2, 2),
		M2: make(map[string]S2, 2),
	}

	obj.M1[1] = S2{
		I: 1,
		S: "1",
	}

	obj.M1[2] = S2{
		I: 2,
		S: "2",
	}

	obj.M2["1"] = S2{
		I: 1,
		S: "1",
	}

	obj.M2["2"] = S2{
		I: 2,
		S: "2",
	}

	data, err := json.Marshal(*obj)
	if err != nil {
		return
	}

	err = ioutil.WriteFile("test.json", data, 0666)
	if err != nil {
		return
	}

}

func TestGetModel(t *testing.T) {
	modelMap, _ := loadAreaModel("../../config_b_2.0/")
	fmt.Println(modelMap)
}

func TestVerifyModel(t *testing.T) {

}
