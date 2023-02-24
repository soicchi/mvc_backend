download:
	go mod tidy && go mod vendor

build:
	go build -o app ./cmd/chatapp

test:
	go test ./...

generate_secretkey:
	openssl rand -base64 32