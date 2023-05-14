package settings

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func BetaFeatures(c *gin.Context) {
	features := datatypes.BetaFeatures{
		Browsing:        false,
		CodeInterpreter: false,
		Plugins:         false,
	}

	if define.ENABLE_PLUGIN {
		features.Plugins = true
	}

	if define.ENABLE_PLUGIN_BROWSING {
		features.Browsing = true
	}

	if define.ENABLE_PLUGIN_CODE_INTERPRETER {
		features.CodeInterpreter = true
	}

	c.JSON(http.StatusOK, features)
}
