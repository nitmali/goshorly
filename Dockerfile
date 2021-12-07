FROM golang:alpine as builder

RUN apk add --no-cache gcc libgo

WORKDIR /go/src/git.ucode.space/goshorly
COPY . .

RUN go get -d -v ./...
RUN go build -o app .

FROM scratch as production

WORKDIR /goshorly
COPY --from=builder /go/src/git.ucode.space/goshorly/app .
CMD ["./app"]
