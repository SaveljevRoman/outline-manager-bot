package main

import (
	"context"
	"os"
	"os/signal"
	"outline-manager-bot/config"
	"outline-manager-bot/internal/app"
	"syscall"
)

func main() {
	cfg := config.LoadConfig()

	ctx, cancelFunc := context.WithCancel(context.Background())

	appInstance := app.NewApp(ctx, cfg)
	appInstance.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	// Ждем сигнала завершения приложения
	<-signalChan
	cancelFunc()
}
