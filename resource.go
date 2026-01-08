package main

import "github.com/fatih/color"

var NoteIcon = map[NoteType]string{
	NoteTypeOther:    color.RGB(200, 200, 2000).Sprint("󰡯"),
	NoteTypeTXT:      color.RGB(0, 102, 102).Sprint(""),
	NoteTypeHTML:     color.RGB(255, 128, 0).Sprint(""),
	NoteTypeMarkdown: color.RGB(0, 128, 255).Sprint(""),
	NoteTypeJSON:     color.RGB(204, 204, 0).Sprint(""),
}
