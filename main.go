package main

import (
	"os"
	"time"
)

var (
	GlobalConfig *Config
	GlobalWorkingTime time.Time
)

func main() {
	var err error
	GlobalConfig, err = LoadConfig()
	if err != nil {
		panic(err)
	}

	args := os.Args

	err = HandleCommand(args[1:])
	if err != nil {
		panic(err)
	}
}
