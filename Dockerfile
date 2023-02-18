FROM golang:1.20-alpine3.16 AS builder
ENV ROOT /app
WORKDIR $ROOT
RUN apk update && apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o ${ROOT}/binary

FROM golang:1.20-alpine3.16 AS dev
ENV ROOT /app \
    TZ Asia/Tokyo
WORKDIR ${ROOT}
RUN apk update && apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["go", "run", "main.go"]

FROM golang:1.20-alpine AS prod
ENV ROOT /app \
    TZ Asia/Tokyo
WORKDIR ${ROOT}
COPY --from=builder /app/binary ${ROOT}
EXPOSE 8080
CMD ["/${ROOT}/binary"]