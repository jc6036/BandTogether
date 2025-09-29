package dbaccess

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func getDBConnection() *sql.DB {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = 
	cfg.Passwd =  // TODO: Remove these hardcoded passwords
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
		return "err" // TODO: Actually handle here
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		// TODO: Next point of implementation
	}

	return ""
}
