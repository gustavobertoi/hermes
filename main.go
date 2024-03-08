package main

import (
	"log"

	"github.com/gustavobertoi/hermes/cmd"
	"github.com/gustavobertoi/hermes/config"
)

func main() {
	_, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	cmd.Execute()
}
