package api

import (
	"github.com/gin-gonic/gin"
)

func AuthSession(r *gin.Engine) {
	normalAPI := r.Group("/api")
	{
		normalAPI.GET("/auth/session", func(c *gin.Context) {

		})
	}
}
