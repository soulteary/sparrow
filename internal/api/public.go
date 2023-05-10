package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/api/conversation"
	"github.com/soulteary/sparrow/internal/define"
	"github.com/soulteary/sparrow/internal/mock"
)

func Public(r *gin.Engine) {
	publicAPI := r.Group("/public-api")
	{
		publicAPI.GET("/conversation_limit", func(c *gin.Context) {
			if define.ENABLE_MOCK {
				c.JSON(http.StatusOK, mock.GetConversationLimit())
			} else {
				c.JSON(http.StatusOK, conversation.GetConversationLimit())
			}
		})
	}
}
