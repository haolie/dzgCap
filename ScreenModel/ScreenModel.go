package ScreenModel

import (
	"fmt"
	"image"

	"dzgCap/model"
)

var (
	// map[taskType]map[modelType]map[key]struct{}
	registerMap = make(map[int32]map[int32]map[string]struct{}, 4)
)

func init() {
	// 注册主界面验证区域
	RegisterModelKey(0, int32(ModelTypeEnum_Rect), model.Sys_Key_Rect_Main_Check)
	// 注册主界面验证图片
	RegisterModelKey(0, int32(ModelTypeEnum_Image), model.Sys_Key_Rect_Main_Check)
	// 注册返回点击位置
	RegisterModelKey(0, int32(ModelTypeEnum_Point), model.Sys_Key_Point_Back)
}

func RegisterModelKey(taskType, modelType int32, key string) {
	if _, exists := registerMap[taskType]; !exists {
		registerMap[taskType] = make(map[int32]map[string]struct{}, 4)
	}

	if _, exists := registerMap[taskType][modelType]; !exists {
		registerMap[taskType][modelType] = make(map[string]struct{}, 4)
	}

	registerMap[taskType][modelType][key] = struct{}{}
}

func BaseVerify(modelKey string) (success bool, errList []error) {
	return VerifyTask(modelKey, 0)
}

func VerifyTask(modeKey string, taskType int32) (success bool, errList []error) {

	modelObj, exists := GetTaskModel(modeKey, taskType)
	if !exists {
		errList = append(errList, fmt.Errorf("load file err"))
		return
	}

	errList = append(errList, verifyPointAndRect(modelObj, taskType)...)
	errList = append(errList, verifyImg(modelObj)...)

	return len(errList) == 0, errList
}

func verifyPointAndRect(modelObj *TaskSaveModel, taskType int32) (errList []error) {

	_, exists := registerMap[taskType]
	if !exists {
		return
	}

	_, exists = registerMap[taskType][int32(ModelTypeEnum_Point)]
	if !exists {
		return
	}

	for key := range registerMap[taskType][int32(ModelTypeEnum_Point)] {
		if modelObj.IsExistsPoint(key) {
			continue
		}

		errList = append(errList, fmt.Errorf("not find point %s", key))
	}

	for key := range registerMap[taskType][int32(ModelTypeEnum_Rect)] {
		if modelObj.IsExistsRect(key) {
			continue
		}

		errList = append(errList, fmt.Errorf("not find point %s", key))
	}

	return
}

func verifyImg(modelObj *TaskSaveModel) (errList []error) {
	return
}

func GetCurrentModelKey() string {
	return model.Sys_Con_Model_Base
}

func GetPointModel(taskType int32, key string) (p image.Point, exists bool) {
	modelObj, exists := GetTaskModel(GetCurrentModelKey(), taskType)
	if !exists {
		return
	}

	pm, exists := modelObj.pointMap[key]
	if !exists {
		return
	}

	p = image.Point{
		X: pm.X,
		Y: pm.Y,
	}

	return
}

func GetRectModel(taskType int32, key string) (r model.Rect, exists bool) {
	modelObj, exists := GetTaskModel(GetCurrentModelKey(), taskType)
	if !exists {
		return
	}

	rm, exists := modelObj.rectMap[key]
	if !exists {
		return
	}

	r = model.Rect{
		X: rm.X,
		Y: rm.Y,
		W: rm.W,
		H: rm.H,
	}

	return
}
