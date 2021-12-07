FROM golang:latest

WORKDIR /go/src/git.ucode.space/goshortly

COPY . .

RUN go get -d -v ./...

CMD ["go", "run", "main.go"]