package pageconfig

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetConfig(c *gin.Context) {
	var pageInfo datatypes.PageInfo

	pageInfo.Props.PageProps.User.ID = define.MOCK_USER_ID
	pageInfo.Props.PageProps.User.Name = define.MOCK_USER_NAME
	pageInfo.Props.PageProps.User.Email = define.MOCK_USER_EMAIL
	pageInfo.Props.PageProps.User.Image = define.MOCK_USER_IMAGE
	pageInfo.Props.PageProps.User.Picture = define.MOCK_USER_IMAGE
	pageInfo.Props.PageProps.User.Mfa = false
	pageInfo.Props.PageProps.User.Idp = define.MOCK_USER_IDP
	pageInfo.Props.PageProps.User.Iat = define.MOCK_USER_IAT
	pageInfo.Props.PageProps.User.IntercomHash = define.MOCK_INTERCOM_HASH
	pageInfo.Props.PageProps.User.Groups = []string{}
	pageInfo.Props.PageProps.UserCountry = define.MOCK_USER_REGION
	pageInfo.Props.PageProps.GeoOk = true
	pageInfo.Props.PageProps.IsUserInCanPayGroup = true
	pageInfo.Props.PageProps.SentryTraceData = define.MOCK_SENTRY_TRACE_DATA
	pageInfo.Props.PageProps.SentryBaggage = define.MOCK_SENTRY_TRACE_ID

	pageInfo.Props.NSsp = true
	pageInfo.IsFallback = false
	pageInfo.Gssp = true
	pageInfo.ScriptLoader = []string{}

	if define.ENABLE_PLUGIN {
		pageInfo.Props.PageProps.User.Groups = append(pageInfo.Props.PageProps.User.Groups, "chatgpt-plugin-partners")
		pageInfo.Query.Model = "text-davinci-002-plugins"
	}

	pageInfo.BuildID = c.Query("BuildID")
	chatId := c.Query("ChatID")
	model := c.Query("Model")
	if model != "" {
		pageInfo.Query.Model = model
	}

	if chatId != "" {
		pageInfo.Page = "/c/[chatId]"
		pageInfo.Query.ChatID = chatId
	} else {
		pageInfo.Page = "/"
	}

	c.JSON(http.StatusOK, &pageInfo)
}
