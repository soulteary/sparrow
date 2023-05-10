package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/mock"
)

func DataExport(c *gin.Context) {
	c.PureJSON(http.StatusOK, mock.DataExport())
}
