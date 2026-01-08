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
			return fmt.Errorf("Unknow command: %s", fs.Args()[0])
		}
	}

	noteIndex := *indexFlag - 1

	if indexValue > 0 {
		noteIndex = indexValue - 1
	}

	notes, err := listNote(workingPath)
	if err != nil {
		return nil
	}

	if noteIndex < 0 || noteIndex > len(notes)-1 {
		noteIndex = len(notes) - 1
	}

	filePath := path.Join(workingPath, notes[noteIndex].Name)

	err = openInEditor(GlobalConfig.Editor, filePath)
	if err != nil {
		return err
	}

	return nil
}
