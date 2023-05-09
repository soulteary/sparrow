package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/define"
	"github.com/soulteary/sparrow/internal/server"
)

// var brokerPool = br.NewManager(256)

func main() {
	if !define.DEV_MODE {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// api.AuthSession(router)
	// api.Public(router)
	// api.Backend(router, brokerPool)

	server.GracefulExit(func() {
		if err := router.Run(define.APP_PORT); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Program start error: %s\n", err)
		}
	})
}
