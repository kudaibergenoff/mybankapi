# Initial stage: download modules
FROM golang:1.23-alpine as builder

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go mod download

# Intermediate stage: Build the binary
FROM golang:1.23-alpine as runner

COPY --from=builder /app /app

RUN apk add --no-cache git && \
    go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app
ENV config=docker

EXPOSE 8080

ENTRYPOINT ["CompileDaemon", "--build=go build -o mybankapi cmd/main.go", "--command=./mybankapi"]