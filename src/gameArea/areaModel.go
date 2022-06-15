package gameArea

import (
	"context"
	"image"
	"sync"
	"time"

	"dzgCap/src/gameAreaModel"
	"dzgCap/src/imageTool"
	. "dzgCap/src/model"
	"dzgCap/src/task/taskCenter"
)

const (
	con_Start_Back_Count   = 5
	con_Start_ReVerifyTime = 300 * time.Millisecond
)

var _ IGameArea = (*GameArea)(nil)

type GameArea struct {
	gameModelKey string
	rect         Rect
	task         ITask
	gTable       IGameTable
	startOnce    *sync.Once
	homeImg      image.Image
}

func (ga *GameArea) GetKey() string {
	return ga.gameModelKey
}

func (ga *GameArea) ClickRectKey(key string, checkKey string) bool {
	ga.clickRectFn(ga.task.GetTaskType(), key)

	return ga.clickVerify(checkKey)
}

func (ga *GameArea) ClickPointKey(key string, checkKey string) bool {
	ga.clickPointKey(ga.task.GetTaskType(), key)

	return ga.clickVerify(checkKey)
}

func (ga *GameArea) ClickPoint(x, y int, checkKey string) bool {
	ga.click(x, y)

	return ga.clickVerify(checkKey)
}

func (ga *GameArea) VerifyRect(taskType TaskEnum, key string) bool {
	return ga.verifyRect(taskType, key)
}

func (ga *GameArea) StartTask(ctx context.Context, taskType TaskEnum) error {
	// 停止当前任务
	ga.Stop()

	// 返回主页面
	ga.clickBack(con_Start_Back_Count)

	// 获取主页标识图片
	{
		r, exists := gameAreaModel.GetRect(ga.gameModelKey, 0, Sys_Key_Rect_Main_Check)
		if !exists {
			panic("can not find rect:" + Sys_Key_Point_Back)
		}

		r = r.Move(ga.getMove())
		ga.homeImg = imageTool.CapScreen(r)
	}

	// 创建并开始任务
	ga.task = taskCenter.CreateTask(taskType, ga)
	ga.task.Start(ctx, GetMinTime(), GetMaxTime())

	return nil
}

func (ga *GameArea) Stop() {
	if ga.task != nil {
		ga.task.Stop()
		ga.task = nil
	}

}

func (ga *GameArea) GetStatus() (status TaskStatusEnum, taskType TaskEnum) {
	if ga.task == nil {
		status = TaskStatusEnum_Unstart
		return
	}

	return ga.task.GetStatus(), ga.task.GetTaskType()
}

func (ga *GameArea) ToHome() error {
	if ga.IsHome() {
		return nil
	}

	ga.clickBack(con_Start_Back_Count)
	if ga.IsHome() {
		return nil
	}

	// 对比宴会邀请按钮区域图形
	canJoin := ga.verifyRect(TaskEnum_Meeting, Sys_Key_Rect_Meeting_Join_Btn)
	// 宴会要求界面  点击确认按钮
	if canJoin {
		ga.clickPointKey(TaskEnum_Meeting, Syc_Key_Point_Meeting_Sure)
		time.Sleep(Sys_Con_jump_Waite)
	}

	ga.clickBack(con_Start_Back_Count)
	if ga.IsHome() {
		return nil
	}

	isEnd := ga.verifyRect(TaskEnum_Meeting, Sys_key_rect_Meeting_End)
	if isEnd {
		ga.clickRectFn(TaskEnum_Meeting, Sys_key_rect_Meeting_End)
		time.Sleep(Sys_Con_jump_Waite)
	}

	ga.clickBack(con_Start_Back_Count)
	if ga.IsHome() {
		return nil
	}

	return nil
}

func (ga *GameArea) IsHome() bool {
	r, exists := gameAreaModel.GetRect(ga.gameModelKey, 0, Sys_Key_Rect_Main_Check)
	if !exists {
		return false
	}

	r = r.Move(ga.getMove())
	img := imageTool.CapScreen(r)

	return imageTool.CompareImage(ga.homeImg, img)
}

func (ga *GameArea) GoBack() {
	ga.clickBack(1)
}

func (ga *GameArea) goHome() {
	p, exists := gameAreaModel.GetPoint(ga.gameModelKey, 0, Sys_Key_Point_Back)
	if !exists {
		return
	}

	ga.click(p.X, p.Y)
}

func (ga *GameArea) clickBack(num int) {
	for i := 0; i < num; i++ {
		ga.goHome()
		time.Sleep(800 * time.Millisecond)
	}
}

func (ga *GameArea) clickRectFn(taskType TaskEnum, rectKey string) {
	r, exists := gameAreaModel.GetRect(ga.gameModelKey, int32(taskType), rectKey)
	if !exists {
		return
	}

	ga.click(r.X+r.W/2, r.Y+r.H/2)
}

func (ga *GameArea) clickPointKey(taskType TaskEnum, pointKey string) {
	p, exists := gameAreaModel.GetPoint(ga.gameModelKey, int32(taskType), pointKey)
	if !exists {
		return
	}

	ga.click(p.X, p.Y)
}

func (ga *GameArea) click(x, y int) {
	mx, my := ga.getMove()
	ga.gTable.Click(x+mx, y+my)
}

func (ga *GameArea) getMove() (x, y int) {
	r, exists := gameAreaModel.GetRect(ga.gameModelKey, 0, Sys_Key_Rect_Game)
	if !exists {
		return
	}

	return ga.rect.X - r.X, ga.rect.Y - r.Y
}

func (ga *GameArea) verifyRect(taskType TaskEnum, rectKey string) bool {
	mx, my := ga.getMove()

	return gameAreaModel.VerifyRect(ga.gameModelKey, int32(taskType), rectKey, mx, my)
}

func (ga *GameArea) clickVerify(rectKey string) bool {
	if len(rectKey) == 0 {
		return true
	}

	for i := 0; i <= 7; i++ {
		if ga.VerifyRect(ga.task.GetTaskType(), rectKey) {
			return true
		}

		time.Sleep(con_Start_ReVerifyTime)
	}

	return false
}
