package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/pflag"
)

func HandleCreate(_args []string, workingPath string) error {
	fs := pflag.NewFlagSet("create", pflag.ContinueOnError)
	title := fs.StringP("title", "t", "", "")
	ext := fs.StringP("ext", "e", "", "")
	c := fs.BoolP("c", "c", false, "")

	fs.Parse(_args)

	filePath, err := createFile(workingPath, fmt.Sprintf(".%s", *ext), *title)
	if err != nil {
		return err
	}

	if !*c {
		err = openInEditor(GlobalConfig.Editor, filePath)
		if err != nil {
			return fmt.Errorf("editor failed: %w", err)
		}

		if *title == "" {
			err := modifyTitle(filePath)
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

	filePath, err := createFile(workingPath, fmt.Sprintf(".%s", ext), *title)
	if err != nil {
		return err
	}

	if !*c {
		err = openInEditor(GlobalConfig.Editor, filePath)
		if err != nil {
			return fmt.Errorf("editor failed: %w", err)
		}

		if *title == "" {
			err := modifyTitle(filePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func openInEditor(editor, filePath string) error {
	parts := strings.Fields(editor)
	if len(parts) == 0 {
		return fmt.Errorf("empty editor command")
	}

	var cmd *exec.Cmd
	if len(parts) > 1 {
		cmd = exec.Command(parts[0], append(parts[1:], filePath)...)
	} else {
		cmd = exec.Command(parts[0], filePath)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: false,
	}

	return cmd.Run()
}

func modifyTitle(filePath string) error {
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
	return changeFileName(filePath, title)
}

func notePrefixNow() string {
	now := time.Now().Format("150405")
	return fmt.Sprintf("note%s:", now)
}
