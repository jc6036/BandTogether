package routes

import (
	"BandTogether/controllers/event_controller"
	"BandTogether/controllers/search_controller"
	"BandTogether/controllers/user_controller"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("page/templates/*")

	r.Static("/styles", "./page/styles")

	// SSR Page Loads
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", user_controller.GetUser(c))
	})

	// Data routes
	r.GET("api/search", func(c *gin.Context) {
		search_controller.UserSearch(c)
	})

	r.GET("api/events", func(c *gin.Context) {
		event_controller.GetUserEvents(c)
	})
}
