// Every program starts here
package main

import (
	"BandTogether/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	routes.RegisterRoutes(engine)

	engine.Run(":8080")
}
