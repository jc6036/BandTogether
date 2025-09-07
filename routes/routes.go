package routes

import (
	"BandTogether/controllers/search_controller"
	"BandTogether/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user_controller.GetUser(c)
	})

	r.GET("/search", func(c *gin.Context) {
		search_controller.UserSearch(c)
	})

	r.Static("/home", "./page")
}
