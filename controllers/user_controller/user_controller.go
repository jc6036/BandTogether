package user_controller

import (
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) gin.H {
	// userId := c.Param("userId")

	// Get DB Connection
	// SELECT JSONB FROM users WHERE users.id = <?param>
	// Return

	return gin.H{
		"id":     "testid",
		"name":   "Jesse Cabell",
		"avatar": "https://images.unsplash.com/photo-1502685104226-ee32379fefbe?w=256&q=80&auto=format&fit=crop",
	}
}
