package user_controller

import (
	"github.com/gin-gonic/gin"

	"database/sql"
	"encoding/json"
)

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func GetUserById(c *gin.Context, db *sql.DB) (gin.H, error) {
	userId := c.Query("userId")
	// TODO: Sanitize against SQL injection
	qstr := "SELECT userId, json FROM btusers WHERE userId = " + userId + ";"

	data, err := queryDB(db, qstr)

	if err != nil {
		return gin.H{}, err
	}

	var user User
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return gin.H{}, err
	}

	return gin.H{
		"id":     user.ID,
		"name":   user.Name,
		"avatar": user.Avatar,
	}, nil
}

func queryDB(db *sql.DB, statement string) (string, error) {
	// Straightforward - Run the provided statement on the provided DB connection, return results
	rows, err := db.Query(statement)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	var retstring string = ""
	for rows.Next() {
		var id int
		var jret string

		err := rows.Scan(&id, &jret)
		if err != nil {
			return "", err
		}

		retstring = jret
	}

	if retstring == "" {
		return "No rows found!", nil
	}

	err = rows.Err()

	if err != nil {
		return "", err
	}

	return retstring, nil
}
