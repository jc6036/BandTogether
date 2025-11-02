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

func RegisterRoutes(engine *gin.Engine) {
	db := getDBConnection()

	engine.LoadHTMLGlob("page/templates/*")

	engine.Static("/styles", "./page/styles")

	// SSR Page Loads
	engine.GET("/home", func(route_context *gin.Context) {
		userdata, err := user_controller.GetUserById(route_context, db)

		if err != nil {
			log.Fatalf("An error has occurred: %s", err.Error())
		}

		route_context.HTML(http.StatusOK, "home.html", userdata)
	})

	// Data routes
	engine.GET("api/search", func(route_context *gin.Context) {
		search_controller.UserSearch(route_context)
	})

	engine.GET("api/events", func(route_context *gin.Context) {
		event_controller.GetUserEvents(route_context)
	})
}

func getDBConnection() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "BandTogether"

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
