package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/define"
)

func SetupEngine() *gin.Engine {
	if !define.DEV_MODE {
		gin.SetMode(gin.ReleaseMode)
	}
	return gin.Default()
}

func Launched(engine *gin.Engine) {
	GracefulExit(func() {
		if err := engine.Run(define.APP_PORT); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Program start error: %s\n", err)
		}
	})
}
