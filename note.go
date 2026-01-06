package main

import (
	"strings"
	"time"
)

type NoteType uint8

const (
	NoteTypeOther NoteType = iota
	NoteTypeTXT
	NoteTypeHTML
	NoteTypeMarkdown
	NoteTypeJSON
)

var extToNoteType = map[string]NoteType{
	"txt":      NoteTypeTXT,
	"html":     NoteTypeHTML,
	"htm":      NoteTypeHTML,
	"md":       NoteTypeMarkdown,
	"markdown": NoteTypeMarkdown,
	"json":     NoteTypeJSON,
}

func NoteTypeFromExt(ext string) NoteType {
	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	if t, ok := extToNoteType[ext]; ok {
		return t
	}
	return NoteTypeOther
}

type Note struct {
	NoteType  NoteType
	Title     string
	CreatedAt int64
}

func WorkingDir(daysAgo *int) string {
	now := time.Now()
	workingTime := now.AddDate(0, 0, -*daysAgo)

	return workingTime.Format("02012006")
}

func HandleCommand(args []string, workingPath string) error {
	cmd := args[0]

	switch cmd {
	case "ls", "list":
		ListNote(workingPath)
	}
	return nil
}
