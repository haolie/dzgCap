package gameAreaModel

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"dzgCap/ConfigManger"
	"dzgCap/Loger"
	"dzgCap/dzgCap"
)

var (
	modelMap   map[string]areaModel
	moduleName = "gameAreaModel"
)

func init() {
	dzgCap.RegisterLoad(moduleName, loadHandler)
}

func loadHandler(ctx context.Context) (errList []error) {
	var err error
	modelMap, err = loadAreaModel(ConfigManger.GetConfigCopy().ConfigPath)
	if err != nil {
		errList = append(errList, err)
	}

	return
}

func loadAreaModel(dirPath string) (modelMap map[string]areaModel, err error) {
	dirList, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}

	modelMap = make(map[string]areaModel, len(dirList))
	for _, d := range dirList {
		modelKey := d.Name()

		modelObj, exists := loadModel(path.Join(dirPath, modelKey, modelKey+".json"))
		if !exists {
			continue
		}

		modelMap[modelKey] = modelObj
	}

	return
}

func loadModel(filePath string) (modelObj areaModel, exists bool) {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		Loger.LogErr(fmt.Sprintf("gameAreaModel.loadModel err:%v", err))
		return
	}

	err = json.Unmarshal(data, &modelObj)
	if err != nil {
		panic(err)
	}

	exists = true

	return
}

func saveAreaModel(modelObj areaModel, key, savePath string) {
	data, err := json.Marshal(&modelObj)
	if err != nil {
		panic(err)
	}

	savePath = path.Join(savePath, key)
	_, err = os.Stat(savePath)
	if err != nil {
		err = os.MkdirAll(savePath, os.ModePerm)
		panic(err)
	}

	err = ioutil.WriteFile(path.Join(savePath, key+".json"), data, 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func getAreaModel(key string) (am areaModel, exists bool) {
	am, exists = modelMap[key]
	return
}

func getAreaTaskModel(key string, taskType int32) (modelObj areaTaskModel, exists bool) {
	rm, exists := getAreaModel(key)
	if !exists {
		return
	}

	modelObj, exists = rm.TaskMap[taskType]
	return
}
