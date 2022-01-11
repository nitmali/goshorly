package utils

var (
	CI_COMMIT_SHORT_SHA string
	CI_COMMIT_BRANCH    string
	CI_COMMIT_TAG       string
	CI_TAGGED           bool
	CI_BUILD            bool
)

func Init_build_vars() {
	if CI_COMMIT_SHORT_SHA == "" && CI_COMMIT_BRANCH == "" {
		CI_BUILD = false
	} else {
		CI_BUILD = true
	}
	if CI_COMMIT_TAG == "" {
		CI_TAGGED = false
	} else {
		CI_TAGGED = true
	}
}
