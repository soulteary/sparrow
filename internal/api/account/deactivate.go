package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/mock"
)

func Deactivate(c *gin.Context) {
	c.JSON(http.StatusOK, mock.Deactivate())
}
