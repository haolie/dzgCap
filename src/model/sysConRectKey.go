package model

import (
	"sync"
)

const (
	// 主界面验证区域
	Sys_Key_Rect_Main_Check = "mainCheckRect"
	// 宴会要求按钮区域
	Sys_Key_Rect_Meeting_Join_Btn = "meetingJoinBtn"
	// 返回点击位置
	Sys_Key_Point_Back = "Point_Back"
	// 游戏区域
	Sys_Key_Rect_Game = "GameRect"

	// 嗣王宴 验证区域
	Sys_key_rect_Meeting_SiWangYan = "SiWangYan"
	// 君王宴 验证区域
	Sys_key_rect_Meeting_JunWangYan = "JunWangYan"
	// 亲王宴 验证区域
	Sys_key_rect_Meeting_QinWangYan = "QinWangYan"
	// 宴会列表 验证区域
	Sys_key_rect_Meeting_List = "MeetingList"
	// 宴会界面 人数奖励点击位置
	Sys_Key_Point_Meeting_GuestNumReward = "GuestNumReward"
	// 宴会人数奖励验证区域
	Sys_key_rect_Meeting_VerifyReward = "VerifyReward"
	// 宴会人数奖励领奖点击位置
	Sys_Key_Point_Meeting_DrawNumReward = "DrawNumReward"
	// 主界面 icon 位置
	Syc_Key_Point_Icon_Line = "IconLine"
	// 主界面 icon 位置
	Syc_Key_Point_Meeting_Sure = "MeetingSure"

	// 宴会结束界面
	Sys_key_rect_Meeting_End = "MeetingEnd"

	// 宴会列表 第一宴会点击位置
	Sys_Key_Point_Meeting_Item1 = "MeetingItem1"
	// 宴会列表 第er宴会点击位置
	Sys_Key_Point_Meeting_Item2 = "MeetingItem2"

	// 连续点击任务 点击位置
	Sys_key_Point_clickTask = "taskClickPoint"
)

var (
	initOnce = new(sync.Once)
	mapList  []map[string]interface{}
)

func appendList(key, val string) {
	dataMap := make(map[string]interface{}, 2)
	dataMap["Key"] = key
	dataMap["Val"] = val

	mapList = append(mapList, dataMap)
}

func GetKeyMap() []map[string]interface{} {
	initOnce.Do(func() {
		mapList = make([]map[string]interface{}, 0, 8)
		appendList(Sys_Key_Rect_Main_Check, "主界面")
		appendList(Sys_Key_Rect_Meeting_Join_Btn, "宴会邀请")
		appendList(Sys_Key_Point_Back, "返回点击")
		appendList(Sys_Key_Rect_Game, "游戏区域")

		appendList(Sys_key_rect_Meeting_List, "宴会列表")
		appendList(Sys_key_rect_Meeting_JunWangYan, "君王宴")
		appendList(Sys_key_rect_Meeting_SiWangYan, "嗣王宴")
		appendList(Sys_key_rect_Meeting_QinWangYan, "亲王宴")
		appendList(Sys_Key_Point_Meeting_GuestNumReward, "宴会-奖励进入")
		appendList(Sys_key_rect_Meeting_VerifyReward, "宴会-奖励验证")
		appendList(Sys_Key_Point_Meeting_DrawNumReward, "宴会-领取")
		appendList(Syc_Key_Point_Meeting_Sure, "宴会邀请确认")
		appendList(Sys_key_rect_Meeting_End, "宴会结束提示")
		appendList(Sys_Key_Point_Meeting_Item1, "宴会列表第一")
		appendList(Sys_Key_Point_Meeting_Item2, "宴会列表第二")

		appendList(Syc_Key_Point_Icon_Line, "主界面ICONLine")

		appendList(Sys_key_Point_clickTask, "连续点击位置")

	})

	return mapList
}
