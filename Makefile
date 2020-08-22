#!bin/sh
define USAGE
Commands:
    run			Run main.go
	build		Build binary
    start       Run built app
endef

run:
	go run main.go

build:
	go build -o dist/flight-service

start:
	./dist/flight-service