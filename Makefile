PKGS ?= $(shell go list ./...)
.PHONY: all services clean

run:
	go run main.go

test:
	go test -covermode=atomic -coverprofile coverage.out -v ${PKGS}

migrate_up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate_down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down		
