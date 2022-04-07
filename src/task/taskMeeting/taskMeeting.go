package taskMeeting

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/go-vgo/robotgo"

	"dzgCap/src/PageView/PageViewCenter"
	"dzgCap/src/ScreenModel"
	. "dzgCap/src/model"
	"dzgCap/src/task/taskCenter"
)

func init() {
	// 组成宴会加入按钮（验证）区域
	ScreenModel.RegisterModelKey(int32(TaskEnum_Meeting), int32(ScreenModel.ModelTypeEnum_Rect), Sys_Key_Rect_Meeting_Join_Btn)
	// 组成宴会加入按钮验证图片
	ScreenModel.RegisterModelKey(int32(TaskEnum_Meeting), int32(ScreenModel.ModelTypeEnum_Image), Sys_Key_Rect_Meeting_Join_Btn)

	taskCenter.RegisterTask(new(meetingTask))
}

type meetingTask struct {
	startTime   time.Time
	endTime     time.Time
	curPv       IPageView
	closeSignal chan struct{}
	status      int32
}

func (m *meetingTask) GetKey() TaskEnum {
	return TaskEnum_Meeting
}

func (m *meetingTask) GetStatus() TaskStatusEnum {
	return TaskStatusEnum(m.status)
}

func (m *meetingTask) GetStartTime() time.Time {
	return m.startTime
}
func (m *meetingTask) GetEndTime() time.Time {
	return m.endTime
}
func (m *meetingTask) Start() {
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

func (m *meetingTask) Pause() {
	status := atomic.LoadInt32(&m.status)
	if status != int32(TaskStatusEnum_Runing) {
		return
	}

	if atomic.CompareAndSwapInt32(&m.status, status, int32(TaskStatusEnum_Pause)) {
		close(m.closeSignal)
	}
}

func (m *meetingTask) Release() {

}

func (m *meetingTask) joinMeeting(closeCh chan struct{}) {

	timeCh := time.After(500)

	for {
		select {
		case <-closeCh:
			return
		case <-timeCh:
			m.doJoin()
			timeCh = time.After(500)
		}
	}

}

func (m *meetingTask) doJoin() {
	if !m.isMeetingJoinView() {
		if !PageViewCenter.IsMainView() {
			PageViewCenter.GoBack()
		}

		return
	}

	ScreenModel.GetCurrentScreenArea().ClickRect(int32(m.GetKey()), Sys_Key_Rect_Meeting_Join_Btn)
	robotgo.MilliSleep(1000)
	ScreenModel.GetCurrentScreenArea().ClickRect(int32(m.GetKey()), Sys_Key_Rect_Meeting_Join_Btn)
	robotgo.MilliSleep(500)
}

func (m *meetingTask) isMeetingJoinView() bool {
	canJoin, err := ScreenModel.GetCurrentScreenArea().CompareRectToCash(int32(m.GetKey()), Sys_Key_Rect_Meeting_Join_Btn)
	if err != nil {
		fmt.Printf("verify meetingJoin faild err:%v\n", err)
		return false
	}

	return canJoin
}
