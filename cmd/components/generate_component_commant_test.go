package components

import (
	"testing"
	"fmt"
	"strings"
)

func TestFeature(t *testing.T){
	componentParts := strings.Split("componentName", "/")
	nestedFolders := strings.Join(componentParts[:len(componentParts)-1], "/")
	if nestedFolders != "" {
		nestedFolders += "/"
	}

	// Extract the componentName here
	componentName := componentParts[len(componentParts)-1]

	type ComponentFile struct {
		SourceFile     string          `json:"source_file"`
		DestinationDir string          `json:"destination_dir"`
	}
	file := ComponentFile{
		DestinationDir:"app/adapter/in/subscription",
	}

	destinationDirParts := strings.Split(file.DestinationDir, "/")
	baseFolder := destinationDirParts[0]
	// Remove the first folder from Dir
	file.DestinationDir =  strings.TrimPrefix(file.DestinationDir, baseFolder+"/")
	fmt.Println("cli.Project: "+"my-project")
	fmt.Println("baseFolder: "+baseFolder)
	fmt.Println("nestedFolders: "+nestedFolders)
	fmt.Println("componentName: "+componentName) // Print the componentName here
	fmt.Println("file.DestinationDir: "+file.DestinationDir)
	fmt.Println("works")
}
