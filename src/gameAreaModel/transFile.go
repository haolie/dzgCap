package gameAreaModel

import (
	"dzgCap/src/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	oldPath = "E:\\haolie\\dzgCap\\config_1.2"
	newPath = "E:/haolie/dzgCap/config_2.0"
)

type OldPointModel struct {
	//{"PointList":[{"Key":"Point_Back","X":912,"Y":140}],"RectList":[{"Key":"mainCheckRect","X":862,"Y":150,"W":19,"H":19}]}
	Key string
	X   int
	Y   int
}

type OldRectModel struct {
	Key string
	X   int
	Y   int
	W   int
	H   int
}

type OldSaveModel struct {
	PointList []*OldPointModel
	RectList  []*OldRectModel
}

func TransFile(key string) {
	oldModleMap, exists := getScreenAreaFromLocal(key)
	if !exists {
		return
	}

	modelMap := make(map[int32]*areaModel, len(oldModleMap))

	for taskId, oldItem := range oldModleMap {
		modelMap[taskId] = &areaModel{
			RectMap:  make(map[string]model.Rect, len(oldItem.RectList)),
			PointMap: make(map[string]model.Point, len(oldItem.PointList)),
		}

		for _, rectItem := range oldItem.RectList {
			modelMap[taskId].RectMap[rectItem.Key] = model.Rect{
				X: rectItem.X,
				Y: rectItem.Y,
				W: rectItem.W,
				H: rectItem.H,
			}
		}

		for _, pointItem := range oldItem.PointList {
			modelMap[taskId].PointMap[pointItem.Key] = model.Point{
				X: pointItem.X,
				Y: pointItem.Y,
			}
		}
	}

	type saveModel struct {
		TaskMap map[int32]*areaModel
	}

	saveObj := saveModel{
		TaskMap: modelMap,
	}

	data, err := json.Marshal(saveObj)
	if err != nil {
		return
	}

	savePath := path.Join(newPath, key)
	_, err = os.Stat(savePath)
	if err != nil {
		err = os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	err = ioutil.WriteFile(path.Join(savePath, key+".json"), data, 0666)
	if err != nil {
		fmt.Println(err)
	}

}

func loadOldModel(modelKey string, task int32, dirPath string) (modelObj *OldSaveModel, exists bool) {

	data, err := ioutil.ReadFile(dirPath)
	if err != nil {
		return nil, false
	}

	var temp OldSaveModel

	err = json.Unmarshal(data, &temp)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	return &temp, true
}

// GetScreenAreaFromLocal
// @description: 从本地读取游戏区域对象
// parameter:
//		@key: 区域键
// return:
//		@srObj:
//		@exists:
func getScreenAreaFromLocal(key string) (taskMap map[int32]*OldSaveModel, exists bool) {

	localPath := path.Join(oldPath, key)

	infoList, err := ioutil.ReadDir(localPath)
	if err != nil {
		return nil, false
	}

	taskMap = make(map[int32]*OldSaveModel, len(infoList))
	for _, info := range infoList {
		strParts := strings.Split(info.Name(), ".")
		if len(strParts) == 0 {
			continue
		}

		tempId, err := strconv.Atoi(strParts[0])
		if err != nil {
			continue
		}

		taskId := int32(tempId)

		fmt.Println(info.Name())
		modelObj, exists := loadOldModel(key, taskId, path.Join(localPath, info.Name()))
		if exists {
			taskMap[taskId] = modelObj
		}
	}

	return taskMap, true
}
