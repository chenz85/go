// +build !debug

package log

// 日志，级别：调试。只在测试环境输出
func D(v ...interface{}) {
}

// 日志，级别：调试。只在测试环境输出
func D2(format string, v ...interface{}) {
}
