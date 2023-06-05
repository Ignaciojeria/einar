# Einar Archetype Code Generator

Cobra installation :
go install github.com/spf13/cobra-cli@latest

Create new command :
/workspace/go/bin/cobra-cli add $command

Compile command line tool application : 
go build -o einar/einar
mkdir example_project
cd example_project

# Create a new application using einar
Create einar project with base installations only in default mode
../einar/einar init --repository https://github.com/Ignaciojeria/einar-cli-template
cd ..

# Inside project :
../einar/einar install pubsub
../einar/einar generate subscription mySubscription

Create einar project with all libraries installed by default
./einar init -n=example_project -m=all-in-one

# TODO LIST
Build integration tests that update all dependency tree (updating release folder) if tests pass sucessfully every week (POWERED BY IA)