package search_controller

import (
	"github.com/gin-gonic/gin"
)

func UserSearch(c *gin.Context) {
	id := c.Query("id")
	c.JSON(200, gin.H{
		"search_id": id,
	})
}
