FROM golang:alpine as builder

RUN apk add --no-cache gcc libgo git

WORKDIR /go/src/git.ucode.space/goshorly
COPY . .

RUN go get -d -v ./...
RUN chmod +x build-ci.sh
RUN ./build-ci.sh

FROM scratch as production
WORKDIR /goshorly
COPY --from=builder /go/src/git.ucode.space/goshorly/app /goshorly/app
CMD ["./app"]