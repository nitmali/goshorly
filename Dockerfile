FROM golang:alpine as builder

RUN apk add --no-cache git make build-base

ARG CI_COMMIT_BRANCH
ARG CI_COMMIT_SHORT_SHA
ARG CI_COMMIT_TAG

ENV CI_COMMIT_BRANCH=$CI_COMMIT_BRANCH
ENV CI_COMMIT_SHORT_SHA=$CI_COMMIT_SHORT_SHA
ENV CI_COMMIT_TAG=$CI_COMMIT_TAG

ENV I_PACKAGE="git.ucode.space/Phil/goshorly/utils"
ENV CGO_ENABLED=0

WORKDIR /go/src/git.ucode.space/goshorly
COPY . .

RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -ldflags=" -X $I_PACKAGE.CI_COMMIT_SHORT_SHA=$CI_COMMIT_SHORT_SHA -X $I_PACKAGE.CI_COMMIT_BRANCH=$CI_COMMIT_BRANCH -X $I_PACKAGE.CI_COMMIT_TAG=$CI_COMMIT_TAG" -o app .

FROM scratch as production
WORKDIR /
COPY --from=builder /go/src/git.ucode.space/goshorly/app /app
ENTRYPOINT [ "/app" ]
