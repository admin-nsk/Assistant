package main

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env") //nolint:errcheck
}

func main() {
	if err := Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
