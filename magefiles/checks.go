package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type CheckState struct {
	Done []string `json:"done"`
}

func (s *CheckState) IsDone(name string) bool {
	return contains(s.Done, name)
}

func Checks() error {
	// these checks have no requirements, so just always print them and trust the
	// control plane to ignore them if they've already run
	fmt.Println("test")
	fmt.Println("lint")

	// NB: it might be better to just parse this from env vars or something
	// instead
	var state CheckState
	if err := json.NewDecoder(os.Stdin).Decode(&state); err != nil {
		return err
	}

	if state.IsDone("test") && state.IsDone("lint") {
		// only run integration if test and lint are done
		fmt.Println("integration")
	}

	return nil
}

func contains[T comparable](xs []T, y T) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}
