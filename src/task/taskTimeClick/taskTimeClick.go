package taskTimeClick

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"dzgCap/Loger"
	"dzgCap/src/ScreenModel"
	. "dzgCap/src/model"
	"dzgCap/src/task/taskCenter"
)

func init() {
	// 注册任务
	taskCenter.RegisterTask(newClickTask())
}

type clickTask struct {
	startTime   time.Time     // 任务开始时间
	endTime     time.Time     // 任务结束时间
	curPv       IPageView     // 当前页面
	closeSignal chan struct{} // 关闭信号
	status      int32         // 状态 TaskEnum
	maxNum      int32         // 最大点击次数（0无效）

	clickCh  chan struct{} // 点击通信
	clickNum int32
}

func newClickTask() *clickTask {
	return &clickTask{
		startTime: time.Now(),
		endTime:   time.Now().Add(time.Minute),
		curPv:     nil,
		status:    0,
	}
}

// 返回任务键值
func (m *clickTask) GetKey() TaskEnum {
	return TaskEnum_Click
}

// 返回任务当前状态
func (m *clickTask) GetStatus() TaskStatusEnum {
	return TaskStatusEnum(atomic.LoadInt32(&m.status))
}

// 返回任务开始时间
func (m *clickTask) GetStartTime() time.Time {
	return m.startTime
}

// 返回任务结束时间
func (m *clickTask) GetEndTime() time.Time {
	return m.endTime
}

// 开始任务
func (m *clickTask) Start(param interface{}) {
	for {
		status := atomic.LoadInt32(&m.status)
		if status != int32(TaskStatusEnum_Unstart) {
			return
		}

		if atomic.CompareAndSwapInt32(&m.status, status, int32(TaskStatusEnum_Runing)) {
			break
		}
	}

	m.clickNum = 0
	numStr, success := param.(string)
	if success {
		fmt.Println(param)
		n, err := strconv.Atoi(numStr)
		if err == nil {
			m.maxNum = int32(n)
		}
	}

	Loger.LogInfo(fmt.Sprintf("taskClick started num:%d", m.maxNum))

	m.closeSignal = make(chan struct{})
	m.clickCh = make(chan struct{})

	m.endTime = time.Now().Add(time.Minute)

	go m.doing()

	go func() {
		select {
		case <-m.closeSignal:
			return
		case <-time.After(m.endTime.Sub(time.Now())):
			m.Stop()
		}
	}()
}

// 停止任务
func (m *clickTask) Stop() {
	status := atomic.LoadInt32(&m.status)
	if status == int32(TaskStatusEnum_Unstart) {
		return
	}
	//
	//// 如果正在运行 先暂停
	//if status == int32(TaskStatusEnum_Runing) {
	//	m.Pause()
	//}

	atomic.StoreInt32(&m.status, int32(TaskStatusEnum_Unstart))

	close(m.closeSignal)
	close(m.clickCh)
	Loger.LogInfo("taskClick stopped")
}

// 暂停任务
func (m *clickTask) Pause() {
	//status := atomic.LoadInt32(&m.status)
	//if status != int32(TaskStatusEnum_Runing) {
	//	return
	//}
	//
	//if atomic.CompareAndSwapInt32(&m.status, status, int32(TaskStatusEnum_Pause)) {
	//	close(m.closeSignal)
	//}
}

// 释放任务
func (m *clickTask) Release() {

}

func (m *clickTask) doing() {
	go m.clickFn()
	for {
		select {
		case <-m.closeSignal:
			return
		case <-m.clickCh:
			go m.clickFn()
		}
	}

}

func (m *clickTask) clickFn() {
	p, exists := ScreenModel.GetCurrentScreenArea().GetPoint(int32(m.GetKey()), Sys_key_Point_clickTask)
	if exists {
		m.clickNum += 1
		// 点击宴会icon
		ScreenModel.GetCurrentScreenArea().ClickPoint(p.X, p.Y)
		Loger.LogInfo(fmt.Sprintf("click x:%d,y:%d  num:%d", p.X, p.Y, m.clickNum))
	} else {
		Loger.LogInfo("miss point")
	}

	if m.maxNum > 0 && m.clickNum >= m.maxNum {
		m.Stop()
	} else {
		time.Sleep(time.Second)
		if m.GetStatus() == TaskStatusEnum_Runing {
			m.clickCh <- struct{}{}
		}
	}
}
