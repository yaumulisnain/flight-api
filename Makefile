#!bin/sh
define USAGE
Commands:
    run			Run main.go
	build		Build binary
    start       Run built app
	test		Run unit tests
endef

run:
	go mod vendor
	go run main.go

build:
	go mod vendor
	go build -o dist/flight-service

start:
	./dist/flight-service

test:
	go mod vendor
	go test ./app/v1/tests/ -v