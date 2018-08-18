package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/czsilence/go/event"
	"github.com/czsilence/go/log"
)

///////////////////////////////////////////////////////////////////
func HandleInterrupt() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-interrupt:
			log.I("[app]", "got interrupt, clear and exiting...")
			// 事件，通知其他功能执行退出前的清理
			event.Emit("on_exit")

			log.I("[app]", "all clear, exit.")
			// 退出程序
			os.Exit(0)
		}
	}
}
