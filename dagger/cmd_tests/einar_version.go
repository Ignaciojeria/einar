package cmd_tests

import (
	"context"
	"fmt"
	"dagger.io/dagger"
)

func EinarVersion(ctx context.Context, client *dagger.Client) error {
	src := client.Host().Directory("./host_output")
	val,err :=client.
	Container().
	From("golang:latest").
	WithDirectory("/src",src).
	WithWorkdir("/src").
	WithExec([]string{"./einar","version"}).
	Stdout(ctx)
	fmt.Println(val)
	if err != nil {
		return err
	}
	return nil
}
