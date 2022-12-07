package db

import "fiber/postgres"

func Query() *postgres.Queries {
	query := postgres.New(Db)
	return query
}
