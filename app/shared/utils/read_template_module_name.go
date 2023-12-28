package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func ReadTemplateModuleName(templatePath string) (string, error) {
	// Construir la ruta completa al archivo go.mod
	modFilePath := templatePath + "/go.mod"

	// Abrir el archivo go.mod
	file, err := os.Open(modFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Leer el archivo línea por línea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			// Extraer el nombre del módulo
			moduleName := strings.TrimSpace(line[len("module "):])
			return moduleName, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Si llegamos aquí, no se encontró la línea del módulo
	return "", errors.New("module declaration not found in go.mod")
}
