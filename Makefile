include .env

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)
PG_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable


BINARY=gravitum

.PHONY: all build run test docker-build docker-run docker-stop swagger

all: build

build:
	go build -o $(GOBIN)/$(BINARY) ./cmd/app/main.go

run:
	go run ./cmd/app/main.go

test:
	go test -v ./...

docker-build:
	docker-compose build

docker-run:
	docker-compose up -d

docker-stop:
	docker-compose down

swagger:
	swag init -g internal/app/app.go -o docs
