package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func fileToNote(inf fs.FileInfo) *Note {
	ext := filepath.Ext(inf.Name())
	noteType := NoteTypeFromExt(ext)

	return &Note{
		NoteType:  noteType,
		Title:     inf.Name()[len(noteFilePrefix):],
		CreatedAt: inf.ModTime().Unix(),
	}
}

func ListNote(workingPath string) error {
	files, err := os.ReadDir(workingPath)
	if err != nil {
		return err
	}

	notes := make([]*Note, 0, len(files))
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		reg, err := regexp.Compile(`^note\:.*`)
		if err != nil {
			return err
		}

		isNote := reg.Match([]byte(f.Name()))
		if !isNote {
			continue
		}

		inf, err := f.Info()
		if err != nil {
			return err
		}

		notes = append(notes, fileToNote(inf))
	}

	for i, note := range notes {
		fmt.Printf("%s %s%s%d%s. %s\n", NoteIcon[note.NoteType], ANSIFgGreen, ANSIBold, i+1, ANSIReset, note.Title)
	}

	return nil
}
