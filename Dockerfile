FROM golang:alpine as builder

RUN apk add --no-cache gcc libgo

WORKDIR /go/src/git.ucode.space/goshorly
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM scratch as production
WORKDIR /goshorly
copy --from=builder /go/src/git.ucode.space/goshorly/app /goshorly/app
CMD ["./app"]