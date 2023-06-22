package main

import (
	"context"
	"github.com/Ignaciojeria/einar/dagger/cmd_tests"
	"github.com/Ignaciojeria/einar/dagger/release"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  fmt.Println(err)
	}
	ctx := context.Background()
	if err := cmd_tests.BuildEinarCli(ctx); err!=nil{
		panic(err)
	}
	if err := cmd_tests.EinarVersion(ctx); err!=nil{
		panic(err)
	}
	if err := cmd_tests.EinarInit(ctx); err!=nil{
		panic(err)
	}
	if err := cmd_tests.EinarInstall(ctx); err!=nil{
		panic(err)
	}
	if err := cmd_tests.EinarGenerate(ctx); err!=nil{
		panic(err)
	}
	if err := release.Publish(ctx); err!=nil{
		panic(err)
	}
}
