package model

type TaskEnum int32

const (
	TaskEnum_Meeting TaskEnum = iota + 1
	TaskEnum_TradeWoar
)

var taskEnumMap = map[TaskEnum]string{
	TaskEnum_Meeting:   "宴会",
	TaskEnum_TradeWoar: "商战",
}

func TaskEnumVerify(taskEnum TaskEnum) bool {
	_, exists := taskEnumMap[taskEnum]
	return exists
}
