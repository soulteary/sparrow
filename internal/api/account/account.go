package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/mock"
)

func AccountCheck(c *gin.Context) {
	c.JSON(http.StatusOK, mock.AccountCheck())
}
