package funcutil

import (
	"fmt"
	"runtime/debug"
)

// Recover 打印panic信息，使用时要放在函数最开始位置
func Recover() {
	if err := recover(); err != nil {
		fmt.Println(err)
		debug.PrintStack()
	}
}
