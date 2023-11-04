package utils

import (
	"fmt"
	"testing"
)

func TestListRelativePathsHardcodedPath(t *testing.T) {
	// Hardcoded path for testing purposes
	basePath := `C:/Users/ignac/Documents/_git/einar/dagger`

	// Call the function to get the list of relative paths
	paths, err := ListRelativePaths(basePath)
	if err != nil {
		t.Fatalf("ListRelativePaths returned an error: %v", err)
	}

	fmt.Println(paths)
}
