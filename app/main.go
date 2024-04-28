package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ganganbiz1/go-echo-gorm-template/app/logger"
	"github.com/ganganbiz1/go-echo-gorm-template/app/wire"
)

func main() {

	logger.Setting()

	di, c, err := wire.DI()
	if err != nil {
		return
	}
	defer c()

	di.Router.Apply(di.Server.Echo)

	go func() {
		err = di.Server.Start()
		if err != nil {
			logger.Error("Server Start failed", err)
			return
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := di.Server.Echo.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown failed", err)
	}
}
