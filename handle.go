package main

import (
	"os"
	"path"
	"slices"
	"time"

	"github.com/spf13/pflag"
)

func HandleCommand(args []string) error {
	fs := pflag.NewFlagSet("handle", pflag.ContinueOnError)
	daysAgo := fs.IntP("ago", "a", 0, "")
	fs.ParseErrorsAllowlist.UnknownFlags = true

	err := fs.Parse(args)
	if err != nil {
		return err
	}

	now := time.Now()
	workingTime := now.AddDate(0, 0, -*daysAgo)
	GlobalWorkingTime = workingTime

	workingDir := workingTime.Format("02012006")
	workingPath := path.Join(GlobalConfig.StorageDir, workingDir)

	os.MkdirAll(workingPath, 0744)

	cmd := extractPrimaryCommand(fs.Args())

	switch cmd {
	case "ls", "list":
		return HandleListNote(workingPath)
	case "c", "create":
		return HandleCreate(args, workingPath)
	case "spec-create":
		return HandleCreateWithExt(args, workingPath, args[0])
	case "o", "open":
		return HandleOpen(args[1:], workingPath)
	}

	return nil
}

func extractPrimaryCommand(args []string) string {
	var cmd string
	if len(args) == 0 {
		cmd = "create"
	} else {
		cmd = args[0]
	}

	if slices.Contains(supportedExt, cmd) {
		cmd = "spec-create"
	}

	return cmd
}
