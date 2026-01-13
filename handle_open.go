package main

import (
	"fmt"
	"path"
	"strconv"

	"github.com/spf13/pflag"
)

func HandleOpen(_args []string, workingPath string) error {
	fs := pflag.NewFlagSet("open", pflag.ContinueOnError)
	indexFlag := fs.IntP("index", "i", 0, "")

	err := fs.Parse(_args)
	if err != nil {
		return err
	}

	var indexValue = -1
	if len(fs.Args()) > 0 {
		var atoiErr error
		indexValue, atoiErr = strconv.Atoi(fs.Args()[0])
		if atoiErr != nil {
			return fmt.Errorf("invalid note index: %w", atoiErr)
		}
	}

	noteIndex := *indexFlag - 1

	if indexValue > 0 {
		noteIndex = indexValue - 1
	}

	notes, err := ListNote(workingPath)
	if err != nil {
		return fmt.Errorf("failed to list notes: %w", err)
	}

	if noteIndex < 0 || noteIndex > len(notes)-1 {
		noteIndex = len(notes) - 1
	}

	filePath := path.Join(workingPath, notes[noteIndex].Name)

	err = OpenFileInEditor(GlobalConfig.Editor, filePath)
	if err != nil {
		return err
	}

	if notes[noteIndex].Title == UntitledTitle {
		err := ModifyTitle(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
