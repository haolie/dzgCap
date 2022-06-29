package common

import (
	"time"

	"dzgCap/tools"
)

const (
	Con_HS_Status_Fail    = 0
	Con_HS_Status_Success = 200
)

type HSResponse struct {
	status  int
	content interface{}
	time    int64
	timeStr string
}

func baseDataMap() map[string]interface{} {
	dataMap := make(map[string]interface{}, 5)
	dataMap["status"] = Con_HS_Status_Success
	dataMap["content"] = ""
	dataMap["data"] = "11111"
	dataMap["time"] = time.Now().Unix()
	dataMap["timeStr"] = tools.ToDateTimeStr(time.Now())

	return dataMap
}

func CreateErrHSResponse(errInfo string) map[string]interface{} {
	dataMap := baseDataMap()
	dataMap["status"] = Con_HS_Status_Fail
	dataMap["content"] = errInfo

	return dataMap
}

func CreateSuccessHSResponse(info interface{}) map[string]interface{} {

	dataMap := baseDataMap()
	dataMap["status"] = Con_HS_Status_Success
	dataMap["content"] = info

	return dataMap
}
