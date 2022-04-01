package ScreenModel

import (
	"fmt"
	"testing"
)

func TestSaveMode(t *testing.T) {
	modeObj := NewTaskSaveModel()
	modeObj.AddPoint(PointModel{"1", 3, 2})
	modeObj.AddPoint(PointModel{"2", 33, 22})

	modeObj.AddRect(RectModel{"2", 33, 22, 33, 22})
	modeObj.AddRect(RectModel{"44", 44, 44, 33, 22})

	err := SaveTaskModel("Base", 1, modeObj)
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
