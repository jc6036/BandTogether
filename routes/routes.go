package routes

import (
	"BandTogether/controllers/event_controller"
	"BandTogether/controllers/search_controller"
	"BandTogether/controllers/user_controller"

	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func RegisterRoutes(r *gin.Engine) {
	db := getDBConnection()

	r.LoadHTMLGlob("page/templates/*")

	r.Static("/styles", "./page/styles")

	// SSR Page Loads
	r.GET("/home", func(c *gin.Context) {
		user, err := user_controller.GetUserById(c, db)

		if err != nil {
			log.Fatalf("An error has occurred: %s", err.Error())
		}

		c.HTML(http.StatusOK, "home.html", user)
	})

	// Data routes
	r.GET("api/search", func(c *gin.Context) {
		search_controller.UserSearch(c)
	})

	r.GET("api/events", func(c *gin.Context) {
		event_controller.GetUserEvents(c)
	})
}

func getDBConnection() *sql.DB {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "BandTogether"

	// Get a database handle.
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return db
}
