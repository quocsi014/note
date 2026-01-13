package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/fatih/color"
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
	"htm":      NoteTypeHTML,
	"html":     NoteTypeHTML,
	"markdown": NoteTypeMarkdown,
	"md":       NoteTypeMarkdown,
	"json":     NoteTypeJSON,
}

const (
	NoteFileNameTemplate = "%s::%s" // prefix::untitled::name.ext
	UntitledTitle        = "Untitled"
)

var supportedExt = func() []string {
	keys := make([]string, 0, len(extToNoteType))
	for k := range extToNoteType {
		keys = append(keys, k)
	}

	return keys
}()

var NoteTypeToExt = func() map[NoteType]string {
	m := map[NoteType]string{}
	for k, v := range extToNoteType {
		m[v] = k
	}

	return m
}()

func NoteTypeFromExt(ext string) NoteType {
	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	if t, ok := extToNoteType[ext]; ok {
		return t
	}
	return NoteTypeOther
}

type Note struct {
	Name       string
	NoteType   NoteType
	Title      string
	ModifiedAt time.Time
}

func (n *Note) Display(titleLen int) string {
	var title = n.Title
	if titleLen < len(n.Title) {
		title = fmt.Sprintf("%s...", title[:titleLen-3])
	}
	return fmt.Sprintf("%s %-*.*s %s",
		NoteIcon[n.NoteType],
		titleLen, titleLen,
		title,
		color.RGB(96, 96, 96).Sprint(TimeSince(n.ModifiedAt)),
	)
}

func WorkingDir(daysAgo *int) string {
	now := time.Now()
	workingTime := now.AddDate(0, 0, -*daysAgo)

	return workingTime.Format("02012006")
}

func CreateNote(workingPath, ext, title string) (string, error) {
	fileName := CreateFileName(title, ext)
	filePath := path.Join(workingPath, fileName)

	_, err := os.OpenFile(filePath, os.O_CREATE, 0744)
	if err != nil {
		return "", fmt.Errorf("failed to create file %s: %w", filePath, err)
	}

	return filePath, nil
}

func CreateFileName(title, ext string) string {
	title = sanitizeFilename(title)

	if title == "" {
		title = UntitledTitle
	}

	fileName := fmt.Sprintf("%s::%s", notePrefixNow(), title)
	if ext != "" {
		fileName = fmt.Sprintf("%s.%s", fileName, ext)
	}

	return fileName
}

func ChangeTitle(relativePath string, title string) error {
	ext := filepath.Ext(relativePath)
	if len(ext) > 0 {
		ext = ext[1:]
	}

	fileName := CreateFileName(title, ext)

	dir, _ := filepath.Split(relativePath)
	newPath := filepath.Join(dir, fileName)

	err := os.Rename(relativePath, newPath)
	if err != nil {
		return fmt.Errorf("failed to rename file from %s to %s: %w", relativePath, newPath, err)
	}
	return nil
}

func notePrefixNow() string {
	now := time.Now().Format("150405")
	return fmt.Sprintf("note%s", now)
}

func ListNote(workingPath string) ([]*Note, error) {
	files, err := os.ReadDir(workingPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", workingPath, err)
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
