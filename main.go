// TODO: Reduce javascript usage as much as possible on home.html. Limit it to what's absolutely necessary

// Every program starts here
package main

import (
	"BandTogether/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
