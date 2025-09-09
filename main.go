// TODO: Clean up CSS classes and categorize more effectively for reuse
// TODO: Reduce javascript usage as much as possible on home.html. Limit it to what's absolutely necessary
// TODO: Clean up home.html's actual html for clarity especially with regards to naming
// TODO: Convert home.html into a gin templated page, so you're not just ajaxing all the dynamic content

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
