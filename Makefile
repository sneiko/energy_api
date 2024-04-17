
create_migration:
	migrate create -ext sql -dir ./internal/infrastructure/pg/migrations -seq ${name}