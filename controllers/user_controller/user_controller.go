package user_controller

import (
	"github.com/gin-gonic/gin"

	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func GetUserById(c *gin.Context) gin.H {
	userId := c.Query("userId")

	qstr := "SELECT userId, json FROM btusers WHERE userId = " + userId + ";"

	db := getDBConnection()
	data := performRead(db, qstr)

	var user User
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		panic(err)
	}

	return gin.H{
		"id":     user.ID,
		"name":   user.Name,
		"avatar": user.Avatar,
	}
}

func getDBConnection() *sql.DB {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = ""
	cfg.Passwd = ""
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

func performRead(db *sql.DB, statement string) string {
	rows, err := db.Query(statement)
	if err != nil {
		return err.Error()
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	var retstring string
	for rows.Next() {
		var id int
		var jret string

		err := rows.Scan(&id, &jret)
		if err != nil {
			log.Fatal(err)
		}

		retstring = jret
	}

	return retstring
}
