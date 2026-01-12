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

	fs.Parse(_args)

	filePath, err := CreateNote(workingPath, fmt.Sprintf("%s", *ext), *title)
	if err != nil {
		return err
	}

	if !*c {
		err = OpenFileInEditor(GlobalConfig.Editor, filePath)
		if err != nil {
			return fmt.Errorf("editor failed: %w", err)
		}

		if *title == "" {
			err := ModifyTitle(filePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func HandleCreateWithExt(_args []string, workingPath string, ext string) error {
	fs := pflag.NewFlagSet("create", pflag.ContinueOnError)
	title := fs.StringP("title", "t", "", "")
	c := fs.BoolP("c", "c", false, "")

	fs.Parse(_args)

	filePath, err := CreateNote(workingPath, fmt.Sprintf("%s", ext), *title)
	if err != nil {
		return err
	}

	if !*c {
		err = OpenFileInEditor(GlobalConfig.Editor, filePath)
		if err != nil {
			return fmt.Errorf("editor failed: %w", err)
		}

		if *title == "" {
			err := ModifyTitle(filePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}


func ModifyTitle(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		return nil
	}

	title := strings.TrimSpace(scanner.Text())
	return ChangeTitle(filePath, title)
}
