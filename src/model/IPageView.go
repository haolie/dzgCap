package model

type IPageView interface {

	// GetId
    // @description: 返回界面id
    // parameter:
    // return:
    //		@int32:
	GetId() int32

	// GetName
    // @description: 返回界面名称
    // parameter:
    // return:
    //		@string:
	GetName() string

	// GetEnterPointList
    // @description: 返回入口列表
    // parameter:
    // return:
	GetEnterPointList()

	// GoToView
    // @description: 进入指定界面
    // parameter:
    //		@key:
    // return:
    //		@bool:
	GoToView(key string) bool

	// IsInView
    // @description: 是否在当前页面
    // parameter:
    // return:
    //		@bool:
	IsInView() bool
}
