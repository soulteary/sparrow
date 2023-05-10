package conversation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
	"github.com/soulteary/sparrow/internal/mock"
)

func GetConversationById(c *gin.Context) {
	if !define.ENABLE_HISTORY_LIST {
		c.JSON(http.StatusForbidden, "History needs to be enabled")
	} else {
		id := c.Param("id")
		c.JSON(http.StatusOK, mock.GetConversationById(id))
	}
}

func UpdateConversation(c *gin.Context) {
	c.JSON(http.StatusOK, datatypes.UpdateConversationResponse{Success: true})
}
