package sff8079

import (
	"testing"
)

func TestAccessorMethods(t *testing.T) {
	// Test with 0x1A, 0x25 (mixed options)
	opts := Options{0x1A, 0x25}

	// Test byte 64 accessors
	if !opts.IsPagingImplemented() {
		t.Error("IsPagingImplemented() should be true for 0x1A")
	}
	if !opts.IsRetimerOrCDR() {
		t.Error("IsRetimerOrCDR() should be true for 0x1A")
	}
	if !opts.IsPowerLevelDeclaration() {
		t.Error("IsPowerLevelDeclaration() should be true for 0x1A")
	}

	// Test byte 65 accessors
	if !opts.IsRateSelect() {
		t.Error("IsRateSelect() should be true for 0x25")
	}
	if !opts.IsLossOfSignalInverted() {
		t.Error("IsLossOfSignalInverted() should be true for 0x25")
	}
	if !opts.IsAdditionalPages() {
		t.Error("IsAdditionalPages() should be true for 0x25")
	}

	// Test some false cases
	if opts.IsTxFault() {
		t.Error("IsTxFault() should be false for 0x25")
	}
	if opts.IsLossOfSignalStandard() {
		t.Error("IsLossOfSignalStandard() should be false for 0x25")
	}
}

func TestGetPowerLevel(t *testing.T) {
	tests := []struct {
		name     string
		options  Options
		expected string
	}{
		{
			name:     "Power Level 4",
			options:  Options{0x40, 0x00},
			expected: "Power Level 4",
		},
		{
			name:     "Power Level 3",
			options:  Options{0x20, 0x00},
			expected: "Power Level 3",
		},
		{
			name:     "Power Level 2",
			options:  Options{0x02, 0x00},
			expected: "Power Level 2",
		},
		{
			name:     "Power Level 1 (default)",
			options:  Options{0x00, 0x00},
			expected: "Power Level 1 (or unspecified)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.GetPowerLevel()
			if result != tt.expected {
				t.Errorf("GetPowerLevel() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestOptionsString(t *testing.T) {
	tests := []struct {
		name     string
		options  Options
		expected string
	}{
		{
			name:     "No options",
			options:  Options{0x00, 0x00},
			expected: "Power Level 1",
		},
		{
			name:     "Power Level 4 with features",
			options:  Options{0x58, 0x00}, // Power Level 4 + Paging + Retimer/CDR
			expected: "High Power Level 4, Paging Implemented, Retimer/CDR",
		},
		{
			name:     "Mixed options",
			options:  Options{0x1A, 0x25}, // Power Level 2 + various features
			expected: "Power Level 2, Paging Implemented, Retimer/CDR, Rate Select, Loss of Signal (Inverted), Additional Pages",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.String()
			if result != tt.expected {
				t.Errorf("String() = %v, want %v", result, tt.expected)
			}
		})
	}
}
