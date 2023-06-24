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

## :hammer: HTTP Server - Inbound Adapter
An HTTP Inbound Adapter is the same as a controller in a web application. To create a new HTTP inbound adapter, we need to execute the following commands:

    einar generate get-controller  get-customer
    einar generate post-controller post-customer
    einar generate put-controller put-customer
    einar generate patch-controller patch-customer
    einar generate delete-controller delete-customer

####  🔍 Directory Structure
When running the Einar commands to generate controllers, the following directory structure will be created in your project:

```bash
/app
  /adapter
    /in
      /controller
        - get_customer.go   # Generated with 'einar generate get-controller get-customer'
        - post_customer.go  # Generated with 'einar generate post-controller post-customer'
        - put_customer.go  # Generated with 'einar generate put-controller put-customer'
        - patch_customer.go # Generated with 'einar generate patch-controller patch-customer'
        - delete_customer.go # Generated with 'einar generate delete-controller delete-customer'
```

#### 🔍 File structure
For example, when you generate a controller using the `einar generate get-controller get-customer` command, Einar CLI will create a new .go file with the following structure:

```go
package controller

import (
	"my-project/app/shared/archetype/container"
	einar "my-project/app/shared/archetype/echo_server"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func init() {
	container.InjectInboundAdapter(func() error {
		einar.Echo.GET("/INSERT_YOUR_PATTERN_HERE", getCustomer)
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func getCustomer(c echo.Context) error {
	//PUT HERE YOUR HANDLING PROCESS
	return c.JSON(http.StatusOK, "INSERT_YOUR_CUSTOM_RESPONSE")
}
```

## ⭐ Features

🔧 **Core Concepts**:
   - No Hidden Dependencies: Ensures transparency in the code dependencies.
   - Dependency Container: Manages object instances and dependencies.
   - Inversion of Control: Decouples modules and promotes code flexibility.

📡 **Google PubSub Integration**:
   - Client Installation & Setup: Streamlines the setup process for Google PubSub.
   - Subscriptions Generation: Easily generate subscriptions.
   - Publishers Generation: Quickly set up publishers for message distribution.

🌐 **Echo Server Integration**:
   - Echo Server Installation: Simplifies the installation process of the Echo Server.
   - Controllers Generation: Generates controllers for handling HTTP requests.

🔥 **Google Firestore Integration**:
   - Firestore Installation & Setup: Streamlines Firestore database setup.
   - Repositories Generation: Quickly generate repositories.

👨‍💻 **Code Ownership and Control**:
   - Full Control of Generated Code: You have complete control over the code that is generated.
   - Code Ownership: The generated code is yours to keep and use as you wish.


## :telephone_receiver: Support

(Contact and support information here)

## :handshake: Contributing

(Contribution guidelines here)
