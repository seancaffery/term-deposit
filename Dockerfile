FROM golang:1.24-alpine

WORKDIR /app

COPY . /app

RUN go build

ENTRYPOINT ["/app/term-deposit"]

