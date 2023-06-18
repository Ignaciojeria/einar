package main

import (
	"context"
	"github.com/Ignaciojeria/einar/dagger/build"
	"github.com/Ignaciojeria/einar/dagger/cmd_tests"
	//"github.com/Ignaciojeria/einar/dagger/release"
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
	
	container,err := build.Binary(ctx,client);

	if err != nil {
		fmt.Println(err)
		return
	}

	if err := cmd_tests.EinarVersion(ctx,container); err!=nil{
		fmt.Println(err)
		return
	}
	
	if err := cmd_tests.EinarInit(ctx,container); err!=nil{
		fmt.Println(err)
		return
	}

	/*
	if err := release.Publish(ctx,client); err != nil {
		fmt.Println(err)
		return
	}*/
	
}
