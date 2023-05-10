package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
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

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				if err, ok := p.(error); ok {
					if errors.Is(err, http.ErrAbortHandler) {
						return
					}
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
