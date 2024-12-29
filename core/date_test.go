package core

import "testing"

func TestIsValidPartialISO8601(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"test", false},       // Valid case (year 0 is acceptable in ISO 8601)
		{"abcd-ef", false},    // Invalid format (non-numeric)
		{"2024-2", false},     // Invalid format (single digit month)
		{"2024-12-01", false}, // Invalid format (extra day part)
		{"2024-00", false},    // Invalid month (zero month)
		{"2024-13", false},    // Invalid month
		{"2024-12", true},     // Valid case
		{"0000-01", true},     // Valid case (year 0 is acceptable in ISO 8601)
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := IsValidPartialISO8601(tt.input)
			if result != tt.expected {
				t.Errorf("isValidPartialISO8601(%q) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
