package taskMeeting

import (
	"context"
	"fmt"
	"image"
	"sync/atomic"
	"time"

	"github.com/go-vgo/robotgo"

	"dzgCap/ConfigManger"
	"dzgCap/src/gameAreaModel"
	. "dzgCap/src/model"
	"dzgCap/src/task/taskCenter"
)

const (
	// 连续点击判断间隔
	con_click_span = time.Second
	// 返回失败最大次数
	con_max_errBack_times = 10
	// icon寻找 等待时间
	con_find_icon_wait = time.Minute * 10
)

func init() {
	taskCenter.RegisterTask(TaskEnum_Meeting, meetingCreater)
}

type meeting struct {
	gArea     IGameArea
	cancelFun func()

	status int32 // 状态 TaskEnum

	// 上次返回点击时间
	lastBackClickTime time.Time
	// 连续返回点击次数
	clickTimes int32

	// 主界面宴会Icon点击位置
	meetingIconP     *image.Point
	lastIconMissTime time.Time
}

func meetingCreater(ga IGameArea) ITask {
	return &meeting{
		gArea: ga,
	}
}

func (m *meeting) GetTaskType() TaskEnum {
	return TaskEnum_Meeting
}

func (m *meeting) GetStatus() TaskStatusEnum {
	return TaskStatusEnum_Unstart
}

func (m *meeting) GetStartTime() time.Time {
	return GetMinTime()
}

func (m *meeting) GetEndTime() time.Time {
	return GetMaxTime()
}

func (m *meeting) Start(contextObj context.Context, startTime, endTime time.Time, param ...interface{}) {
	for {
		status := atomic.LoadInt32(&m.status)
		if status == int32(TaskStatusEnum_Runing) {
			return
		}

		if atomic.CompareAndSwapInt32(&m.status, status, int32(TaskStatusEnum_Runing)) {
			break
		}
	}

	var ctx context.Context
	ctx, m.cancelFun = context.WithCancel(contextObj)
	go m.doMeetingTask(ctx)
}

func (m *meeting) Stop() {
	status := atomic.LoadInt32(&m.status)
	if status == int32(TaskStatusEnum_Unstart) {
		return
	}

	// 如果正在运行 先暂停
	if status == int32(TaskStatusEnum_Runing) {
		m.cancelFun()
	}

	atomic.StoreInt32(&m.status, int32(TaskStatusEnum_Unstart))
}

// 异步宴会
func (m *meeting) doMeetingTask(ctx context.Context) {
	timeCh := time.After(5000 * time.Millisecond)
	drawCh := time.After(20 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return
		case <-timeCh:
			m.joinMeeting()
			timeCh = time.After(300* time.Millisecond)
		case <-drawCh:
			if ConfigManger.GetMeetingRewardTime() <= 0 {
				break
			}

			m.drawMeetingReward()
			drawCh = time.After(time.Second * time.Duration(ConfigManger.GetMeetingRewardTime()))
		}
	}
}

func (m *meeting) joinMeeting() {
	if !m.isMeetingJoinView() {
		if !m.gArea.IsHome() {
			m.gArea.GoBack()
		}

		return
	}

	// 点击要求按钮
	m.gArea.ClickRectKey(Sys_Key_Rect_Meeting_Join_Btn, "")

	for i:=0;i<8;i++{
		if m.isMeetingJoinView(){
			break
		}

		robotgo.MilliSleep(100)
	}

	robotgo.MilliSleep(400)

	// 宴会要求已过期 todo
	if m.isMeetingJoinView() {
		m.gArea.ClickPointKey(Syc_Key_Point_Meeting_Sure, "")
		return
	}

	// 点击参宴按钮
	m.gArea.ClickRectKey(Sys_Key_Rect_Meeting_Join_Btn, "")
	robotgo.MilliSleep(800)

	// 返回操作
	m.gArea.GoBack()
	robotgo.MilliSleep(500)
}

// 判断是否是宴会界面
func (m *meeting) isMeetingJoinView() bool {
	return m.gArea.VerifyRect(m.GetTaskType(), Sys_Key_Rect_Meeting_Join_Btn)
}

func (m *meeting) drawMeetingReward() {
	// 判断宴会入库是否有效
	if m.meetingIconP == nil {
		// 无效则查找 查找失败直接返回
		if !m.findMeetingIcon() {
			return
		}
	}

	// 返回主页面
	if m.gArea.ToHome() != nil {
		panic(fmt.Errorf("screen err"))
	}

	// 点击宴会icon
	m.gArea.ClickPoint(m.meetingIconP.X, m.meetingIconP.Y, "")
	time.Sleep(Sys_Con_jump_Waite * 2)

	// 抢占
	if m.grab() {
		return
	}

	// 如果是宴会列表 点击第宴会
	if m.isMeetingList() {
		m.gArea.ClickPointKey(Sys_Key_Point_Meeting_Item1, "")
		time.Sleep(Sys_Con_jump_Waite * 2)

		// 抢占
		if m.grab() {
			return
		}
	}

	// 如果是宴会界面 开始领奖操作
	if m.isMeeting() {
		m.rewardDrawFn()
	} else {
		m.meetingIconP = nil
		return
	}

	for i := 0; i < 5; i++ {
		m.gArea.GoBack()
		time.Sleep(Sys_Con_jump_Waite)

		// 抢占
		if m.grab() {
			return
		}

		if m.isMeetingList() {
			m.gArea.ClickPointKey(Sys_Key_Point_Meeting_Item2, "")
			time.Sleep(Sys_Con_jump_Waite * 2)

			// 抢占
			if m.grab() {
				return
			}

			// 如果是宴会界面 开始领奖操作
			if m.isMeeting() {
				m.rewardDrawFn()
				return
			}

			break
		}

		// 返回主页面
		if m.gArea.ToHome() != nil {
			panic(fmt.Errorf("screen err"))
		}
	}
}

// 查找宴会icon
func (m *meeting) findMeetingIcon() bool {
	// 间隔时间内 不查找
	if m.lastIconMissTime.Add(con_find_icon_wait).After(time.Now()) {
		return false
	}

	m.lastIconMissTime = time.Now()

	// 返回主页面
	if m.gArea.ToHome() != nil {
		panic(fmt.Errorf("screen err"))
	}

	// 抢占
	if m.grab() {
		return false
	}

	// 获取游戏区域
	gameRect, exists := gameAreaModel.GetRect(m.gArea.GetKey(), 0, Sys_Key_Rect_Game)
	if !exists {
		return false
	}

	// icon 点位
	p, exists := gameAreaModel.GetPoint(m.gArea.GetKey(), 0, Syc_Key_Point_Icon_Line)
	if !exists {
		return false
	}

	// 左右宽度
	padding := 2
	iconW := (gameRect.W - padding*2) / 6
	for i := 0; i < 6; i++ {
		if m.gArea.ToHome() != nil {
			panic(fmt.Errorf("screen err"))
		}

		// 抢占
		if m.grab() {
			return false
		}

		px := gameRect.X + iconW*i + padding + iconW/2

		// 点击Icon
		m.gArea.ClickPoint(px, p.Y, "")
		time.Sleep(time.Second * 2)
		// 抢占
		if m.grab() {
			return false
		}

		//找到宴会列表 返回true|| 找到宴会 返回true
		if m.isMeetingList() || m.isMeeting() {
			m.meetingIconP = &image.Point{
				X: px,
				Y: p.Y,
			}
			return true
		}
	}

	return false
}

// 领奖操作
func (m *meeting) rewardDrawFn() {
	// 点击宴会人数奖励Icon
	m.gArea.ClickPointKey(Sys_Key_Point_Meeting_GuestNumReward, "")
	time.Sleep(Sys_Con_jump_Waite)

	// 抢占
	if m.grab() {
		return
	}

	// 领奖
	m.gArea.ClickPointKey(Sys_Key_Point_Meeting_DrawNumReward, "")
	time.Sleep(Sys_Con_jump_Waite)
}

// 判断是否是宴会列表界面
func (m *meeting) isMeetingList() bool {
	return m.gArea.VerifyRect(m.GetTaskType(), Sys_key_rect_Meeting_List)
}

// 判断是否是宴会界面
func (m *meeting) isMeeting() bool {
	if m.gArea.VerifyRect(m.GetTaskType(), Sys_key_rect_Meeting_SiWangYan) {
		return true
	}

	if m.gArea.VerifyRect(m.GetTaskType(), Sys_key_rect_Meeting_QinWangYan) {
		return true
	}

	if m.gArea.VerifyRect(m.GetTaskType(), Sys_key_rect_Meeting_JunWangYan) {
		return true
	}

	return false
}

// 宴会抢占
func (m *meeting) grab() bool {
	if !m.isMeetingJoinView() {
		return false
	}

	m.lastIconMissTime = time.Now().Add(-con_find_icon_wait * 2)
	return true
}
