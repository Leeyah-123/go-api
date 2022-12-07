package main

import (
	"fiber/app"
	"fiber/platform/db"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db.ConnectToDatabase()
	app.StartApp()
}
