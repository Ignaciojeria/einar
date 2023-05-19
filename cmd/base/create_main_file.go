package base

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateMainFile() error {
	// Create the main.go file in the moduleName directory
	mainFilePath := filepath.Join("main.go")

	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		err := fmt.Errorf("error creating main.go file for project %s:", err)
		fmt.Println(err)
		return err
	}
	defer mainFile.Close()

	// Write the desired content to the main.go file
	mainFileContent := fmt.Sprintf(`package main

import (
	"archetype/app/shared/archetype"
	"os"
)

func main() {
	if err := archetype.Setup(); err != nil {
		os.Exit(0)
	}
}
`)

	_, err = io.WriteString(mainFile, mainFileContent)
	if err != nil {
		err := fmt.Errorf("error writing to main.go file for project %s: ", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("Main.go file generated successfully at %s.\n", mainFilePath)
	return nil
}
