export I_PACKAGE="git.ucode.space/Phil/goshorly/utils"
export I_GitCommitShort=$(git rev-parse --short HEAD)
export I_GitBranch=$(git rev-parse --abbrev-ref HEAD)

go mod download

case $1 in
    "linux")
    go build -a -installsuffix cgo -ldflags "-X $I_PACKAGE.GitCommitShort=$I_GitCommitShort -X $I_PACKAGE.GitBranch=$I_GitBranch" -o app-linux-amd64 .
    ;;
    "windows")
    go build -a -installsuffix cgo -ldflags "-X $I_PACKAGE.GitCommitShort=$I_GitCommitShort -X $I_PACKAGE.GitBranch=$I_GitBranch" -o app-windows-amd64 .
    ;;
    "darwin")
    go build -a -installsuffix cgo -ldflags "-X $I_PACKAGE.GitCommitShort=$I_GitCommitShort -X $I_PACKAGE.GitBranch=$I_GitBranch" -o app-darwin-amd64 .
    ;;
    *)
    go build -a -installsuffix cgo -ldflags "-X $I_PACKAGE.GitCommitShort=$I_GitCommitShort -X $I_PACKAGE.GitBranch=$I_GitBranch" -o app .
    ;;
esac
