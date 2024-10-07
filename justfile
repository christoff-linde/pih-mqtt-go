set dotenv-load := true

# Define variables

API_DIR := "api"
CONSUMER_DIR := "consumer"
DB_DIR := "db"
API_SQLC_DIR := "db/sqlc/api"
CONSUMER_SQLC_DIR := "db/sqlc/consumer"
MIGRATIONS_DIR := "db/schema"

default:
    just --fmt --unstable
    just --list

# build the api service
build-api:
    cd {{ API_DIR }} && go build -o bin/api cmd/main.go

# build the consumer service
build-consumer:
    cd {{ CONSUMER_DIR }} && go build -o bin/consumer main.go

# run the api service
run-api:
    cd {{ API_DIR }} && go run cmd/main.go

# run the consumer service
run-consumer:
    cd {{ CONSUMER_DIR }} && go run main.go

# clean all build artifacts
clean:
    rm -rf {{ API_DIR }}/bin
    rm -rf {{ CONSUMER_DIR }}/bin

# build all services
build: build-api build-consumer

# run all services
run: run-api run-consumer

# generate sqlc code for api service
sqlcgen-api:
    cd {{ API_SQLC_DIR }} && sqlc generate

# generate sqlc code for consumer service
sqlcgen-consumer:
    cd {{ CONSUMER_SQLC_DIR }} && sqlc generate

# generate sqlc code for all services
sqlcgen: sqlcgen-api sqlcgen-consumer

# apply all db migrations
goose-up:
    goose -dir {{ MIGRATIONS_DIR }} postgres $DB_URL up

# rollback last db migration
goose-down:
    goose -dir {{ MIGRATIONS_DIR }} postgres $DB_URL down

# show the status of db migrations
goose-status:
    goose -dir {{ MIGRATIONS_DIR }} postgres $DB_URL status

# create a new db migration with name
goose-create name:
    goose -dir {{ MIGRATIONS_DIR }} create {{ name }} sql

# vendor all dependencies
vendor:
    go work vendor

# run go mod tidy on all services
tidy:
    cd {{ API_DIR }} && go mod tidy
    cd {{ CONSUMER_DIR }} && go mod tidy
