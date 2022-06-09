package taskMeeting

import (
	"context"
	"time"

	. "dzgCap/src/model"
)

func init() {

}

type meeting struct {
	gArea      IGameArea
	contextObj context.Context
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

}

func (m *meeting) Stop() {
}
