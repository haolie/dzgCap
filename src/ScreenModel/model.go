package ScreenModel

type PointModel struct {
	Key string
	X   int
	Y   int
}

type RectModel struct {
	Key string
	X   int
	Y   int
	W   int
	H   int
}

type TaskSaveModel struct {
	PointList []PointModel
	RectList  []RectModel

	pointMap map[string]PointModel `json:" - "`
	rectMap  map[string]RectModel  `json:" - "`
}

func NewTaskSaveModel() *TaskSaveModel {
	return &TaskSaveModel{
		PointList: nil,
		RectList:  nil,
		pointMap:  make(map[string]PointModel, 4),
		rectMap:   make(map[string]RectModel, 4),
	}
}

func (tsm *TaskSaveModel) AddPoint(p PointModel) {
	tsm.pointMap[p.Key] = p
}

func (tsm *TaskSaveModel) AddRect(r RectModel) {
	tsm.rectMap[r.Key] = r
}

func (tsm *TaskSaveModel) IsExistsPoint(key string) bool {
	_, exists := tsm.pointMap[key]
	return exists
}

func (tsm *TaskSaveModel) IsExistsRect(key string) bool {
	_, exists := tsm.rectMap[key]
	return exists
}
