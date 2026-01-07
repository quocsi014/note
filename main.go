package main

import (
	"os"
)

var GlobalConfig *Config

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
