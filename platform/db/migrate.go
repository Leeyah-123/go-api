package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

// Used for running SQL migrations
func StartMigration() {
	driver, err := postgres.WithInstance(Db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("migration: Running migration ...")
	m, err := migrate.NewWithDatabaseInstance(
		"file://platform/migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		log.Printf("migration: %v\n", err.Error())
	} // or m.Step(2) if you want to explicitly set the number of migrations to run
}
