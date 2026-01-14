package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func OpenFileInEditor(editor, filePath string) error {
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

	return cmd.Run()
}
