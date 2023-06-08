package conversation

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	eb "github.com/soulteary/sparrow/components/event-broker"
	lcs "github.com/soulteary/sparrow/components/local-conversation-storage"
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

		// TODO bind user
		userID := c.Request.Header.Get("x-user-id")
		if userID != "" {
			fmt.Println("[user]", userID)
		} else {
			userID = define.DEFAULT_USER_NAME
		}

		c.JSON(http.StatusOK, mock.GetConversationById(userID, id))
	}
}

func UpdateConversation(c *gin.Context) {
	type RequestBody struct {
		IsVisible bool `json:"is_visible"`
	}
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !requestBody.IsVisible {
		// TODO bind user
		userID := c.Request.Header.Get("x-user-id")
		if userID != "" {
			fmt.Println("[user]", userID)
		} else {
			userID = define.DEFAULT_USER_NAME
		}

		id := c.Param("id")
		mock.ClearConversationByID(userID, id)
	}

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

		if data.ConversationID == "" {
			if data.Action == "next" && len(data.Messages) == 1 && strings.EqualFold(data.Messages[0].Author.Role, "user") {
				data.ConversationID = data.ParentMessageID
			} else {
				fmt.Println("unhandle missing conversation id")
			}
		}
		fmt.Println("[conversation]", data.ParentMessageID, data.ConversationID)
		c.Request.Header.Set("x-conversation-id", data.ConversationID)
		c.Request.Header.Set("x-parent-message-id", data.ParentMessageID)

		// TODO bind user
		userID := c.Request.Header.Get("x-user-id")
		if userID != "" {
			fmt.Println("[user]", userID)
		} else {
			userID = define.DEFAULT_USER_NAME
		}

		broker := brokerPool.GetBroker(userID, data.ConversationID, data.ParentMessageID)
		userModel := strings.TrimSpace(strings.ToLower(data.Model))
		userPrompt := strings.TrimSpace(data.Messages[0].Content.Parts[0])
		parentMessageID := data.ParentMessageID
		messageID := define.GenerateUUID()
		nextMessageID := data.Messages[0].ID

		if define.DEV_MODE {
			fmt.Println("[user request]", "Model:", userModel)
			fmt.Println("[user request]", "Prompt:", userPrompt)
			fmt.Println("[user request]", "Data:", data)
			fmt.Println()
		}

		if lcs.IsRootMessage(userID, data.ConversationID, parentMessageID) {
			lcs.SetRootMessage(userID, parentMessageID, messageID, userPrompt)
		} else {
			lcs.SetMessage(userID, parentMessageID, messageID, userPrompt, true)
		}
		lcs.SetMessage(userID, messageID, nextMessageID, userPrompt, true)

		messageChan := make(eb.EventChan)

		switch userModel {
		case datatypes.MODEL_MIDJOURNEY.Slug:
			message := midjourney.BuildMessage(data.ConversationID, parentMessageID, messageID, nextMessageID, userPrompt)
			midjourney.PostMessage(midjourney.GetConn(), message)
			broker.Serve(c, messageChan)
			return
		case datatypes.MODEL_FLAGSTUDIO.Slug:
			streamGenerated := sr.StreamBuilder(userID, data.ConversationID, parentMessageID, messageID, nextMessageID, userModel, broker, userPrompt, sr.MSG_STATUS_AUTO_MODE)
			if streamGenerated {
				broker.Serve(c, messageChan)
			}
			return
		case datatypes.MODEL_CLAUDE.Slug:
			message := []byte(fmt.Sprintf("%s\n%s", parentMessageID, userPrompt))
			claude.PostMessage(claude.GetConn(), message)
			broker.Serve(c, messageChan)
			return
		default:
			streamGenerated := sr.StreamBuilder(userID, data.ConversationID, parentMessageID, messageID, nextMessageID, userModel, broker, userPrompt, sr.MSG_STATUS_AUTO_MODE)
			if streamGenerated {
				broker.Serve(c, messageChan)
			}
			return
		}
	}
}
