package common

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func CommonHandler(name string, ctx iris.Context, err error) {
	ctx.WriteString(name)
	if err != nil {
		ctx.JSON(CreateErrHSResponse(fmt.Sprintf("%s Fail:%s", name, err)))
		return
	} else {
		ctx.JSON(CreateSuccessHSResponse(name + "success"))
	}
}
