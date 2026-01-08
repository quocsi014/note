package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"slices"

	"github.com/fatih/color"
)

func fileToNote(inf fs.FileInfo) *Note {
	ext := filepath.Ext(inf.Name())
	noteType := NoteTypeFromExt(ext)

	return &Note{
		Name:       inf.Name(),
		NoteType:   noteType,
		Title:      inf.Name()[notePrefixLength : len(inf.Name())-len(ext)],
		ModifiedAt: inf.ModTime(),
	}
}

func HandleListNote(workingPath string) error {
	notes, err := listNote(workingPath)
	if err != nil {
		return err
	}

	titleLenMax := 40
	titleLen := 0
	for _, note := range notes {
		if l := len(note.Title); l > titleLen {
			if l > titleLenMax {
				l = titleLenMax
				note.Title = fmt.Sprintf("%s...", note.Title[:titleLenMax-3])
			}

			titleLen = l
		}
	}

	fmt.Println("")

	for i, note := range notes {
		fmt.Printf("%s. %s\n\n",
			color.New(color.Bold, color.FgGreen).Sprintf("%3.3d", i+1),
			note.Display(titleLen),
		)
	}

	return nil
}

func listNote(workingPath string) ([]*Note, error) {
	files, err := os.ReadDir(workingPath)
	if err != nil {
		return nil, err
	}

	notes := make([]*Note, 0, len(files))
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		reg, err := regexp.Compile(notePrefixPattern)
		if err != nil {
			return nil, err
		}

		isNote := reg.Match([]byte(f.Name()))
		if !isNote {
			continue
		}

		inf, err := f.Info()
		if err != nil {
			return nil, err
		}

		notes = append(notes, fileToNote(inf))
	}

	slices.SortFunc(notes, func(a, b *Note) int {
		if a.ModifiedAt.After(b.ModifiedAt) {
			return 1
		} else {
			return -1
		}
	})

	return notes, nil
}
