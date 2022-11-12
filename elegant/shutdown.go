package elegant

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// 优雅停机

func Shutdown(cancelFn context.CancelFunc) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown
	cancelFn()
}
