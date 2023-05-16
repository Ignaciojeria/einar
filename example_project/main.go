package main

import (
	"archetype/app/shared/archetype"
	"os"
)

func main() {
	if err := archetype.Setup(); err != nil {
		os.Exit(0)
	}
}
