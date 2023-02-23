download:
	go mod tidy && go mod vendor

build:
	go build -o app

generate_secretkey:
	openssl rand -base64 32