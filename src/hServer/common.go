package hServer

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"dzgCap/src/model"
)

func getRectParam(ctx iris.Context) (key string, taskId int32, r model.Rect, err error) {
	key = ctx.URLParam(con_Params_Key)
	if len(key) == 0 {
		ctx.WriteString("need key")
		return
	}

	taskId32, err := ctx.URLParamInt(con_Params_TaskType)
	if err != nil {
		ctx.WriteString("need taskId")
		return
	}

	taskId = int32(taskId32)

	r.X, err = ctx.URLParamInt(con_Params_X)
	if err != nil {
		ctx.WriteString("need x")
		return
	}

	r.Y, err = ctx.URLParamInt(con_Params_Y)
	if err != nil {
		ctx.WriteString("need y")
		return
	}

	r.W, err = ctx.URLParamInt(con_Params_W)
	if err != nil {
		ctx.WriteString("need w")
		return
	}

	r.H, err = ctx.URLParamInt(con_Params_H)
	if err != nil {
		ctx.WriteString("need h")
		return
	}

	return
}

func commonHandler(name string, ctx iris.Context, err error) {
	ctx.WriteString(name)
	if err != nil {
		ctx.JSON(createErrHSResponse(fmt.Sprintf("%s Fail:%s", name, err)))
		return
	} else {
		ctx.JSON(createSuccessHSResponse(name + "success"))
	}
}
