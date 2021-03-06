package ScreenModel

import (
	"fmt"

	"dzgCap/src/model"
)

var (
	// map[taskType]map[modelType]map[key]struct{}
	registerMap     = make(map[int32]map[int32]map[string]struct{}, 4)
	currentModelKey = model.Sys_Con_Model_Base

	taskStatusFun      func() model.TaskStatusEnum
	screenModelCash    = make(map[string]*ScreenArea, 4)
	currentScreenModel *ScreenArea
)

func init() {
	// 注册主界面验证区域
	RegisterModelKey(0, int32(ModelTypeEnum_Rect), model.Sys_Key_Rect_Main_Check)
	// 注册主界面验证图片
	RegisterModelKey(0, int32(ModelTypeEnum_Image), model.Sys_Key_Rect_Main_Check)
	// 注册返回点击位置
	RegisterModelKey(0, int32(ModelTypeEnum_Point), model.Sys_Key_Point_Back)
	// 注册游戏区域
	RegisterModelKey(0, int32(ModelTypeEnum_Rect), model.Sys_Key_Rect_Game)

}

// RegisterModelKey
// @description: 注册模型（点、区域、图片）
// parameter:
//		@taskType: 任务类型（任务Id）
//		@modelType: 模型类型（点、区域、图片）
//		@key: 键值
// return:
func RegisterModelKey(taskType, modelType int32, key string) {
	if _, exists := registerMap[taskType]; !exists {
		registerMap[taskType] = make(map[int32]map[string]struct{}, 4)
	}

	if _, exists := registerMap[taskType][modelType]; !exists {
		registerMap[taskType][modelType] = make(map[string]struct{}, 4)
	}

	registerMap[taskType][modelType][key] = struct{}{}
}

// 返回当前游戏区域模型
func GetCurrentScreenArea() *ScreenArea {
	return currentScreenModel
}

// 加载当前游戏区域模型
func LoadScreenArea(key string) {
	var exists bool
	currentScreenModel, exists = GetScreenArea(key)
	if !exists {
		panic("not find ScreenArea:" + key)
	}

	err := currentScreenModel.FreshArea()
	if err != nil {
		panic(err)
	}
}

// 返回指定游戏区域模型
func GetScreenArea(key string) (sr *ScreenArea, exists bool) {
	sr, exists = screenModelCash[key]
	if exists {
		return
	}

	sr, exists = GetScreenAreaFromLocal(key)
	if exists {
		screenModelCash[key] = sr
	}

	return
}

// 验证基本项
func BaseVerify(modelKey string) (success bool, errList []error) {
	return VerifyTask(modelKey, 0)
}

// VerifyTask
// @description: 验证任务模型
// parameter:
//		@modeKey: 游戏区域键值
//		@taskType:  任务类型
// return:
//		@success:
//		@errList:
func VerifyTask(modeKey string, taskType int32) (success bool, errList []error) {

	modelObj, exists := GetScreenArea(modeKey)
	if !exists {
		errList = append(errList, fmt.Errorf("load file err"))
		return
	}

	// 验证点和区域
	errList = append(errList, verifyPointAndRect(modelObj, taskType)...)
	// 验证图片
	errList = append(errList, verifyImg(modelObj)...)

	return len(errList) == 0, errList
}

// 验证点和区域
func verifyPointAndRect(modelObj *ScreenArea, taskType int32) (errList []error) {

	_, exists := registerMap[taskType]
	if !exists {
		return
	}

	_, exists = registerMap[taskType][int32(ModelTypeEnum_Point)]
	if !exists {
		return
	}

	// 验证点
	for key := range registerMap[taskType][int32(ModelTypeEnum_Point)] {
		if modelObj.IsExistsPoint(taskType, key) {
			continue
		}

		errList = append(errList, fmt.Errorf("not find point %s", key))
	}

	// 验证区域
	for key := range registerMap[taskType][int32(ModelTypeEnum_Rect)] {
		if modelObj.IsExistsRect(taskType, key) {
			continue
		}

		errList = append(errList, fmt.Errorf("not find point %s", key))
	}

	return
}

func verifyImg(modelObj *ScreenArea) (errList []error) {
	return
}

func GetCurrentModelKey() string {
	return GetCurrentScreenArea().Key
}
