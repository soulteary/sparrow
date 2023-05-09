package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/api/conversation"
)

func Public(r *gin.Engine) {
	publicAPI := r.Group("/public-api")
	{
		publicAPI.GET("/conversation_limit", func(c *gin.Context) {
			c.JSON(http.StatusOK, conversation.GetConversationLimit())
		})
	}
}
