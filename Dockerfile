FROM golang:alpine as builder

RUN apk add --no-cache gcc libgo git

WORKDIR /go/src/git.ucode.space/goshorly
COPY . .

RUN go get -d -v ./...

RUN export I_PACKAGE="git.ucode.space/Phil/goshorly/utils" && \
    export I_GitCommitShort=$(git rev-parse --short HEAD) && \
    export I_GitBranch=$(git rev-parse --abbrev-ref HEAD) && \
    go build -a -installsuffix cgo -ldflags "-X $I_PACKAGE.GitCommitShort=$I_GitCommitShort -X $I_PACKAGE.GitBranch=$I_GitBranch" -o app .

FROM scratch as production
WORKDIR /goshorly
COPY --from=builder /go/src/git.ucode.space/goshorly/app /goshorly/app
CMD ["./app"]