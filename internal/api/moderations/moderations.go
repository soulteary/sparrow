package moderations

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetModerations(c *gin.Context) {
	c.JSON(http.StatusOK, datatypes.Moderations{
		Flagged:      false, // maybe
		Blocked:      false, // must be
		ModerationID: "modr-" + define.GenerateRandomString(29),
	})
}
