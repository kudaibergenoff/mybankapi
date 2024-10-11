# Initial stage: download modules
FROM golang:1.23-alpine as builder

WORKDIR /app

COPY .. /app

RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go mod download

ENTRYPOINT CompileDaemon --build="go build -o mybankapi cmd/main.go" --command=./mybankapi