package server

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func GracefulExit(fn func()) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go fn()
	fmt.Println("Sparrow Service has been launched 🚀")

	<-ctx.Done()

	stop()
	fmt.Println("The program is closing, if you want to end it immediately, please press `CTRL+C`")
}
