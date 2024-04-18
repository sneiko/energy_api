
create_migration:
	migrate create -ext sql -dir ./internal/infrastructure/mariadb/migrations -seq ${name}