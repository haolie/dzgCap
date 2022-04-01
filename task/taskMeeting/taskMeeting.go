package taskMeeting

import (
	"fmt"
	"image"
	"sync/atomic"
	"time"

	"github.com/go-vgo/robotgo"

	"dzgCap/capMager"
	. "dzgCap/model"
)

var (
	rectMeetingBtn      = *NewRect(0, 0, 0, 0)
	pointMeetingJoinBtn = image.Point{X: 0, Y: 0}
)

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
	go joinMeeting(m.closeSignal)
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

func joinMeeting(closeCh chan struct{}) {

	timeCh := time.After(500)

	for {
		select {
		case <-closeCh:
			return
		case <-timeCh:
			doJoin()
			timeCh = time.After(500)
		}
	}

}

func doJoin() {
	if !isMeetingJoinView() {
		return
	}

	capMager.ClickRect(rectMeetingBtn)
	robotgo.MilliSleep(1000)
	capMager.ClickPoint(pointMeetingJoinBtn.X, pointMeetingJoinBtn.Y)
	robotgo.MilliSleep(200)
}

func isMeetingJoinView() bool {
	canJoin, err := capMager.CompareRectToCash(rectMeetingBtn, Sys_Key_Rect_Meeting_Join_Btn)
	if err != nil {
		fmt.Printf("verify meetingJoin faild err:%v\n", err)
		return false
	}

	return canJoin
}
