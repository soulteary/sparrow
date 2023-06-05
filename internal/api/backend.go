package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	eb "github.com/soulteary/sparrow/components/event-broker"
	"github.com/soulteary/sparrow/internal/api/account"
	"github.com/soulteary/sparrow/internal/api/aip"
	"github.com/soulteary/sparrow/internal/api/conversation"
	"github.com/soulteary/sparrow/internal/api/conversations"
	"github.com/soulteary/sparrow/internal/api/gentitle"
	"github.com/soulteary/sparrow/internal/api/models"
	"github.com/soulteary/sparrow/internal/api/moderations"
	"github.com/soulteary/sparrow/internal/api/opengraph"
	"github.com/soulteary/sparrow/internal/api/pageconfig"
	"github.com/soulteary/sparrow/internal/api/settings"
	"github.com/soulteary/sparrow/internal/mock"
)

func Backend(r *gin.Engine, brokerPool *eb.BrokersPool) {
	backendAPI := r.Group("/backend-api")
	{
		// account
		backendAPI.GET("/accounts/check", account.AccountCheck)
		backendAPI.GET("/accounts/check/v4-2023-04-27", account.AccountTempCheck)
		backendAPI.POST("/accounts/data_export", account.DataExport)
		backendAPI.POST("/accounts/deactivate", account.Deactivate)
		backendAPI.GET("/settings/beta_features", settings.BetaFeatures)
		backendAPI.POST("/settings/beta_features", settings.BetaFeatures)
		backendAPI.GET("/user_system_messages", func(c *gin.Context) {
			c.JSON(http.StatusNotFound, mock.GetMessageSystem2())
		})

		// share
		// backendAPI.POST("/share/create", share.Create)
		// backendAPI.PATCH("/share/", share.Create)
		// model
		backendAPI.GET("/models", models.GetModels)
		// conversation
		backendAPI.POST("/moderations", moderations.GetModerations)
		backendAPI.POST("/conversation/gen_title/", gentitle.GetTitle)
		backendAPI.POST("/conversation/gen_title/:uuid", gentitle.GetTitle)
		backendAPI.GET("/conversation/:id", conversation.GetConversationById)
		backendAPI.PATCH("/conversation/:id", conversation.UpdateConversation)
		backendAPI.POST("/conversation", conversation.CreateConversation(brokerPool))
		backendAPI.GET("/conversations", conversations.GetConversationList)
		backendAPI.PATCH("/conversations", conversations.ClearConversationList)
		// The interface submission method used in the historical version is POST, which is temporarily reserved
		backendAPI.POST("/conversations", conversations.GetConversationList)
		// The page configuration interface prepared for the front end
		backendAPI.GET("/pageconfig", pageconfig.GetConfig)
		// AI Plugin
		backendAPI.GET("/aip/p", aip.GetAIP)
		// Misc
		backendAPI.GET("/opengraph/tags", opengraph.OpengraphTag)
	}
}
