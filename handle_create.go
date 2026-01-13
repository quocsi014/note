package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

func HandleCreate(_args []string, workingPath string) error {
	fs := pflag.NewFlagSet("create", pflag.ContinueOnError)
	title := fs.StringP("title", "t", "", "")
	ext := fs.StringP("ext", "e", "", "")
	c := fs.BoolP("c", "c", false, "")

	err := fs.Parse(_args)
	if err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	filePath, err := CreateNote(workingPath, fmt.Sprintf("%s", *ext), *title)
	if err != nil {
		return fmt.Errorf("failed to create note: %w", err)
	}

	if !*c {
		err = OpenFileInEditor(GlobalConfig.Editor, filePath)
		if err != nil {
			return fmt.Errorf("failed to open editor (%s): %w", GlobalConfig.Editor, err)
		}

		if *title == "" {
			err := ModifyTitle(filePath)
			if err != nil {
				return fmt.Errorf("failed to update note title: %w", err)
			}
		}
	}

	return nil
}

func HandleCreateWithExt(_args []string, workingPath string, ext string) error {
	fs := pflag.NewFlagSet("create", pflag.ContinueOnError)
	title := fs.StringP("title", "t", "", "")
	c := fs.BoolP("c", "c", false, "")

	err := fs.Parse(_args)
	if err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	filePath, err := CreateNote(workingPath, fmt.Sprintf("%s", ext), *title)
	if err != nil {
		return fmt.Errorf("failed to create note with ext %s: %w", ext, err)
	}

	if !*c {
		err = OpenFileInEditor(GlobalConfig.Editor, filePath)
		if err != nil {
			return fmt.Errorf("failed to open editor (%s): %w", GlobalConfig.Editor, err)
		}

		if *title == "" {
			err := ModifyTitle(filePath)
			if err != nil {
				return fmt.Errorf("failed to update note title: %w", err)
			}
		}
	}

	return nil
}


func ModifyTitle(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file for title modification: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		return nil
	}

	title := strings.TrimSpace(scanner.Text())
	err = ChangeTitle(filePath, title)
	if err != nil {
		return fmt.Errorf("failed to change title: %w", err)
	}
	return nil
}
