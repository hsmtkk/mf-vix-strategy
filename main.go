package main

import (
	"log"

	"github.com/hsmtkk/mf-vix-strategy/command"
)

func main() {
	cmd := command.Command
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
