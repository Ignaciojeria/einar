package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddImportStatement(filePath, importPath string, tagFolder string) error {
	if tagFolder > "4.2.0" {
		//Automatic imports are not needed for future einar templates.
		return nil
	}
	importPath = strings.ReplaceAll(importPath, "\\", "/")
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	var inImportBlock bool
	scanner := bufio.NewScanner(file)
	alreadyImported := false
	for scanner.Scan() {
		line := scanner.Text()
		if inImportBlock && strings.TrimSpace(line) == ")" {
			inImportBlock = false
			// add the import just before the closing parenthesis of the import block
			if !alreadyImported {
				lines = append(lines, fmt.Sprintf("\t_ \"%s\"", importPath))
			}
		}
		// check if the import already exists
		if inImportBlock && strings.Contains(line, importPath) {
			alreadyImported = true
		}
		lines = append(lines, line)
		if strings.Contains(line, "import (") {
			inImportBlock = true
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	// write the updated lines back to the file
	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		fmt.Fprintln(file, line)
	}

	return nil
}
