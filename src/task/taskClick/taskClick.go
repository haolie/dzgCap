package taskClick

import (
	"context"
	"fmt"
	"image"
	"sync/atomic"
	"time"

	"dzgCap/Loger"
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
	con_click_times    = 180
)

func init() {
	taskCenter.RegisterTask(TaskEnum_Click, clickTaskCreater)
}

type clickTask struct {
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

func clickTaskCreater(ga IGameArea) ITask {
	return &clickTask{
		gArea: ga,
	}
}

func (m *clickTask) GetTaskType() TaskEnum {
	return TaskEnum_Click
}

func (m *clickTask) GetStatus() TaskStatusEnum {
	return TaskStatusEnum_Unstart
}

func (m *clickTask) GetStartTime() time.Time {
	return GetMinTime()
}

func (m *clickTask) GetEndTime() time.Time {
	return GetMaxTime()
}

func (m *clickTask) Start(contextObj context.Context, startTime, endTime time.Time, param ...interface{}) {
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

func (m *clickTask) Stop() {
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
func (m *clickTask) doMeetingTask(ctx context.Context) {
	timeCh := time.After(2000 * time.Millisecond)

	var num int

	for {
		select {
		case <-ctx.Done():
			return
		case <-timeCh:
			if num > con_click_times {
				m.Stop()
			}
			Loger.LogInfo(fmt.Sprintf("click num:%d \n",num))
			 m.gArea.ClickPointKey(Sys_key_Point_clickTask, "")
			num++
			timeCh = time.After(con_click_span)
		}
	}
}
