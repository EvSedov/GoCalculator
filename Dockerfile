FROM golang:1.21-alpine3.19
RUN apk add --no-cache nodejs npm
WORKDIR /usr/src/app
RUN go install github.com/cosmtrek/air@latest
COPY . .
RUN go mod tidy
