package ScreenModel

import (
	"fmt"
)

var (
	// map[taskType]map[modelType]map[key]struct{}
	registerMap = make(map[int32]map[int32]map[string]struct{}, 4)
)

func RegisterModelKey(taskType, modelType int32, key string) {
	if _, exists := registerMap[taskType]; !exists {
		registerMap[taskType] = make(map[int32]map[string]struct{}, 4)
	}

	if _, exists := registerMap[taskType][modelType]; !exists {
		registerMap[taskType][modelType] = make(map[string]struct{}, 4)
	}

	registerMap[taskType][modelType][key] = struct{}{}
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
