package main

import (
	"context"
	"dagger/pipelines"
	"fmt"
)

func main() {
	if err := pipelines.Build(context.Background()); err != nil {
		fmt.Println(err)
	}
}
