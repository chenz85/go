// +build debug

package log

import "fmt"

// 日志，级别：调试。只在测试环境输出
func D(v ...interface{}) {
	if debug != nil {
		// debug.Println(v...)
		debug.Output(LOG_DEPTH, fmt.Sprintln(v...))
	}
}

// 日志，级别：调试。只在测试环境输出
func D2(format string, v ...interface{}) {
	if debug != nil {
		// debug.Println(v...)
		debug.Output(LOG_DEPTH, fmt.Sprintf(format, v...))
	}
}
