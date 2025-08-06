package utils

import (
	"testing"
)

func TestSanitizeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "XSS Script Tag",
			input:    "<script>alert('XSS')</script>Hello",
			expected: "&lt;script&gt;alert(XSS)&lt;/script&gt;Hello",
		},
		{
			name:     "SQL Injection Attempt",
			input:    "'; DROP TABLE books; --",
			expected: " DROP TABLE books --",
		},
		{
			name:     "Normal Text with Whitespace",
			input:    "   Normal Book Title   ",
			expected: "Normal Book Title",
		},
		{
			name:     "Special Characters",
			input:    "Book & Author (2024)",
			expected: "Book  Author 2024",
		},
		{
			name:     "Empty String",
			input:    "",
			expected: "",
		},
		{
			name:     "Only Whitespace",
			input:    "   ",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeString(tt.input)
			if result != tt.expected {
				t.Errorf("SanitizeString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestValidateID(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid Positive ID",
			input:    "123",
			expected: true,
		},
		{
			name:     "Valid Single Digit",
			input:    "1",
			expected: true,
		},
		{
			name:     "Invalid Letters",
			input:    "abc",
			expected: false,
		},
		{
			name:     "Invalid Mixed",
			input:    "123abc",
			expected: false,
		},
		{
			name:     "Empty String",
			input:    "",
			expected: false,
		},
		{
			name:     "Negative Number",
			input:    "-1",
			expected: false,
		},
		{
			name:     "Zero",
			input:    "0",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateID(tt.input)
			if result != tt.expected {
				t.Errorf("ValidateID() = %v, want %v", result, tt.expected)
			}
		})
	}
}
