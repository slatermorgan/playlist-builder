# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

ENV PORT=8080
ENV TABLE_NAME=playlists

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY /cmd ./cmd
COPY /pkg ./pkg
COPY /playlists ./playlists

EXPOSE 8080

CMD PORT=8080 TABLE_NAME=example-playlists go run cmd/server/main.go

