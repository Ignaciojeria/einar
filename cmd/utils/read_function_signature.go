package utils

import (
	"bufio"
	"errors"
	"os"
	"regexp"
)

func ReadFunctionSignature(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// The regex now matches any function signature with any parameters and multiple return values
	regex := regexp.MustCompile(`func\((.*?)\) \(.*\)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)

		// The match will be the first element in the matches slice, because we want to return the whole match
		if len(matches) > 0 {
			return matches[0], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", errors.New("no matching function signature found in file")
}
