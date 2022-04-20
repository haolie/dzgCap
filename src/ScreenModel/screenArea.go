package ScreenModel

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"

	"dzgCap/src/imageTool"
	"dzgCap/src/model"
)

type ScreenArea struct {
	Key          string
	currentRect  model.Rect
	taskPointMap map[int32]map[string]image.Point
	taskRectMap  map[int32]map[string]model.Rect
	taskImageMap map[int32]map[string]image.Image
}

func NewScreenArea(key string) *ScreenArea {
	return &ScreenArea{
		Key:          key,
		currentRect:  model.Rect{},
		taskPointMap: make(map[int32]map[string]image.Point, 8),
		taskRectMap:  make(map[int32]map[string]model.Rect, 8),
		taskImageMap: make(map[int32]map[string]image.Image, 8),
	}
}

func (sr *ScreenArea) AddPoint(taskId int32, key string, p image.Point) {
	if sr.taskPointMap == nil {
		sr.taskPointMap = make(map[int32]map[string]image.Point, 8)
	}

	if _, exists := sr.taskPointMap[taskId]; !exists {
		sr.taskPointMap[taskId] = make(map[string]image.Point, 8)
	}

	sr.taskPointMap[taskId][key] = p
}

func (sr *ScreenArea) AddRect(taskId int32, key string, r model.Rect) {
	if sr.taskRectMap == nil {
		sr.taskRectMap = make(map[int32]map[string]model.Rect, 8)
	}

	if _, exists := sr.taskRectMap[taskId]; !exists {
		sr.taskRectMap[taskId] = make(map[string]model.Rect, 8)
	}

	sr.taskRectMap[taskId][key] = r
}

func (sr *ScreenArea) IsExistsPoint(taskId int32, key string) bool {

	_, exists := sr.taskPointMap[taskId]
	if !exists {
		return exists
	}

	_, exists = sr.taskPointMap[taskId][key]

	return exists
}

func (sr *ScreenArea) IsExistsRect(taskId int32, key string) bool {
	_, exists := sr.taskRectMap[taskId]
	if !exists {
		return exists
	}

	_, exists = sr.taskRectMap[taskId][key]

	return exists
}

func (sr *ScreenArea) FreshArea() error {
	img := imageTool.CapFullScreen()
	r, exists := imageTool.FindMinRect(img, color.RGBA{
		R: 20,
		G: 24,
		B: 31,
		A: 255,
	})

	if !exists {
		return fmt.Errorf("miss game area")
	}

	empt := model.Rect{}
	if sr.currentRect != empt && (r.H != sr.currentRect.H || r.W != sr.currentRect.W) {
		return fmt.Errorf("screen area changed")
	}

	sr.currentRect = r

	return nil
}

//#region image

// 组装图片保存路径
func baseImgPath(taskId int32, screenkey string) string {
	dirPath := fmt.Sprintf("./%s/%s/%d/pic", model.Sys_Con_Path_Config, screenkey, taskId)
	_, err := os.Stat(dirPath)
	if err != nil {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	return dirPath
}

func (sr *ScreenArea) GetCashImg(taskId int32, name string) (img image.Image, exists bool, err error) {
	if _, exists := sr.taskImageMap[taskId]; !exists {
		sr.taskImageMap[taskId] = make(map[string]image.Image)
	}

	img, exists = sr.taskImageMap[taskId][name]
	if exists {
		return
	}

	fileName := path.Join(baseImgPath(taskId, sr.Key), name+".png")
	img, err = imageTool.LoadImage(fileName)
	if err != nil {
		return
	}

	sr.taskImageMap[taskId][name] = img

	return img, true, nil
}

func (sr *ScreenArea) SaveRectImg(taskId int32, rectKey string) error {
	if !sr.IsExistsRect(taskId, rectKey) {
		return fmt.Errorf("not find rect:" + rectKey)
	}

	img := imageTool.CapScreen(sr.taskRectMap[taskId][rectKey])

	fileName := path.Join(baseImgPath(taskId, sr.Key), rectKey+".png")
	err := imageTool.SaveImage(img, fileName)
	if err != nil {
		return err
	}

	if _, exists := sr.taskImageMap[taskId]; !exists {
		sr.taskImageMap[taskId] = make(map[string]image.Image)
	}

	sr.taskImageMap[taskId][rectKey] = img

	return nil
}

// 相对位置
func (sr *ScreenArea) CompareRectToCash(taskId int32, rectKey string) (isSame bool, err error) {
	if !sr.IsExistsRect(taskId, rectKey) {
		return false, nil
	}

	cashImg, exists, err := sr.GetCashImg(taskId, rectKey)
	if err != nil || !exists {
		return
	}

	mx, my := sr.getMove()
	r := sr.taskRectMap[taskId][rectKey]
	r = r.Move(mx, my)

	img := imageTool.CapScreen(r)
	isSame = imageTool.CompareImage(cashImg, img)

	return
}

//#endregion

//#region 位移

func (sr *ScreenArea) getMove() (x, y int) {
	if !sr.IsExistsRect(0, model.Sys_Key_Rect_Game) {
		return
	}

	r := sr.taskRectMap[0][model.Sys_Key_Rect_Game]

	return sr.currentRect.X - r.X, sr.currentRect.Y - r.Y
}

//#endregion

//#region 点击

// 相对位置
func (sr *ScreenArea) ClickPoint(x, y int) {
	robotgo.MoveClick(x, y)
}

// 相对位置
func (sr *ScreenArea) ClickPointKey(taskId int32, pKey string) {
	if !sr.IsExistsPoint(taskId, pKey) {
		return
	}

	p := sr.taskPointMap[taskId][pKey]

	mx, my := sr.getMove()
	robotgo.MoveClick(p.X+mx, p.Y+my)
}

func (sr *ScreenArea) ClickRect(taskId int32, rectKey string) {
	if !sr.IsExistsRect(taskId, rectKey) {
		return
	}

	rect := sr.taskRectMap[taskId][rectKey]
	mx, my := sr.getMove()
	robotgo.MoveClick(rect.X+mx+rect.W/2, rect.Y+my+rect.H/2)
}

//#endregion

func GetScreenAreaFromLocal(key string) (srObj *ScreenArea, exists bool) {
	srObj = NewScreenArea(key)
	srObj.currentRect = srObj.taskRectMap[0][model.Sys_Key_Rect_Game]

	localPath := fmt.Sprintf("./%s/%s", model.Sys_Con_Path_Config, key)

	infoList, err := ioutil.ReadDir(localPath)
	if err != nil {
		return nil, false
	}

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

		saveModel, exists := getTaskModel(key, taskId)
		if !exists {
			continue
		}

		for _, item := range saveModel.RectList {
			srObj.AddRect(taskId, item.Key, item.Rect)
		}

		for _, item := range saveModel.PointList {
			srObj.AddPoint(taskId, item.Key, image.Point{X: item.X, Y: item.Y})
		}
	}

	return srObj, true
}

func SaveScreenModel(sr *ScreenArea) error {

	saveModelMap := make(map[int32]*TaskSaveModel, 8)

	for taskId, subMap := range sr.taskRectMap {
		if _, exists := saveModelMap[taskId]; !exists {
			saveModelMap[taskId] = NewTaskSaveModel()
		}

		for key, r := range subMap {
			saveModelMap[taskId].RectList = append(saveModelMap[taskId].RectList, RectModel{Key: key, Rect: r})
		}
	}

	for taskId, subMap := range sr.taskPointMap {
		if _, exists := saveModelMap[taskId]; !exists {
			saveModelMap[taskId] = NewTaskSaveModel()
		}

		for key, p := range subMap {
			saveModelMap[taskId].PointList = append(saveModelMap[taskId].PointList, PointModel{Key: key, X: p.X, Y: p.Y})
		}
	}

	for taskId, sm := range saveModelMap {
		err := saveTaskModel(sr.Key, taskId, sm)
		if err != nil {
			return err
		}
	}

	return nil
}
