# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -x -v -o /app ./cmd/main.go

CMD ["/app"]
