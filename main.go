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

	if err := pipelines.CreateReleaseTag(ctx,"v"+version,"Release version v"+version);err!=nil{
		fmt.Println(err)
		return
	}


}
