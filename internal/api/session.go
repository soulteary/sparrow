package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/define"
	"github.com/soulteary/sparrow/internal/mock"
)

func AuthSession(r *gin.Engine) {
	group := r.Group("/api")
	{
		if define.ENABLE_MOCK {
			group.GET("/auth/session", func(c *gin.Context) {
				c.JSON(http.StatusOK, mock.AuthSession())
			})
		} else {
			// TODO custom API
		}
	}
}
