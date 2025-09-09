// TODO: Get the hardcoded dummy data out of the home page
// TODO: Sub-task since I'm breaking for the night - make an AJAX call to the events API to grab the data, it's there and ready

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
