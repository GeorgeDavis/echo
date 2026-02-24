FROM golang:1.26-trixie

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy