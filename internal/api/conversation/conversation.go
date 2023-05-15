package conversation

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	eb "github.com/soulteary/sparrow/components/event-broker"
	sr "github.com/soulteary/sparrow/components/stream-responser"
	claude "github.com/soulteary/sparrow/connectors/claude"
	midjourney "github.com/soulteary/sparrow/connectors/mid-journey"
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
		_, data, err := sr.ParseConversationBody(c.Request.Body)
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

		messageChan := make(eb.EventChan)
		c.Request.Header.Set("x-message-id", data.ParentMessageID)

		switch userModel {
		case datatypes.MODEL_MIDJOURNEY.Slug:
			message := []byte(fmt.Sprintf("%s\n%s", data.ParentMessageID, userPrompt))
			midjourney.PostMessage(midjourney.GetConn(), message)
			broker.Serve(c, messageChan)
			return
		case datatypes.MODEL_FLAGSTUDIO.Slug:
			streamGenerated := sr.StreamBuilder(data.ParentMessageID, data.ConversationID, userModel, broker, userPrompt, sr.MSG_STATUS_AUTO_MODE)
			if streamGenerated {
				broker.Serve(c, messageChan)
			}
			return
		case datatypes.MODEL_CLAUDE.Slug:
			message := []byte(fmt.Sprintf("%s\n%s", data.ParentMessageID, userPrompt))
			midjourney.PostMessage(claude.GetConn(), message)
			broker.Serve(c, messageChan)
			return
		default:
			streamGenerated := sr.StreamBuilder(data.ParentMessageID, data.ConversationID, userModel, broker, userPrompt, sr.MSG_STATUS_AUTO_MODE)
			if streamGenerated {
				broker.Serve(c, messageChan)
			}
			return
		}
	}
}
