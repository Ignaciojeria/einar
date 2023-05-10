package archetype

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateMainFile(moduleName string) error {
	// Create the main.go file in the moduleName directory
	mainFilePath := filepath.Join(moduleName, "main.go")
	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		err := fmt.Errorf("error creating main.go file for module %s: %s", moduleName, err)
		fmt.Println(err)
		return err
	}
	defer mainFile.Close()

	// Write the desired content to the main.go file
	mainFileContent := `package main

import "example/app/config"

func main() {
	config.Setup(config.ArchetypeConfiguration{
		EnableHTTPServer: true,
	})
}`
	_, err = io.WriteString(mainFile, mainFileContent)
	if err != nil {
		err := fmt.Errorf("error writing to main.go file for module %s: %s", moduleName, err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("Main.go file generated successfully at %s.\n", mainFilePath)
	return nil
}
