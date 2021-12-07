FROM golang:alpine

RUN apk add --no-cache gcc libgo

WORKDIR /go/src/git.ucode.space/goshorly
COPY . .

RUN go get -d -v ./...
RUN go build -o app .

CMD ["./app"]
