.PHONY: build build-docker test run run-docker migrate migrate_test

migrate:
	migrate -path migrations -database "postgres://poolcrm:poolcrm@localhost:5433/poolcrm?sslmode=disable" up
migrate-test:
	migrate -path migrations -database "postgres://poolcrm:poolcrm@localhost:5433/poolcrm_test?sslmode=disable" up
build:
	GOOS=linux GOARCH=amd64 go build -v ./cmd/bot
build-docker:
	docker build -t bot .
	docker run --rm -v $(CURDIR):/local bot cp /bin/bot /local/
test:
	go test -v -race ./...
test-docker:
	docker run --rm -w="/project" -v C:\Users\mikea\Documents\code\aggy-rest:/project golang:1.13 go test -v -race ./...
run: build
	./bot
run-docker: build-docker
	docker-compose up -d
generate-doc:
	docker run --rm -v $(CURDIR):/local openapitools/openapi-generator-cli generate -i /local/api/openapi.yaml -g html2 -o /local/docs
generate:
	docker run --rm -v $(CURDIR):/local openapitools/openapi-generator-cli generate -i /local/api/openapi.yaml -g go-server -o /local/openapi

.DEFAULT_GOAL := run