package business

import (
	"fmt"
	"strings"
	"testing"
)

func TestExtractTac(t *testing.T) {
	templateFilePath := `C:\Users\ignac\go\bin\github.com\Ignaciojeria\einar-cli-template\1.1.0`
	normalizedPath := strings.ReplaceAll(templateFilePath, "\\", "/")
	pathParts := strings.Split(normalizedPath, "/")
	if len(pathParts) < 1 {

	}
	latestGitTag := pathParts[len(pathParts)-1]

	fmt.Println(latestGitTag)
}
