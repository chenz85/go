package app

import (
	"os"
	"os/signal"
	"syscall"

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
			Exit(0)
		}
	}
}
