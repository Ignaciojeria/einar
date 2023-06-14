# Intall Goreleaser :
go install github.com/goreleaser/goreleaser@latest


# Tag geneartion
 1.- git tag -d v1.0.3
 2.- git tag v1.0.3
 3.- git push origin v1.0.3
 4.- go run ./dagger/main.go

# Einar einar Code Generator Installation
go install github.com/Ignaciojeria/einar@v1.14.0

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

