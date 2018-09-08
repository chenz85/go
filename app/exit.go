package app

import (
	"os"

	"github.com/smartswarm/go/event"
	"github.com/smartswarm/go/log"
)

var (
	is_waiting_exit bool
)

func IsWaitingExit() bool {
	return is_waiting_exit
}

func Exit(code int) {
	is_waiting_exit = true

	// 事件，通知其他功能执行退出前的清理
	// os_exit 事件的响应者都应该阻塞的完成退出前的清理操作
	event.Emit("on_exit")

	log.I("[app]", "all clear, exit.")
	// 退出程序
	os.Exit(code)
}
