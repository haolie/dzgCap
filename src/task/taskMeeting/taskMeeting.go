// ************************************
// @package: taskMeeting
// @description: 宴会任务
// @author:
// @revision history:
// @create date: 2022-04-22 09:48:24
// ************************************

package taskMeeting

import (
	"fmt"
	"image"
	"sync/atomic"
	"time"

	"github.com/go-vgo/robotgo"

	"dzgCap/src/PageView/PageViewCenter"
	"dzgCap/src/ScreenModel"
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
	// 组成宴会加入按钮（验证）区域
	ScreenModel.RegisterModelKey(int32(TaskEnum_Meeting), int32(ScreenModel.ModelTypeEnum_Rect), Sys_Key_Rect_Meeting_Join_Btn)
	// 组成宴会加入按钮验证图片
	ScreenModel.RegisterModelKey(int32(TaskEnum_Meeting), int32(ScreenModel.ModelTypeEnum_Image), Sys_Key_Rect_Meeting_Join_Btn)
	// 注册任务
	taskCenter.RegisterTask(newMeetingTask())
}

// 宴会任务结构体
type meetingTask struct {
	startTime   time.Time     // 任务开始时间
	endTime     time.Time     // 任务结束时间
	curPv       IPageView     // 当前页面
	closeSignal chan struct{} // 关闭信号
	status      int32         // 状态 TaskEnum

	// 上次返回点击时间
	lastBackClickTime time.Time
	// 连续返回点击次数
	clickTimes int32

	// 主界面宴会Icon点击位置
	meetingIconP     *image.Point
	lastIconMissTime time.Time
}

func newMeetingTask() *meetingTask {
	return &meetingTask{
		startTime:         time.Now(),
		endTime:           time.Now(),
		curPv:             nil,
		closeSignal:       nil,
		status:            0,
		lastBackClickTime: time.Now(),
		clickTimes:        0,
		meetingIconP:      nil,
		lastIconMissTime:  time.Now().Add(-2 * con_find_icon_wait),
	}
}

// 返回任务键值
func (m *meetingTask) GetKey() TaskEnum {
	return TaskEnum_Meeting
}

// 返回任务当前状态
func (m *meetingTask) GetStatus() TaskStatusEnum {
	return TaskStatusEnum(atomic.LoadInt32(&m.status))
}

// 返回任务开始时间
func (m *meetingTask) GetStartTime() time.Time {
	return m.startTime
}

// 返回任务结束时间
func (m *meetingTask) GetEndTime() time.Time {
	return m.endTime
}

// 开始任务
func (m *meetingTask) Start(param interface{}) {
	for {
		status := atomic.LoadInt32(&m.status)
		if status == int32(TaskStatusEnum_Runing) {
			return
		}

		if atomic.CompareAndSwapInt32(&m.status, status, int32(TaskStatusEnum_Runing)) {
			break
		}
	}

	m.closeSignal = make(chan struct{})
	go m.joinMeeting(m.closeSignal)
}

// 停止任务
func (m *meetingTask) Stop() {
	status := atomic.LoadInt32(&m.status)
	if status == int32(TaskStatusEnum_Unstart) {
		return
	}

	// 如果正在运行 先暂停
	if status == int32(TaskStatusEnum_Runing) {
		m.Pause()
	}

	atomic.StoreInt32(&m.status, int32(TaskStatusEnum_Unstart))
}

// 暂停任务
func (m *meetingTask) Pause() {
	status := atomic.LoadInt32(&m.status)
	if status != int32(TaskStatusEnum_Runing) {
		return
	}

	if atomic.CompareAndSwapInt32(&m.status, status, int32(TaskStatusEnum_Pause)) {
		close(m.closeSignal)
	}
}

// 释放任务
func (m *meetingTask) Release() {

}

// 异步宴会
func (m *meetingTask) joinMeeting(closeCh chan struct{}) {

	timeCh := time.After(500)
	//drawCh := time.After(2000)

	for {
		select {
		case <-closeCh:
			return
		case <-timeCh:
			m.doJoin()
			timeCh = time.After(500)
			//case <-drawCh:
			//	m.drawMeetingReward()
			//	drawCh = time.After(time.Minute)
		}
	}

}

// 宴会加入
func (m *meetingTask) doJoin() {
	// 不是宴会要求界面 执行返回操作 直至主页面
	if !m.isMeetingJoinView() {
		// 判断主页面
		if !PageViewCenter.IsMainView() {
			// 返回操作
			PageViewCenter.GoBack()

			// 记录返回操作时间与次数 判断连续返回失败   连续返回失败panic
			{
				if m.lastBackClickTime.Add(con_click_span).After(time.Now()) {
					m.clickTimes += 1
				} else {
					m.clickTimes = 0
				}

				m.lastBackClickTime = time.Now()

				if m.clickTimes > con_max_errBack_times {
					err := ScreenModel.GetCurrentScreenArea().FreshArea()
					if err != nil {
						panic(err)
					}
				}
			}

		}
		return
	}

	// 点击要求按钮
	ScreenModel.GetCurrentScreenArea().ClickKeyRect(int32(m.GetKey()), Sys_Key_Rect_Meeting_Join_Btn)
	robotgo.MilliSleep(800)

	if m.isMeetingJoinView() {
		ScreenModel.GetCurrentScreenArea().ClickKeyRect(int32(m.GetKey()), Sys_Key_Rect_Meeting_Join_Btn)
		robotgo.MilliSleep(800)

		if m.isMeetingJoinView() {
			// 宴会要求已过期 todo

		}
	}

	// 点击参宴按钮
	ScreenModel.GetCurrentScreenArea().ClickKeyRect(int32(m.GetKey()), Sys_Key_Rect_Meeting_Join_Btn)
	robotgo.MilliSleep(800)
	// 返回操作
	PageViewCenter.GoBack()
	robotgo.MilliSleep(500)
}

// 判断是否是宴会界面
func (m *meetingTask) isMeetingJoinView() bool {
	// 对比宴会邀请按钮区域图形
	canJoin, err := ScreenModel.GetCurrentScreenArea().CompareRectToCash(int32(m.GetKey()), Sys_Key_Rect_Meeting_Join_Btn)
	if err != nil {
		fmt.Printf("verify meetingJoin faild err:%v\n", err)
		return false
	}

	return canJoin
}

// 领取宴会奖励
func (m *meetingTask) drawMeetingReward() bool {
	// 判断宴会入库是否有效
	if m.meetingIconP == nil {
		// 无效则查找 查找失败直接返回
		if !m.findMeetingIcon() {
			return false
		}
	}

	// 返回主页面
	if !PageViewCenter.GoToMainView() {
		panic(fmt.Errorf("screen err"))
	}

	// 点击宴会icon
	ScreenModel.GetCurrentScreenArea().ClickPoint(m.meetingIconP.X, m.meetingIconP.Y)
	time.Sleep(Sys_Con_jump_Waite)

	// 抢占
	if m.grab() {
		return false
	}

	// 如果是宴会列表 点击第宴会
	if m.isMeetingList() {
		ScreenModel.GetCurrentScreenArea().ClickPointKey(1, Sys_Key_Point_Meeting_Item1)
		time.Sleep(Sys_Con_jump_Waite)

		// 抢占
		if m.grab() {
			return false
		}
	}

	// 如果是宴会界面 开始领奖操作
	if m.isMeeting() {
		m.rewardDrawFn()
		return true
	}

	return false
}

// 领奖操作
func (m *meetingTask) rewardDrawFn() {

	// 点击宴会人数奖励Icon
	ScreenModel.GetCurrentScreenArea().ClickPointKey(1, Sys_Key_Point_Meeting_GuestNumReward)
	time.Sleep(Sys_Con_jump_Waite)

	// 抢占
	if m.grab() {
		return
	}

	// 领奖
	ScreenModel.GetCurrentScreenArea().ClickPointKey(1, Sys_Key_Point_Meeting_DrawNumReward)
	time.Sleep(Sys_Con_jump_Waite)
}

// 查找宴会icon
func (m *meetingTask) findMeetingIcon() bool {
	// 间隔时间内 不查找
	if m.lastIconMissTime.Add(con_find_icon_wait).After(time.Now()) {
		return false
	}

	m.lastBackClickTime = time.Now()

	// 返回主页面
	if !PageViewCenter.GoToMainView() {
		panic(fmt.Errorf("screen err"))
	}

	// 抢占
	if m.grab() {
		return false
	}

	// 获取游戏区域
	gameRect, exists := ScreenModel.GetCurrentScreenArea().GetRect(0, Sys_Key_Rect_Game)
	if !exists {
		return false
	}

	// icon 点位
	p, exists := ScreenModel.GetCurrentScreenArea().GetPoint(0, Syc_Key_Point_Icon_Line)
	if !exists {
		return false
	}

	// 左右宽度
	padding := 2
	iconW := (gameRect.W - padding*2) / 6
	for i := 0; i < 6; i++ {
		if !PageViewCenter.GoToMainView() {
			panic(fmt.Errorf("screen err"))
		}

		// 抢占
		if m.grab() {
			return false
		}

		px := gameRect.X + iconW*i + padding + iconW/2

		// 点击Icon
		ScreenModel.GetCurrentScreenArea().ClickPoint(px, p.Y)
		time.Sleep(Sys_Con_jump_Waite)
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

// 判断是否是宴会列表界面
func (m *meetingTask) isMeetingList() bool {
	sm, err := ScreenModel.GetCurrentScreenArea().CompareRectToCash(1, Sys_key_rect_Meeting_List)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return sm
}

// 判断是否是宴会界面
func (m *meetingTask) isMeeting() bool {
	sm, err := ScreenModel.GetCurrentScreenArea().CompareRectToCash(1, Sys_key_rect_Meeting_SiWangYan)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if sm {
		return true
	}

	sm, err = ScreenModel.GetCurrentScreenArea().CompareRectToCash(1, Sys_key_rect_Meeting_QinWangYan)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if sm {
		return true
	}

	sm, err = ScreenModel.GetCurrentScreenArea().CompareRectToCash(1, Sys_key_rect_Meeting_JunWangYan)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return sm
}

// 宴会抢占
func (m *meetingTask) grab() bool {
	if !m.isMeetingJoinView() {
		return false
	}

	m.lastBackClickTime = time.Now().Add(-con_find_icon_wait * 2)
	return true
}
