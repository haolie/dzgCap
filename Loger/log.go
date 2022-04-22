package Loger

import (
	"fmt"
	"time"

	"dzgCap/tools"
)

func LogInfo(info string) {
	fmt.Printf("%s:%s\n", tools.ToTimeStr(time.Now()), info)
}
