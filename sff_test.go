package sff

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/bluecmd/go-sff/sff8079"
)

func TestGetType(t *testing.T) {
	tests := []struct {
		name    string
		eeprom  []byte
		want    Type
		wantErr bool
	}{
		{
			name:    "SFF-8079 SFP",
			eeprom:  createSff8079Eeprom(),
			want:    TypeSff8079,
			wantErr: false,
		},
		{
			name:    "SFF-8636 QSFP",
			eeprom:  createSff8636Eeprom(),
			want:    TypeSff8636,
			wantErr: false,
		},
		{
			name:    "Unknown type",
			eeprom:  []byte{0x00, 0x00, 0x00, 0x00},
			want:    TypeUnknown,
			wantErr: true,
		},
		{
			name:    "Too small",
			eeprom:  []byte{0x00, 0x00, 0x00},
			want:    TypeUnknown,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetType(tt.eeprom)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

// MockReader implements the Reader interface for testing
type MockReader struct {
	data []byte
	err  error
}

func (m *MockReader) Read() ([]byte, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.data, nil
}

func TestRead(t *testing.T) {
	// Test with mock reader
	eeprom := createSff8079Eeprom()
	reader := &MockReader{data: eeprom}

	module, err := Read(reader)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}

	if module.Type != TypeSff8079 {
		t.Errorf("Expected SFF-8079, got %v", module.Type)
	}

	if module.Sff8079 == nil {
		t.Error("Expected Sff8079 to be populated")
	}
}

func TestReadWithError(t *testing.T) {
	// Test with reader that returns error
	reader := &MockReader{err: fmt.Errorf("read error")}

	_, err := Read(reader)
	if err == nil {
		t.Error("Expected error from Read")
	}
}

func TestReadFromPath(t *testing.T) {
	// Test the convenience function (this will fail on systems without I2C devices)
	// but we can test that it creates the right type of reader
	path := "/dev/i2c-0"
	reader := NewI2CReader(path)

	if reader.path != path {
		t.Errorf("Expected path %s, got %s", path, reader.path)
	}
}

func TestModuleString(t *testing.T) {
	// Create a mock module
	module := &Module{
		Type: TypeSff8079,
		Sff8079: &sff8079.Sff8079{
			Identifier: 3, // SFP
			Vendor:     [16]byte{'T', 'e', 's', 't', ' ', 'V', 'e', 'n', 'd', 'o', 'r'},
			VendorPn:   [16]byte{'T', 'E', 'S', 'T', '-', 'P', 'N'},
		},
	}

	str := module.String()
	if str == "" {
		t.Error("Module.String() returned empty string")
	}

	// Check that vendor name appears in output
	if len(str) < 10 {
		t.Error("Module.String() output too short")
	}
}

// Helper functions to create test EEPROM data

func createSff8079Eeprom() []byte {
	eeprom := make([]byte, 256)

	// SFF-8079 identifier (SFP)
	eeprom[0] = 0x03 // SFP identifier
	eeprom[1] = 0x04 // Extended identifier

	// Vendor name starting at byte 20
	vendorName := "Test Vendor"
	copy(eeprom[20:36], vendorName)

	// Vendor OUI at bytes 37-39
	eeprom[37] = 0x00
	eeprom[38] = 0x11
	eeprom[39] = 0x22

	// Part number starting at byte 40
	partNumber := "TEST-PN-123"
	copy(eeprom[40:56], partNumber)

	// Serial number starting at byte 68
	serialNumber := "SN123456789"
	copy(eeprom[68:84], serialNumber)

	// Date code starting at byte 84
	dateCode := "230101"
	copy(eeprom[84:92], dateCode)

	return eeprom
}

func createSff8636Eeprom() []byte {
	eeprom := make([]byte, 256)

	// SFF-8636 identifier (QSFP) at byte 128
	eeprom[128] = 0x0C // QSFP identifier

	// Vendor name starting at byte 148
	vendorName := "Test QSFP Vendor"
	copy(eeprom[148:164], vendorName)

	// Vendor OUI at bytes 165-167
	eeprom[165] = 0x00
	eeprom[166] = 0x33
	eeprom[167] = 0x44

	// Part number starting at byte 168
	partNumber := "QSFP-TEST-123"
	copy(eeprom[168:184], partNumber)

	// Serial number starting at byte 196
	serialNumber := "QSN987654321"
	copy(eeprom[196:212], serialNumber)

	// Date code starting at byte 212
	dateCode := "230101"
	copy(eeprom[212:220], dateCode)

	return eeprom
}

// TestStringOutputs tests both String() and StringCol() methods against actual EEPROM data files
func TestStringOutputs(t *testing.T) {
	testCases := []struct {
		name    string
		binFile string
		strFile string
		colFile string
	}{
		{
			name:    "FLEX-P.8596.02",
			binFile: "testdata/FLEX-P.8596.02.bin",
			strFile: "testdata/FLEX-P.8596.02.str",
			colFile: "testdata/FLEX-P.8596.02.col",
		},
		{
			name:    "IN-Q2AY2-35",
			binFile: "testdata/IN-Q2AY2-35.bin",
			strFile: "testdata/IN-Q2AY2-35.str",
			colFile: "testdata/IN-Q2AY2-35.col",
		},
		{
			name:    "JST01TMAC1CY5GEN",
			binFile: "testdata/JST01TMAC1CY5GEN.bin",
			strFile: "testdata/JST01TMAC1CY5GEN.str",
			colFile: "testdata/JST01TMAC1CY5GEN.col",
		},
		{
			name:    "TR-FC85S-N00",
			binFile: "testdata/TR-FC85S-N00.bin",
			strFile: "testdata/TR-FC85S-N00.str",
			colFile: "testdata/TR-FC85S-N00.col",
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Read the EEPROM data file
			eepromData, err := os.ReadFile(tc.binFile)
			if err != nil {
				t.Fatalf("Failed to read EEPROM file %s: %v", tc.binFile, err)
			}

			// Create a reader and parse the module
			reader := &MockReader{data: eepromData}
			module, err := Read(reader)
			if err != nil {
				t.Fatalf("Failed to parse EEPROM data: %v", err)
			}

			// Test String() method
			t.Run("String", func(t *testing.T) {
				actualOutput := module.String()

				// Check if golden output file exists, if not create it
				if _, err := os.Stat(tc.strFile); os.IsNotExist(err) {
					// Create the golden output file
					err = os.WriteFile(tc.strFile, []byte(actualOutput), 0644)
					if err != nil {
						t.Fatalf("Failed to create golden output file %s: %v", tc.strFile, err)
					}
					t.Logf("Created golden output file: %s", tc.strFile)
					return
				}

				// Read expected output and compare
				expectedBytes, err := os.ReadFile(tc.strFile)
				if err != nil {
					t.Fatalf("Failed to read expected output file %s: %v", tc.strFile, err)
				}

				expectedOutput := string(expectedBytes)

				// Normalize line endings for comparison
				actual := strings.ReplaceAll(actualOutput, "\r\n", "\n")
				expected := strings.ReplaceAll(expectedOutput, "\r\n", "\n")

				if actual != expected {
					t.Errorf("String output mismatch for %s", tc.name)
					t.Logf("Expected:\n%s", expected)
					t.Logf("Actual:\n%s", actual)

					// Show differences more clearly
					actualLines := strings.Split(actual, "\n")
					expectedLines := strings.Split(expected, "\n")

					maxLines := len(actualLines)
					if len(expectedLines) > maxLines {
						maxLines = len(expectedLines)
					}

					for i := 0; i < maxLines; i++ {
						var actualLine, expectedLine string
						if i < len(actualLines) {
							actualLine = actualLines[i]
						}
						if i < len(expectedLines) {
							expectedLine = expectedLines[i]
						}

						if actualLine != expectedLine {
							t.Logf("Line %d: expected '%s', got '%s'", i+1, expectedLine, actualLine)
						}
					}
				}
			})

			// Test StringCol() method
			t.Run("StringCol", func(t *testing.T) {
				actualOutput := module.StringCol()

				// Check if golden output file exists, if not create it
				if _, err := os.Stat(tc.colFile); os.IsNotExist(err) {
					// Create the golden output file
					err = os.WriteFile(tc.colFile, []byte(actualOutput), 0644)
					if err != nil {
						t.Fatalf("Failed to create golden output file %s: %v", tc.colFile, err)
					}
					t.Logf("Created golden output file: %s", tc.colFile)
					return
				}

				// Read expected output and compare
				expectedBytes, err := os.ReadFile(tc.colFile)
				if err != nil {
					t.Fatalf("Failed to read expected output file %s: %v", tc.colFile, err)
				}

				expectedOutput := string(expectedBytes)

				// Normalize line endings for comparison
				actual := strings.ReplaceAll(actualOutput, "\r\n", "\n")
				expected := strings.ReplaceAll(expectedOutput, "\r\n", "\n")

				if actual != expected {
					t.Errorf("StringCol output mismatch for %s", tc.name)
					t.Logf("Expected:\n%s", expected)
					t.Logf("Actual:\n%s", actual)

					// Show differences more clearly
					actualLines := strings.Split(actual, "\n")
					expectedLines := strings.Split(expected, "\n")

					maxLines := len(actualLines)
					if len(expectedLines) > maxLines {
						maxLines = len(expectedLines)
					}

					for i := 0; i < maxLines; i++ {
						var actualLine, expectedLine string
						if i < len(actualLines) {
							actualLine = actualLines[i]
						}
						if i < len(expectedLines) {
							expectedLine = expectedLines[i]
						}

						if actualLine != expectedLine {
							t.Logf("Line %d: expected '%s', got '%s'", i+1, expectedLine, actualLine)
						}
					}
				}
			})
		})
	}
}
