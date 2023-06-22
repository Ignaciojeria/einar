package cmd_tests

import (
	"context"
	"fmt"
	"dagger.io/dagger"
	"os"
)

func EinarVersion(ctx context.Context) error {

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	defer client.Close()
	if err != nil {
		return err
	}

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
