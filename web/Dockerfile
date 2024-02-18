FROM golang:1.21-alpine3.19
RUN apk update && apk add --no-cache bash nodejs npm
WORKDIR /usr/src/app/web/
COPY . .
RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy
