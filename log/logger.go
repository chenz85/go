package log

import (
	"fmt"
	"log"
	"os"
)

const (
	LOG_DEPTH int = 2
)

var (
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
	debug *log.Logger
)

// 初始化日志
func init() {
	outFile := os.Stdout
	flag := log.LstdFlags | log.Lmicroseconds

	info = log.New(outFile, "[INFO] ", flag)
	warn = log.New(outFile, "[WARN] ", flag)
	err = log.New(outFile, "[ERROR] ", flag|log.Lshortfile)
	debug = log.New(outFile, "[DEBUG] ", flag|log.Lshortfile)
}

// 日志，级别：信息
func I(v ...interface{}) {
	if info != nil {
		info.Println(v...)
	}
}

// 日志，级别：信息
func I2(format string, v ...interface{}) {
	if info != nil {
		info.Printf(format, v...)
	}
}

// 日志，级别：警告
func W(v ...interface{}) {
	if warn != nil {
		warn.Println(v...)
	}
}

// 日志，级别：警告
func W2(format string, v ...interface{}) {
	if warn != nil {
		warn.Printf(format, v...)
	}
}

// 日志，级别：错误。输出并退出程序
func E(v ...interface{}) {
	if err != nil {
		// err.Fatalln(v...)
		err.Output(LOG_DEPTH, fmt.Sprintln(v...))
		panic(err)
		//os.Exit(1)
	}
}

// 日志，级别：错误。输出并退出程序
func E2(format string, v ...interface{}) {
	if err != nil {
		// err.Fatalln(v...)
		err.Output(LOG_DEPTH, fmt.Sprintf(format, v...))
		panic(err)
		//os.Exit(1)
	}
}
