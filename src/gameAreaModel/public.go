package gameAreaModel

import (
	"fmt"
	"image"

	"dzgCap/Loger"
	"dzgCap/src/imageTool"
	. "dzgCap/src/model"
)

func GetModelKeyWithRect(r Rect) (modelKey string, exists bool) {
	for k, mod := range modelMap {
		if _, exists := mod.TaskMap[0]; !exists {
			continue
		}

		if mod.TaskMap[0].RectMap[Sys_Key_Rect_Game] == r {
			return k, true
		}
	}

	return
}

func GetRect(areaModel string, taskType int32, key string) (r Rect, exists bool) {
	areaTaskModelObj, exists := getAreaTaskModel(areaModel, taskType)
	if !exists {
		return
	}

	r, exists = areaTaskModelObj.RectMap[key]

	return
}

func GetPoint(areaModel string, taskType int32, key string) (p Point, exists bool) {
	areaTaskModelObj, exists := getAreaTaskModel(areaModel, taskType)
	if !exists {
		return
	}

	p, exists = areaTaskModelObj.PointMap[key]

	return
}

func GetImage(areaModel string, taskType int32, key string) (img image.Image, exists bool) {
	areaTaskModelObj, exists := getAreaTaskModel(areaModel, taskType)
	if !exists {
		return
	}

	if areaTaskModelObj.imgMap == nil {
		return
	}

	img, exists = areaTaskModelObj.imgMap[key]
	if exists {
		return
	}

	img, exists = loadImg(areaModel, taskType, key)
	if exists {
		areaTaskModelObj.imgMap[key] = img
	}

	return
}

func SaveAreaModel(areaModelKey string, r Rect) error {
	_, exists := modelMap[areaModelKey]
	if !exists {
		modelMap[areaModelKey] = createNewArea(areaModelKey)
	}

	if _, exists := modelMap[areaModelKey].TaskMap[0]; !exists {
		modelMap[areaModelKey].TaskMap[0] = createAreaTaskModel()
	}

	modelMap[areaModelKey].TaskMap[0].RectMap[Sys_Key_Rect_Game] = r

	saveDir := getAreaModelPath(areaModelKey)
	return saveAreaModel(modelMap[areaModelKey], saveDir)
}

func SaveRect(areaKey string, taskType int32, key string, r Rect) error {
	_, exists := modelMap[areaKey]
	if !exists {
		return fmt.Errorf("not find gameArea")
	}

	_, exists = modelMap[areaKey].TaskMap[taskType]
	if !exists {
		modelMap[areaKey].TaskMap[taskType] = createAreaTaskModel()
	}

	modelMap[areaKey].TaskMap[taskType].RectMap[key] = r

	saveDir := getAreaModelPath(areaKey)
	return saveAreaModel(modelMap[areaKey], saveDir)
}

func SavePoint(areaKey string, taskType int32, key string, p Point) error {
	_, exists := modelMap[areaKey]
	if !exists {
		return fmt.Errorf("not find gameArea")
	}

	_, exists = modelMap[areaKey].TaskMap[taskType]
	if !exists {
		modelMap[areaKey].TaskMap[taskType] = createAreaTaskModel()
	}

	modelMap[areaKey].TaskMap[taskType].PointMap[key] = p

	saveDir := getAreaModelPath(areaKey)
	return saveAreaModel(modelMap[areaKey], saveDir)
}

func SaveImage(areaKey string, taskType int32, key string) error {
	_, exists := modelMap[areaKey]
	if !exists {
		return fmt.Errorf("not find gameArea")
	}

	_, exists = modelMap[areaKey].TaskMap[taskType]
	if !exists {
		modelMap[areaKey].TaskMap[taskType] = createAreaTaskModel()
	}

	img, err := CapRectImg(areaKey, taskType, key, 0, 0)
	if err != nil {
		return err
	}

	err = saveImg(areaKey, taskType, key, img)
	if err != nil {
		return err
	}

	modelMap[areaKey].TaskMap[taskType].imgMap[key] = img

	return nil
}

func CapRectImg(areaKey string, taskType int32, rectKey string, mx, my int) (img image.Image, err error) {
	r, exists := GetRect(areaKey, taskType, rectKey)
	if !exists {
		err = fmt.Errorf("not find rect")
		return
	}

	r = r.Move(mx, my)
	img = imageTool.CapScreen(r)
	return
}

func VerifyRect(areaModel string, taskType int32, key string, mx, my int) bool {
	oldImg, exists := GetImage(areaModel, taskType, key)
	if !exists {
		return false
	}

	img, err := CapRectImg(areaModel, taskType, key, mx, my)
	if err != nil {
		Loger.LogErr(fmt.Sprintf("%s.VerifyRect 抓取图片出错 err:%s", moduleName, err))
		return false
	}

	saveImg(areaModel, taskType, "temp", img)

	return imageTool.CompareImage(img, oldImg)
}
