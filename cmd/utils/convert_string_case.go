package utils

import (
	"regexp"
	"strings"
)

// ConvertStringCase converts a string to a given case.
// Supported cases are: snake_case, PascalCase and camelCase.
func ConvertStringCase(s, caseType string) string {
	switch caseType {
	case "snake_case":
		return toSnakeCase(s)
	case "PascalCase":
		return toPascalCase(s)
	case "camelCase":
		return toCamelCase(s)
	case "kebab":
		return toKebabCase(s)
	default:
		return s
	}
}
func toKebabCase(s string) string {
	re := regexp.MustCompile(`([a-z])([A-Z])`)
	s = re.ReplaceAllStringFunc(s, func(match string) string {
		return match[:1] + "-" + match[1:]
	})
	return strings.ToLower(s)
}

func toSnakeCase(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), "-", "_")
}

func toPascalCase(s string) string {
	words := strings.Split(s, "-")
	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func toCamelCase(s string) string {
	s = toPascalCase(s)
	return strings.ToLower(s[:1]) + s[1:]
}
