package user_controller

import (
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) gin.H {
	// name := c.Param("name")

	return gin.H {
		"id":     "testid",
		"name":   "Jesse Cabell",
		"avatar": "https://images.unsplash.com/photo-1502685104226-ee32379fefbe?w=256&q=80&auto=format&fit=crop",
	}
}