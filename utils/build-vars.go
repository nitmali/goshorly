package utils

var (
	GitCommitShort string
	GitBranch      string
	GitBuild       bool
)

func Init_build_vars() {
	if GitCommitShort == "" && GitBranch == "" {
		GitBuild = false
	} else {
		GitBuild = true
	}
}
