package sff8636

import (
	"testing"
)

func TestOptionsDecode(t *testing.T) {
	// Test case 1: No options implemented
	options := Options{0x00, 0x00, 0x00}
	byte0, byte1, byte2 := options.Decode()

	// Verify all bits are false
	if byte0.LPModeTxDisConfigurable || byte0.IntLRxLOSLConfigurable ||
		byte0.TxInputAdaptiveEqualizers || byte0.TxInputEqualizersAuto ||
		byte0.TxInputEqualizersFixed || byte0.RxOutputEmphasisFixed ||
		byte0.RxOutputAmplitudeFixed {
		t.Error("Expected all byte 0 options to be false")
	}

	if byte1.TxCDROnOffControl || byte1.RxCDROnOffControl ||
		byte1.TxCDRLossOfLock || byte1.RxCDRLossOfLock ||
		byte1.RxSquelchDisable || byte1.RxOutputDisable ||
		byte1.TxSquelchDisable || byte1.TxSquelch {
		t.Error("Expected all byte 1 options to be false")
	}

	if byte2.MemoryPage02 || byte2.MemoryPage01 ||
		byte2.RateSelectImplemented || byte2.TxDisableImplemented ||
		byte2.TxFaultImplemented || byte2.TxSquelchType ||
		byte2.TxLossOfSignal || byte2.Pages20_21Implemented {
		t.Error("Expected all byte 2 options to be false")
	}

	// Test case 2: Some options implemented
	options = Options{0x10, 0x80, 0x20}
	byte0, byte1, byte2 = options.Decode()

	// Verify specific bits are set
	if !byte0.TxInputAdaptiveEqualizers {
		t.Error("Expected TxInputAdaptiveEqualizers to be true")
	}
	if !byte1.TxCDROnOffControl {
		t.Error("Expected TxCDROnOffControl to be true")
	}
	if !byte2.RateSelectImplemented {
		t.Error("Expected RateSelectImplemented to be true")
	}

	// Test case 3: All options implemented
	options = Options{0xFF, 0xFF, 0xFF}
	byte0, byte1, byte2 = options.Decode()

	// Verify all bits are true
	if !byte0.LPModeTxDisConfigurable || !byte0.IntLRxLOSLConfigurable ||
		!byte0.TxInputAdaptiveEqualizers || !byte0.TxInputEqualizersAuto ||
		!byte0.TxInputEqualizersFixed || !byte0.RxOutputEmphasisFixed ||
		!byte0.RxOutputAmplitudeFixed {
		t.Error("Expected all byte 0 options to be true")
	}

	if !byte1.TxCDROnOffControl || !byte1.RxCDROnOffControl ||
		!byte1.TxCDRLossOfLock || !byte1.RxCDRLossOfLock ||
		!byte1.RxSquelchDisable || !byte1.RxOutputDisable ||
		!byte1.TxSquelchDisable || !byte1.TxSquelch {
		t.Error("Expected all byte 1 options to be true")
	}

	if !byte2.MemoryPage02 || !byte2.MemoryPage01 ||
		!byte2.RateSelectImplemented || !byte2.TxDisableImplemented ||
		!byte2.TxFaultImplemented || !byte2.TxSquelchType ||
		!byte2.TxLossOfSignal || !byte2.Pages20_21Implemented {
		t.Error("Expected all byte 2 options to be true")
	}
}

func TestOptionsString(t *testing.T) {
	// Test case 1: No options implemented
	options := Options{0x00, 0x00, 0x00}
	expected := "No options implemented"
	if result := options.String(); result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	// Test case 2: Some options implemented
	options = Options{0x10, 0x80, 0x20}
	result := options.String()
	expectedParts := []string{
		"Tx input adaptive equalizers freeze capable",
		"Tx CDR On/Off Control",
		"Rate select implemented",
	}

	for _, part := range expectedParts {
		if !contains(result, part) {
			t.Errorf("Expected result to contain '%s', got '%s'", part, result)
		}
	}
}

func TestOptionsList(t *testing.T) {
	// Test case 1: No options implemented
	options := Options{0x00, 0x00, 0x00}
	if result := options.List(); len(result) != 0 {
		t.Errorf("Expected empty list, got %v", result)
	}

	// Test case 2: Some options implemented
	options = Options{0x10, 0x80, 0x20}
	result := options.List()
	expectedCount := 3
	if len(result) != expectedCount {
		t.Errorf("Expected %d options, got %d", expectedCount, len(result))
	}
}

func TestOptionsStringCol(t *testing.T) {
	// Test case 1: No options implemented
	options := Options{0x00, 0x00, 0x00}
	result := options.StringCol()

	// Should contain the header with hex values
	expectedHeader := "Options [193-195]"
	if !contains(result, expectedHeader) {
		t.Errorf("Expected result to contain '%s', got '%s'", expectedHeader, result)
	}

	// Should contain hex values
	expectedHex := "0x00 0x00 0x00"
	if !contains(result, expectedHex) {
		t.Errorf("Expected result to contain '%s', got '%s'", expectedHex, result)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) && (s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			contains(s[1:], substr))))
}

// TestOptionsWithRealData tests the Options functionality with actual EEPROM data
func TestOptionsWithRealData(t *testing.T) {
	// This test would require access to the main sff package
	// For now, we'll test the Options type directly with sample data

	// Sample data that might represent real EEPROM values
	// These values are based on common SFF-8636 transceiver capabilities
	options := Options{0x10, 0x80, 0x20}

	// Verify the decoding works correctly
	byte0, byte1, byte2 := options.Decode()

	// Byte 0: 0x10 = 00010000 in binary
	// Bit 4 should be set (Tx input adaptive equalizers freeze capable)
	if !byte0.TxInputAdaptiveEqualizers {
		t.Error("Expected TxInputAdaptiveEqualizers to be true for 0x10")
	}

	// Byte 1: 0x80 = 10000000 in binary
	// Bit 7 should be set (Tx CDR On/Off Control implemented)
	if !byte1.TxCDROnOffControl {
		t.Error("Expected TxCDROnOffControl to be true for 0x80")
	}

	// Byte 2: 0x20 = 00100000 in binary
	// Bit 5 should be set (Rate select implemented)
	if !byte2.RateSelectImplemented {
		t.Error("Expected RateSelectImplemented to be true for 0x20")
	}

	// Test string representation
	result := options.String()
	expectedParts := []string{
		"Tx input adaptive equalizers freeze capable",
		"Tx CDR On/Off Control",
		"Rate select implemented",
	}

	for _, part := range expectedParts {
		if !contains(result, part) {
			t.Errorf("Expected result to contain '%s', got '%s'", part, result)
		}
	}
}
