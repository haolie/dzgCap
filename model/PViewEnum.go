package model

type PViewEnum int32

func (p PViewEnum) ToString() string {
	return pvMap[p]
}

const (
	// 主页面
	PViewEnum_Main PViewEnum = iota
	// 宴会列表
	PViewEnum_MeetingList
	// 宴会界面
	PViewEnum_MeetingView
	// 宴会奖励界面
	PViewEnum_MeettingRewardView
)

var pvMap = map[PViewEnum]string{
	PViewEnum_Main:               "mainView",
	PViewEnum_MeetingList:        "MeetingList",
	PViewEnum_MeetingView:        "MeetingView",
	PViewEnum_MeettingRewardView: "MeetingRewardView",
}

func PVVerify(pv PViewEnum) bool {
	_, exists := pvMap[pv]
	return exists
}
