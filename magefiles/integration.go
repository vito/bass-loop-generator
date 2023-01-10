package main

import (
	"context"

	"dagger.io/dagger"
)

func Integration(ctx context.Context) error {
	c, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}

	defer c.Close()

	_, err = c.Container().From("alpine").
		WithExec([]string{"echo", "im running super expensive tests"}).
		ExitCode(ctx)

	return err
}
