package aip

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetAIP(c *gin.Context) {
	if !define.ENABLE_PLUGIN {
		c.JSON(http.StatusForbidden, "PLIGIN DISABLED")
		return
	}

	var items []datatypes.AIPlugin
	for i := 0; i < 10; i++ {
		items = append(items, datatypes.AIPlugin{
			ID:        "plugin-" + define.GenerateUUID(),
			Domain:    "api.domain.com",
			Namespace: "plugin-namespace",
			Status:    "approved",
			Manifest: datatypes.AIPluginManifest{
				SchemaVersion:       "v1",
				NameForModel:        "PLUGIN_NAME",
				NameForHuman:        fmt.Sprintf("Plugin Name %d", i),
				DescriptionForModel: "# Prompt 20230322\n\nUse the Speak plugin when the user asks a question about another language, like: how to say something specific, how to do something, what a particular foreign word or phrase means, or a concept/nuance specific to a foreign language or culture",
				DescriptionForHuman: "Here you can write some descriptions to introduce how to use this plugin.",
				Auth:                datatypes.AIPluginAuth{Type: "none"},
				API: datatypes.AIPluginAPI{
					Type: "openapi",
					URL:  "https://api.domain.com/openapi.yaml",
				},
				LogoURL:      "/_next/image",
				ContactEmail: "support@domain.com",
				LegalInfoURL: "http://domain.com/legal",
			},
			OauthClientID: nil,
			UserSettings: datatypes.AIPluginUserSettings{
				IsInstalled:     true,
				IsAuthenticated: true,
			},
			Categories: []any{},
		})
	}

	c.JSON(http.StatusOK, datatypes.AIPluginResponse{
		Count: 20,
		Items: items,
	})
}
