package main

import (
	"context"
	"dagger/pipelines"
	"fmt"
	"os"
	"dagger.io/dagger"
	"github.com/joho/godotenv"
)
var version = "1.0.1"

func main() {
	err := godotenv.Load()
	if err != nil {
	  fmt.Println(err)
	}
	os.Setenv("EINAR_CLI_VERSION",version)
	// initialize Dagger client
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	defer client.Close()
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	if err := pipelines.Build(ctx,client); err != nil {
		fmt.Println(err)
		return
	}
	tagName := "v"+version
	if err := pipelines.SetReleaseTag(ctx,tagName,"Release version v"+version);err!=nil{
		fmt.Println(err)
		//return
	}

	if err := pipelines.PublishRelease(ctx,client,tagName); err != nil {
		fmt.Println(err)
		return
	}


}
