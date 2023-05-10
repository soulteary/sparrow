package conversations

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/mock"
)

func GetConversationList(c *gin.Context) {
	c.JSON(http.StatusOK, mock.GetConversationList())
}
