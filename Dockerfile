FROM golang:alpine as builder

RUN apk add --no-cache gcc libgo

WORKDIR /go/src/git.ucode.space/goshorly
COPY . .

RUN go get -d -v ./...
RUN ./build-ci.sh

FROM gruebel/upx:latest as upx
COPY --from=builder /go/src/git.ucode.space/goshorly/app /app.org
RUN upx --best --lzma -o /app /app.org

FROM scratch as production
WORKDIR /goshorly
copy --from=upx /app /goshorly/app
CMD ["./app"]