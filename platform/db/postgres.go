package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var Db *sql.DB

func ConnectToDatabase() {
	var err error
	Db, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
}
