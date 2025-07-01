SEEDING_PATH=./cmd/database/main.go
SERVER_PATH=./cmd/server/main.go

.PHONY = run seedDBs startServer

seedDBs:
	@echo "Seeding all databases..."
	@GO_ENV=development go run ${SEEDING_PATH} seedTenantDBs

startServer:
	@echo "Starting server..."
	@GO_ENV=development go run ${SERVER_PATH}

run: seedDBs startServer
	@echo "Server ready!"