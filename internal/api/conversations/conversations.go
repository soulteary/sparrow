package conversations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/define"
	"github.com/soulteary/sparrow/internal/mock"
)

func GetConversationList(c *gin.Context) {
	// TODO bind user
	userID := c.Request.Header.Get("x-user-id")
	if userID != "" {
		fmt.Println("[user]", userID)
	} else {
		userID = define.DEFAULT_USER_NAME
	}
	c.JSON(http.StatusOK, mock.GetConversationList(userID))
}

func ClearConversationList(c *gin.Context) {
	// TODO bind user
	userID := c.Request.Header.Get("x-user-id")
	if userID != "" {
		fmt.Println("[user]", userID)
	} else {
		userID = define.DEFAULT_USER_NAME
	}
	c.JSON(http.StatusOK, mock.ClearConversationList(userID))
}
