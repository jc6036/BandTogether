package user_controller

import (
	"github.com/gin-gonic/gin"

	"database/sql"
	"encoding/json"
	"log"
)

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func GetUserById(c *gin.Context, db *sql.DB) gin.H {
	userId := c.Query("userId")
	// TODO: Sanitize against SQL injection
	qstr := "SELECT userId, json FROM btusers WHERE userId = " + userId + ";"

	data := queryDB(db, qstr)

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

func queryDB(db *sql.DB, statement string) string {
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
