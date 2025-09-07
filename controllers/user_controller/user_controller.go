package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"user": name,
	})
}
