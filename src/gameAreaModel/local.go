package gameAreaModel

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"

	"dzgCap/ConfigManger"
	"dzgCap/Loger"
	"dzgCap/dzgCap"
	"dzgCap/src/imageTool"
	"dzgCap/src/model"
)

var (
	modelMap   map[string]*areaModel
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

func loadAreaModel(dirPath string) (modelMap map[string]*areaModel, err error) {
	dirList, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}

	modelMap = make(map[string]*areaModel, len(dirList))
	for _, d := range dirList {
		modelKey := d.Name()

		modelObj, exists := loadModel(path.Join(dirPath, modelKey, modelKey+".json"))
		if !exists {
			continue
		}

		modelMap[modelKey] = &modelObj
	}

	return
}

// 从本地文件中加载游戏模型
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
	if modelObj.TaskMap == nil {
		modelObj.TaskMap = make(map[int32]*areaTaskModel, 4)
	}

	//
	for taskId := range modelObj.TaskMap {
		if modelObj.TaskMap[taskId].PointMap == nil {
			modelObj.TaskMap[taskId].PointMap = make(map[string]model.Point, 4)
		}

		if modelObj.TaskMap[taskId].RectMap == nil {
			modelObj.TaskMap[taskId].RectMap = make(map[string]model.Rect, 4)
		}

		if modelObj.TaskMap[taskId].imgMap == nil {
			modelObj.TaskMap[taskId].imgMap = make(map[string]image.Image, 4)
		}
	}

	return
}

// saveAreaModel
// @description: 保存游戏区域模型
// parameter:
//		@modelObj:
//		@savePath:
// return:
//		@error:
func saveAreaModel(modelObj *areaModel, savePath string) error {
	data, err := json.Marshal(modelObj)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(savePath)
	if err != nil {
		err = os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(path.Join(savePath, modelObj.Key+".json"), data, 0666)
}

// 返回游戏区域模型
func getAreaModel(key string) (am *areaModel, exists bool) {
	am, exists = modelMap[key]
	return
}

// 返回任务模型
func getAreaTaskModel(key string, taskType int32) (modelObj *areaTaskModel, exists bool) {
	rm, exists := getAreaModel(key)
	if !exists {
		return
	}

	modelObj, exists = rm.TaskMap[taskType]
	return
}

// 返回游戏区域保存路径
func getAreaModelPath(key string) string {
	return path.Join(ConfigManger.GetConfigCopy().ConfigPath, key)
}

func loadImg(areaKey string, taskId int32, key string) (img image.Image, exists bool) {
	fileName := path.Join(getAreaModelPath(areaKey), fmt.Sprintf("%d", taskId), "pic", key+".png")
	img, err := imageTool.LoadImage(fileName)
	if err != nil {
		return
	}

	exists = true
	return
}

func saveImg(areaKey string, taskId int32, key string, img image.Image) error {
	savePath := path.Join(getAreaModelPath(areaKey), fmt.Sprintf("%d", taskId), "pic")
	_, err := os.Stat(savePath)
	if err != nil {
		err = os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return imageTool.SaveImage(img, path.Join(savePath, key+".png"))
}
