# Makefile with helpful commands
#

BIN             = reporting
OUTPUT_DIR      = bin

.PHONY: help 
.DEFAULT_GOAL := help

## Build ##
build: clean ## builds the binary
	go build -o $(OUTPUT_DIR)/$(BIN) .

tests: ## Runs the unit tests
	go test -cover -v ./...

docker-build: clean ## builds the docker image
	docker build . -t $(BIN)

docker-run: docker-build ## Start the application
	docker-compose up --remove-orphans

docker-init: ## Initialize the Database, this should be run after docker-run
	# Create the DB if it does not exist
	-docker exec -it reporting_db_1 sh -c  "psql -U postgres -tc \"SELECT 1 FROM pg_database WHERE datname = 'reporting'\" | grep -q 1 | psql -U postgres -c \"CREATE DATABASE reporting\""
	# Load the seed data
	-docker exec -it reporting_db_1 sh -c "psql -h localhost -d reporting -U postgres -f  /var/data/db_init.sql"

clean: ## Remove build artifacts
	$(RM) $(OUTPUT_DIR)/*
	-docker image rm -f $(BIN)

help: ## Display this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_\/-]+:.*?## / {printf "\033[34m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | \
		sort | \
		grep -v '#'
