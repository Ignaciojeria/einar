# Intall Goreleaser :
go install github.com/goreleaser/goreleaser@latest


# Tag geneartion
 1.- git tag -d v1.0.1
 2.- git tag v1.0.1
 3.- git push origin v1.0.1 here : package pipelines

# Einar einar Code Generator Installation

Cobra installation :
go install github.com/spf13/cobra-cli@latest

Create new command :
/workspace/go/bin/cobra-cli add $command

Compile command line tool application : 
go run main.go
cd build

# Check einar-cli version
einar version

# Create a new application using einar
Create einar project based on your external template : 
einar init my-project https://github.com/Ignaciojeria/einar-cli-template no-auth
einar init my-project https://github.com/private_repository user:token
cd ..

# Inside project :
einar install pubsub
einar generate subscription mySubscription

einar install firestore
einar generate firestore-repository myRepository

# Building binaries : 

## For Windows (64-bit):
    GOOS=windows GOARCH=amd64 go build -o einar.exe main.go

## For Mac (64-bit):
    GOOS=darwin GOARCH=amd64 go build -o einar main.go

## For Linux (64-bit):
    GOOS=linux GOARCH=amd64 go build -o einar main.go
