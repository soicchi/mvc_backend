download:
	go mod tidy && go mod vendor

build:
	go build -o app ./cmd/chatapp

test:
	go test ./...

fmt:
	go fmt ./...

secret_key:
	openssl rand -base64 32