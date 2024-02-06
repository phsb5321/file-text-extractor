package main

import (
	"fmt"
	"os"

	"textractor/config"
	"textractor/processor"
)

func main() {
	cfg, err := config.ParseCommandLineArguments(os.Args[1:])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err := processor.ProcessDirectory(cfg); err != nil {
		fmt.Printf("Error processing directory: %v\n", err)
		os.Exit(1)
	}
}
