package ScreenModel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"dzgCap/src/model"
)

var (
	modelCashMap = make(map[string]map[int32]*TaskSaveModel, 4)
	basePath     = "./"
)

func GetTaskModel(modelKey string, taskId int32) (model *TaskSaveModel, exists bool) {
	if _, exists := modelCashMap[modelKey]; exists {
		if model, exists = modelCashMap[modelKey][taskId]; exists {
			return model, exists
		}
	}

	fileName := getSavePath(modelKey, taskId)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, false
	}

	var temp TaskSaveModel

	err = json.Unmarshal(data, &temp)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	temp.pointMap = make(map[string]PointModel, len(temp.PointList))
	temp.rectMap = make(map[string]RectModel, len(temp.RectList))

	for _, item := range temp.PointList {
		temp.pointMap[item.Key] = item
	}

	for _, item := range temp.RectList {
		temp.rectMap[item.Key] = item
	}

	saveCash(modelKey, taskId, &temp)

	return &temp, true
}

func saveCash(modelKey string, taskId int32, model *TaskSaveModel) {

	if _, exists := modelCashMap[modelKey]; !exists {
		modelCashMap[modelKey] = make(map[int32]*TaskSaveModel)
	}

	modelCashMap[modelKey][taskId] = model
}

func SaveTaskModel(modelKey string, taskId int32, model *TaskSaveModel) error {
	model.PointList = make([]PointModel, 0, len(model.pointMap))
	for _, item := range model.pointMap {
		model.PointList = append(model.PointList, item)
	}

	model.RectList = make([]RectModel, 0, len(model.rectMap))
	for _, item := range model.rectMap {
		model.RectList = append(model.RectList, item)
	}

	data, err := json.Marshal(*model)
	if err != nil {
		return err
	}

	fileName := getSavePath(modelKey, taskId)

	err = ioutil.WriteFile(fileName, data, 0666)
	if err != nil {
		return err
	}

	saveCash(modelKey, taskId, model)

	return nil
}

func getSavePath(modelKey string, taskId int32) string {

	dirPath := fmt.Sprintf("%s%s/%s", basePath, model.Sys_Con_Path_Config, modelKey)

	_, err := os.Stat(dirPath)
	if err != nil {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	return fmt.Sprintf("%s/%d.ini", dirPath, taskId)
}
