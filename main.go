package main

import (
	"os"
)

func main() {
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}
