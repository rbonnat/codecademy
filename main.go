package main

import (
	"context"
	"log"
	"os"

	"github.com/rbonnat/codecademy/configuration"
	"github.com/rbonnat/codecademy/envvarstore"
	"github.com/rbonnat/codecademy/httpserver"
)

func main() {
	// Initialize configuration
	cfg, err := configuration.Load(envvarstore.New())
	if err != nil {
		log.Printf("Error while loading configuration: %v", err)
		os.Exit(1)
	}

	// Launch http server
	err = httpserver.Run(context.TODO(), cfg)
	if err != nil {
		log.Fatal(err)
	}
}
