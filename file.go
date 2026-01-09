package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func sanitizeFilename(s string) string {
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := s

	for _, char := range invalidChars {
		result = strings.ReplaceAll(result, char, "-")
	}

	result = strings.Trim(result, " .")

	result = strings.Join(strings.Fields(result), " ")

	if len(result) > 100 {
		result = result[:100]
	}

	if result == "" {
		result = untitled
	}

	return result
}

func createFile(workingPath, ext string, title string) (string, error) {
	fileName := createFileName(title, ext)
	filePath := path.Join(workingPath, fileName)

	_, err := os.OpenFile(filePath, os.O_CREATE, 0744)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func changeFileName(relativePath string, title string) error {
	ext := filepath.Ext(relativePath)
	fileName := createFileName(title, ext)

	dir, _ := filepath.Split(relativePath)
	newPath := filepath.Join(dir, fileName)

	return os.Rename(relativePath, newPath)
}

func createFileName(title, ext string) string {
	title = sanitizeFilename(title)
	fileName := fmt.Sprintf("%s%s", notePrefixNow(), title)
	if ext != "" {
		fileName = fmt.Sprintf("%s%s", fileName, ext)
	}

	return fileName
}
