package main

import (
	"context"

	"dagger.io/dagger"
)

func Test(ctx context.Context) error {
	c, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}

	defer c.Close()

	_, err = c.Container().From("alpine").
		WithExec([]string{"echo", "im testing"}).
		ExitCode(ctx)

	return err
}
