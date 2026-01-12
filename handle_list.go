package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"time"

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
	notes, err := ListNote(workingPath)
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

	printTitle()
	fmt.Println("")

	if len(notes) > 0 {
		for i, note := range notes {
			fmt.Printf("%s. %s\n\n",
				color.New(color.Bold, color.FgGreen).Sprintf("%3.3d", i+1),
				note.Display(titleLen),
			)
		}
	} else {
		fmt.Print("empty")
	}

	return nil
}

func printTitle() {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	if GlobalWorkingTime.After(today) {
		fmt.Println("Today's notes")
		return
	}

	if today.Sub(GlobalWorkingTime) <= time.Hour*24 {
		fmt.Println("Yesterday's notes")
		return
	}

	fmt.Printf("Notes dated %s\n", GlobalWorkingTime.Format("02 Jan"))	
}
