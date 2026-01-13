package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var (
	GlobalConfig *Config
	GlobalWorkingTime time.Time
)

func main() {
	var err error
	GlobalConfig, err = LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", color.RedString("Config Error"), err)
		os.Exit(1)
	}

	args := os.Args

	err = HandleCommand(args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", color.RedString("Error"), err)
		os.Exit(1)
	}
}
