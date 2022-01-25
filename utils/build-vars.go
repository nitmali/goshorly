package utils

import (
	"fmt"
	"time"
)

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

func Print_Starting_Screen() {
	if CI_BUILD {
		if CI_TAGGED {
			fmt.Println("---- Starting goshorly " + CI_COMMIT_TAG + " ----")
		} else {
			fmt.Println("---- Starting goshorly " + CI_COMMIT_SHORT_SHA + " ----")
		}
	} else {
		fmt.Println("---- Starting goshorly " + "unknown version" + " ----")
	}
	time.Sleep(1 * time.Second)
}
