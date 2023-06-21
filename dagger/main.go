package main

import (
	"context"
	"github.com/Ignaciojeria/einar/dagger/export"
	"github.com/Ignaciojeria/einar/dagger/cmd_tests"

	"fmt"
	"os"
	"dagger.io/dagger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  fmt.Println(err)
	}
	// initialize Dagger client
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	defer client.Close()
	
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := export.Binary(ctx,client); err!=nil{
		panic(err)
	}
	
	if err := cmd_tests.EinarVersion(ctx,client); err!=nil{
		panic(err)
	}

	if err := cmd_tests.EinarInit(ctx,client); err!=nil{
		panic(err)
	}

	if err := cmd_tests.EinarInstall(ctx,client); err!=nil{
		panic(err)
	}

	if err := cmd_tests.EinarGenerate(ctx,client); err!=nil{
		panic(err)
	}

}
