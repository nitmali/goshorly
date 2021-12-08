export I_PACKAGE="git.ucode.space/Phil/goshorly/utils"
export I_GitCommitShort=$(git rev-parse --short HEAD)
export I_GitBranch=$(git rev-parse --abbrev-ref HEAD)

go mod download

case $1 in
    "linux")
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X $I_PACKAGE.GitCommitShort=$I_GitCommitShort -X $I_PACKAGE.GitBranch=$I_GitBranch" -o app-linux-amd64 .
    ;;
    "windows")
    CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -ldflags "-X $I_PACKAGE.GitCommitShort=$I_GitCommitShort -X $I_PACKAGE.GitBranch=$I_GitBranch" -o app-windows-amd64 .
    ;;
    "darwin")
    CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -ldflags "-X $I_PACKAGE.GitCommitShort=$I_GitCommitShort -X $I_PACKAGE.GitBranch=$I_GitBranch" -o app-darwin-amd64 .
    ;;
    *)
    CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags "-X $I_PACKAGE.GitCommitShort=$I_GitCommitShort -X $I_PACKAGE.GitBranch=$I_GitBranch" -o app .
    ;;
esac
