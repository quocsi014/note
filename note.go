package main

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/pflag"
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

func HandleCommand(args []string) error {

	fs := pflag.NewFlagSet("handle", pflag.ContinueOnError)
	fs.ParseErrorsAllowlist.UnknownFlags = true

	daysAgo := pflag.IntP("ago", "a", 0, "")

	fs.Parse(args)

	workingDir := WorkingDir(daysAgo)
	workingPath := path.Join(GlobalConfig.StorageDir, workingDir)

	os.MkdirAll(workingPath, 0744)

	var cmd string
	if len(fs.Args()) == 0 {
		cmd = "create"
	} else {
		cmd = args[0]
	}

	if slices.Contains(supportedExt, cmd) {
		cmd = "spec-create"
	}

	switch cmd {
	case "ls", "list":
		ListNote(workingPath)
	case "create":
		Create(args, workingPath)
	case "spec-create":
		CreateWithExt(args, workingPath, args[0])
	}

	return nil
}
