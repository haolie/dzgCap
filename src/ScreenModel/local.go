package ScreenModel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"dzgCap/src/model"
)

var (
	basePath = "./"
)

// getTaskModel
// @description: 根据模式键和任务从磁盘获取任务模型
// parameter:
//		@modelKey: 屏幕模式键
//		@taskId: 任务Id
// return:
//		@model: 任务保存模型
//		@exists: 是否存在
func getTaskModel(modelKey string, taskId int32) (model *TaskSaveModel, exists bool) {
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

	return &temp, true
}

// saveTaskModel
// @description: 保存任务模型
// parameter:
//		@modelKey: 屏幕模式Id
//		@taskId: 任务Id
//		@model: 任务模型
// return:
//		@error:
func saveTaskModel(modelKey string, taskId int32, model *TaskSaveModel) error {

	data, err := json.Marshal(*model)
	if err != nil {
		return err
	}

	fileName := getSavePath(modelKey, taskId)

	err = ioutil.WriteFile(fileName, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

// getSavePath
// @description: 组装磁盘地址
// parameter:
//		@modelKey:
//		@taskId:
// return:
//		@string:
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
