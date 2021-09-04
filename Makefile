include linting.mk

local_run:
	go run main.go

deps:
	go mod tidy
	go mod vendor

unit-test:
	go test -race -count=1 -v -cover ./...