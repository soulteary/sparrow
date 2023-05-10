package conversation

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	eb "github.com/soulteary/sparrow/components/event-broker"
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

func CreateConversation(brokerPool *eb.BrokersPool) func(c *gin.Context) {
	return func(c *gin.Context) {
		_, data, err := ParseConversationBody(c.Request.Body)
		if err != nil {
			c.Data(http.StatusTeapot, "application/json; charset=utf-8", []byte(fmt.Sprintf("%v", err)))
			return
		}

		if data.ParentMessageID == "" {
			c.Data(http.StatusTeapot, "application/json; charset=utf-8", []byte(fmt.Sprintf("%v", "missing id")))
			return
		}

		// TODO bind user
		userID := c.Request.Header.Get("x-user-id")
		if userID != "" {
			fmt.Println("[user]", userID)
		}

		broker := brokerPool.GetBroker(userID, data.ParentMessageID, data.ConversationID)
		userModel := strings.TrimSpace(strings.ToLower(data.Model))
		userPrompt := strings.TrimSpace(data.Messages[0].Content.Parts[0])

		if define.DEV_MODE {
			fmt.Println("[user request]", "Model:", userModel)
			fmt.Println("[user request]", "Prompt:", userPrompt)
			fmt.Println("[user request]", "Data:", data)
			fmt.Println()
		}

		switch userModel {
		default:
			streamGenerated := StreamBuilder(data, broker, userPrompt)
			if streamGenerated {
				c.Request.Header.Set("x-message-id", data.ParentMessageID)
				broker.Serve(c)
			}
		}
	}
}
