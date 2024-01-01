package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestReadTemplateModuleName(t *testing.T) {
	// Crear un directorio temporal
	tempDir, err := ioutil.TempDir("", "test-module")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Crear un archivo go.mod de prueba en el directorio temporal
	moduleName := "example.com/my/module"
	goModContent := "module " + moduleName
	err = ioutil.WriteFile(filepath.Join(tempDir, "go.mod"), []byte(goModContent), 0666)
	if err != nil {
		t.Fatalf("Failed to write go.mod: %v", err)
	}

	// Llamar a la función ReadTemplateModuleName
	resultModuleName, err := ReadTemplateModuleName(tempDir)
	if err != nil {
		t.Errorf("ReadTemplateModuleName returned an error: %v", err)
	}
	if resultModuleName != moduleName {
		t.Errorf("Expected module name %q, got %q", moduleName, resultModuleName)
	}
}
