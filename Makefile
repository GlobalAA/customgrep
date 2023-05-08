.PHONY: build

build:
	go build -o grep cmd/app/main.go

.DEFAULT_GOAL := build
