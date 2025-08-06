package utils

import (
	"html"
	"regexp"
	"strings"
)

// SanitizeString cleans user input to prevent XSS and other attacks
func SanitizeString(input string) string {
	// Remove HTML tags and escape HTML entities
	input = html.EscapeString(input)

	// Trim whitespace
	input = strings.TrimSpace(input)

	// Remove potentially dangerous characters for SQL (defense in depth)
	// Note: This is extra protection, GORM already handles SQL injection
	dangerousChars := regexp.MustCompile(`[<>\"'%;()&+]`)
	input = dangerousChars.ReplaceAllString(input, "")

	return input
}

// ValidateID validates that an ID is a positive integer
func ValidateID(id string) bool {
	if id == "" {
		return false
	}

	// Check if contains only digits
	matched, _ := regexp.MatchString(`^\d+$`, id)
	return matched
}
