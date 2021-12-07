FROM golang:latest

COPY . /app

WOKRDIR /app

RUN go mod download

CMD["go", "run", "."]