package base

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func CreateVersion(project string) error {
	versionPath := filepath.Join(project, ".version")
	versionContentBytes, err := ioutil.ReadFile("cmd/base/version/.version")
	if err != nil {
		return fmt.Errorf("error reading version file: %v", err)
	}

	versionContent := string(versionContentBytes)

	err = ioutil.WriteFile(versionPath, []byte(versionContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing version file: %v", err)
	}

	fmt.Printf(".version file generated successfully at %s.\n", versionPath)
	return nil
}
