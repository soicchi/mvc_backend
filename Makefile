.PHONY: up download build test fmt secret_key

up:
	docker compose up -d && sleep 3 && docker compose up

stop:
	docker compose stop

download:
	docker compose run --rm api go mod tidy && go mod vendor

build:
	docker compose run --rm api go build -o app ./cmd/chatapp

test:
	docker compose run --rm api go test -cover ./...

fmt:
	docker compose run --rm api go fmt ./...

secret_key:
	docker compose run --rm api openssl rand -base64 32