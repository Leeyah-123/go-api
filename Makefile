DATABASE_URL = postgres://postgres:admin@localhost:5432/newdb

migrate:
	migrate -databse $(DATABASE_URL) -path platform/migrations up

sqlc:
	sqlc generate