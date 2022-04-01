package model

type IPageView interface {
	GetId() int32
	GetName() string
	GetEnterPointList()
	GoToView(key string) bool
	IsInView() bool
}
