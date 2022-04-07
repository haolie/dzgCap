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
