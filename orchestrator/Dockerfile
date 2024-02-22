FROM golang:1.22
WORKDIR /usr/src/app/orchestrator
COPY . .
RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy
