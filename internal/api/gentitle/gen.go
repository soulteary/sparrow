package gentitle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
)

func GetTitle(c *gin.Context) {
	// tips: When the content is empty, do not update the page title
	c.JSON(http.StatusOK, datatypes.GenTitle{Title: ""})
}
