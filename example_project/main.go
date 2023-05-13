package main

import (
	"archetype/app/shared/archetype"
	"archetype/app/shared/config"
	"os"
)

func main() {
	if err := archetype.Setup(config.ArchetypeConfiguration{
		EnableHTTPServer:                  true,
		EnableFirestore:                   false,
		EnablePubSub:                      false,
		EnableRedis:                       false,
		EnablePostgreSQLDB:                false,
	}); err != nil {
		os.Exit(0)
	}
}
