DATABASE_PATH=./cmd/database/main.go
SERVER_PATH=./cmd/server/main.go

.PHONY = run migrateDBs dropDBs startServer

migrateDBs:
	@echo "Migrating all databases..."
	@GO_ENV=development go run ${DATABASE_PATH} migrateDBs

dropDBs:
	@echo "Dropping all databases..."
	@GO_ENV=development go run ${DATABASE_PATH} dropDBs

startServer:
	@echo "Starting server..."
	@GO_ENV=development go run ${SERVER_PATH}

run: startServer
	@echo "Server ready!"