# :fire:EinarCLI: Accelerate Your Application Development:fire:

EinarCLI is a blazingly fast :fire: and powerful command-line interface tool designed to supercharge :rocket: your application development process. It helps to streamline and expedite the creation of your applications, enabling you to code faster, smoother, and smarter.

With EinarCLI, get ready to turn up the heat :fire: on your programming skills and develop applications like never before!

## Minimum Requirements

- [Golang 1.20](https://golang.org/doc/install)

## 🔧 Installation
To install EinarCLI, use the following command:

    go install github.com/Ignaciojeria/einar@v1.16.0

To verify your EinarCLI installation, run:

    einar version

## :computer: Usage
#### Starting a new Einar project with the official template
![Diagram](ddd_sketching/einar.drawio.svg)
To use the official EinarCLI template for creating a new project, execute the following command:

    einar init my-project-name https://github.com/Ignaciojeria/einar-cli-template no-auth
    
#### Starting a new Einar project with your custom private template
If you wish to utilize your own corporate private template to initiate a new project, use the following command:

    einar init my-project-name https://github.com/your-corporative-private-einar-cli-template user:token

## 🔧 Installing HTTP Server
We are using Echo to manage HTTP requests efficiently. To install the Echo Server, run the following command in your terminal:

    einar install echo-server
    
#### 🔍 Check Application Port 
To check the port that your HTTP server is using for your application, you can refer to the .env file located in the root directory of your Einar project

#### 🔍 Run Application
To run the application, you can start it as a traditional Golang application by executing the following command:

    go run main.go

This will start your application on the default port specified in the configuration.

#### 🔍 Check Running Application Status
To check the status of the running application, visit the following URL in your web browser:

[http://localhost:8080/health](http://localhost:8080/health) 

![Img](ddd_sketching/check_running_application_status.jpeg)

## :star: Features

(Feature list here)

## :telephone_receiver: Support

(Contact and support information here)

## :handshake: Contributing

(Contribution guidelines here)
