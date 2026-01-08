package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
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

	filePath, err := create(workingPath, *ext, *title)
	if err != nil {
		return err
	}

	if !*c {
		err = openInEditor(GlobalConfig.Editor, filePath)
		if err != nil {
			return fmt.Errorf("editor failed: %w", err)
		}
	}

	return nil
}

func HandleCreateWithExt(_args []string, workingPath string, ext string) error {
	fs := pflag.NewFlagSet("create", pflag.ContinueOnError)
	title := fs.StringP("title", "t", "", "")
	c := fs.BoolP("c", "c", false, "")

	fs.Parse(_args)

	filePath, err := create(workingPath, ext, *title)
	if err != nil {
		return err
	}

	if !*c {
		err = openInEditor(GlobalConfig.Editor, filePath)
		if err != nil {
			return fmt.Errorf("editor failed: %w", err)
		}
	}

	return nil
}

func create(workingPath, ext string, title string) (string, error) {
	title = sanitizeFilename(title)
	fileName := fmt.Sprintf("%s%s", notePrefixNow(), title)
	if ext != "" {
		fileName = fmt.Sprintf("%s.%s", fileName, ext)
	}
	filePath := path.Join(workingPath, fileName)

	_, err := os.OpenFile(filePath, os.O_CREATE, 0744)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func sanitizeFilename(s string) string {
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := s

	for _, char := range invalidChars {
		result = strings.ReplaceAll(result, char, "-")
	}

	result = strings.Trim(result, " .")

	result = strings.Join(strings.Fields(result), " ")

	if len(result) > 100 {
		result = result[:100]
	}

	if result == "" {
		result = untitled
	}

	return result
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
		Setpgid: false, // Không tạo process group mới
	}

	return cmd.Run()
}

func notePrefixNow() string {
	now := time.Now().Format("150405")
	return fmt.Sprintf("note%s:", now)
}
