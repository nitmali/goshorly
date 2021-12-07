FROM golang:alpine

WORKDIR /go/src/git.ucode.space/goshortly

COPY . .

RUN apk add gcc libgo

RUN go get -d -v ./...

CMD ["go", "run", "main.go"]