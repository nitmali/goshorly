FROM golang:latest

COPY . /app

WORKDIR /app

RUN go mod download

CMD ["go", "run", "."]