package main

import (
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
