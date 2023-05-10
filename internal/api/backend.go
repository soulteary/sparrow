package api

import (
	"github.com/gin-gonic/gin"
	eb "github.com/soulteary/sparrow/components/event-broker"
	"github.com/soulteary/sparrow/internal/api/account"
	"github.com/soulteary/sparrow/internal/api/models"
)

func Backend(r *gin.Engine, brokerPool *eb.BrokersPool) {
	backendAPI := r.Group("/backend-api")
	{
		// account
		backendAPI.GET("/accounts/check", account.AccountCheck)
		backendAPI.POST("/accounts/data_export", account.DataExport)
		backendAPI.POST("/accounts/deactivate", account.Deactivate)
		// // share
		// backendAPI.POST("/share/create", share.Create)
		// backendAPI.PATCH("/share/", share.Create)
		// model
		backendAPI.GET("/models", models.GetModels)
		// // conversation
		// backendAPI.POST("/moderations", moderations.GetModerations)
		// backendAPI.POST("/conversation/gen_title/", gentitle.GetTitle)
		// backendAPI.POST("/conversation/gen_title/:uuid", gentitle.GetTitle)
		// backendAPI.GET("/conversation/:id", conversation.GetConversationById)
		// backendAPI.PATCH("/conversation/:id", conversation.UpdateConversation)
		// backendAPI.POST("/conversation", conversation.CreateConversation(brokers))
		// backendAPI.GET("/conversations", conversations.GetConversationList)
		// backendAPI.PATCH("/conversations", conversations.GetConversationList)
		// // The interface submission method used in the historical version is POST, which is temporarily reserved
		// backendAPI.POST("/conversations", conversations.GetConversationList)
		// // The page configuration interface prepared for the front end
		// backendAPI.GET("/pageconfig", pageconfig.GetConfig)
		// // AI Plugin
		// backendAPI.GET("/aip/p", aip.GetAIP)
		// // Misc
		// backendAPI.GET("/opengraph/tags", opengraph.OpengraphTag)
	}
}
