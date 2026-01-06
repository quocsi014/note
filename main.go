package main

import (
	"os"
	"path"

	"github.com/spf13/pflag"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	daysAgo := pflag.IntP("ago", "a", 0, "")

	pflag.Parse()

	args := pflag.Args()

	workingDir := WorkingDir(daysAgo)
	workingPath := path.Join(config.StorageDir, workingDir)

	os.MkdirAll(workingPath, 0744)

	HandleCommand(args, workingPath)
}
