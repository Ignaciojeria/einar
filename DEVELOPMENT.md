# Intall Goreleaser :
go install github.com/goreleaser/goreleaser@latest

# Einar Archetype Code Generator

Cobra installation :
go install github.com/spf13/cobra-cli@latest

Create new command :
/workspace/go/bin/cobra-cli add $command

Compile command line tool application : 
go build -o einar/einar
mkdir example_project
cd example_project

# Check einar-cli version
../einar/einar version

# Create a new application using einar
Create einar project based on your external template : 
../einar/einar init my-project https://github.com/Ignaciojeria/einar-cli-template no-auth
../einar/einar init my-project https://github.com/private_repository user:token
cd ..

# Inside project :
../einar/einar install pubsub
../einar/einar generate subscription mySubscription

../einar/einar install firestore
../einar/einar generate firestore-repository myRepository

Create einar project with all libraries installed by default
./einar init -n=example_project -m=all-in-one

# TODO LIST
Build integration tests that update all dependency tree (updating release folder) if tests pass sucessfully every week (POWERED BY IA)


# Building binaries : 

## For Windows (64-bit):
    GOOS=windows GOARCH=amd64 go build -o einar.exe main.go

## For Mac (64-bit):
    GOOS=darwin GOARCH=amd64 go build -o einar main.go

## For Linux (64-bit):
    GOOS=linux GOARCH=amd64 go build -o einar main.go
