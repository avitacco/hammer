package main

import (
	"os"

	"github.com/avitacco/hammer/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}
