package main

import (
	"fmt"
	"os"
	"strings"
)

func Checks() error {
	// these checks have no requirements, so just always print them and trust the
	// control plane to ignore them if they've already run
	fmt.Println("test")
	fmt.Println("lint")

	var doneChecks CheckSet
	if done := os.Getenv("DONE"); done != "" {
		doneChecks = strings.Fields(done)
	}

	if doneChecks.Contains("test") && doneChecks.Contains("lint") {
		// only run integration if test and lint are done
		fmt.Println("integration")
	}

	return nil
}

type CheckSet []string

func (checks CheckSet) Contains(y string) bool {
	for _, x := range checks {
		if x == y {
			return true
		}
	}
	return false
}
