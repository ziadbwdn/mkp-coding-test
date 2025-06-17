package utils

import (
	"regexp"
	"strings"
)

// ValidateTerminalCode validates terminal code format
func ValidateTerminalCode(code string) bool {
	// Terminal code should be alphanumeric, 2-10 characters
	matched, _ := regexp.MatchString("^[A-Z0-9]{2,10}$", strings.ToUpper(code))
	return matched
}

// ValidateUsername validates username format
func ValidateUsername(username string) bool {
	// Username should be alphanumeric with underscores, 3-50 characters
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_]{3,50}$", username)
	return matched
}

// SanitizeString removes extra whitespace and limits length
func SanitizeString(input string, maxLength int) string {
	cleaned := strings.TrimSpace(input)
	if len(cleaned) > maxLength {
		return cleaned[:maxLength]
	}
	return cleaned
}