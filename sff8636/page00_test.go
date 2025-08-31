package sff8636

import (
	"testing"
	"unsafe"
)

func TestRevisionCompliance(t *testing.T) {
	tests := []struct {
		value    RevisionCompliance
		expected string
	}{
		{RevNotSpecified, "Revision not specified. Do not use for SFF-8636 rev 2.5 or higher."},
		{RevSff8436Rev48, "SFF-8436 Rev 4.8 or earlier"},
		{RevSff8636Rev25, "SFF-8636 Rev 2.5, 2.6 and 2.7"},
		{RevisionCompliance(0xFF), "Reserved (0xff)"},
	}

	for _, test := range tests {
		result := test.value.String()
		if result != test.expected {
			t.Errorf("RevisionCompliance(0x%02x).String() = %q, want %q", byte(test.value), result, test.expected)
		}
	}
}

func TestChannelMonitoring(t *testing.T) {
	cm := &ChannelMonitoring{
		Rx1Power: [2]byte{0x00, 0x64}, // 100 in decimal
		Rx2Power: [2]byte{0x00, 0xC8}, // 200 in decimal
		Tx1Bias:  [2]byte{0x00, 0x32}, // 50 in decimal
		Tx1Power: [2]byte{0x00, 0x96}, // 150 in decimal
	}

	// Test Rx power calculations
	if power := cm.GetRxPower(1); power != 0.0100 { // 100 * 0.0001
		t.Errorf("GetRxPower(1) = %f, want 0.0100", power)
	}
	if power := cm.GetRxPower(2); power != 0.0200 { // 200 * 0.0001
		t.Errorf("GetRxPower(2) = %f, want 0.0200", power)
	}

	// Test Tx bias calculations
	if bias := cm.GetTxBias(1); bias != 0.1000 { // 50 * 0.002
		t.Errorf("GetTxBias(1) = %f, want 0.1000", bias)
	}

	// Test Tx power calculations
	if power := cm.GetTxPower(1); power < 0.0149 || power > 0.0151 { // 150 * 0.0001 with tolerance
		t.Errorf("GetTxPower(1) = %f, want 0.0150 ± 0.0001", power)
	}

	// Test invalid channel
	if power := cm.GetRxPower(0); power != 0.0 {
		t.Errorf("GetRxPower(0) = %f, want 0.0", power)
	}
	if power := cm.GetRxPower(5); power != 0.0 {
		t.Errorf("GetRxPower(5) = %f, want 0.0", power)
	}
}

func TestControlStatus(t *testing.T) {
	// Test all bits set
	cs := ControlStatus(0xFF)
	if !cs.IsSoftwareReset() {
		t.Error("IsSoftwareReset() should be true when bit 7 is set")
	}
	if !cs.IsHighPowerClass8Enabled() {
		t.Error("IsHighPowerClass8Enabled() should be true when bit 5 is set")
	}
	if !cs.IsHighPowerClass5to7Enabled() {
		t.Error("IsHighPowerClass5to7Enabled() should be true when bit 4 is set")
	}
	if !cs.IsLowPowerMode() {
		t.Error("IsLowPowerMode() should be true when bit 1 is set")
	}

	// Test no bits set
	cs = ControlStatus(0x00)
	if cs.IsSoftwareReset() {
		t.Error("IsSoftwareReset() should be false when bit 7 is not set")
	}
	if cs.IsHighPowerClass8Enabled() {
		t.Error("IsHighPowerClass8Enabled() should be false when bit 5 is not set")
	}
	if cs.IsHighPowerClass5to7Enabled() {
		t.Error("IsHighPowerClass5to7Enabled() should be false when bit 4 is not set")
	}
	if cs.IsLowPowerMode() {
		t.Error("IsLowPowerMode() should be false when bit 1 is not set")
	}
}

func TestSff8636WithPage00(t *testing.T) {
	// Create sample EEPROM data
	eeprom := make([]byte, 256)

	// Set identifier (byte 0)
	eeprom[0] = 0x0C // QSFP+

	// Set revision compliance (byte 1)
	eeprom[1] = 0x07 // SFF-8636 Rev 2.5, 2.6 and 2.7

	// Set some sample channel monitoring values (bytes 34-81)
	// Rx1 Power: 1.0 mW (10000 * 0.0001)
	eeprom[34] = 0x27 // MSB
	eeprom[35] = 0x10 // LSB (10000 in decimal)

	// Set control status (byte 93)
	eeprom[93] = 0xA2 // Software reset + High Power Class 8 + Low Power Mode

	// Set the standard SFF-8636 identifier (byte 128)
	eeprom[128] = 0x0C // QSFP+

	// Decode the EEPROM
	sff, err := Decode(eeprom)
	if err != nil {
		t.Fatalf("Error decoding EEPROM: %v", err)
	}

	// Test that Page00 fields are properly decoded
	if sff.Identifier != 0x0C {
		t.Errorf("Expected identifier 0x0C, got 0x%02x", sff.Identifier)
	}

	if sff.RevisionCompliance != RevSff8636Rev25 {
		t.Errorf("Expected revision compliance %v, got %v", RevSff8636Rev25, sff.RevisionCompliance)
	}

	if sff.IdentifierPage01 != 0x0C {
		t.Errorf("Expected page01 identifier 0x0C, got 0x%02x", sff.IdentifierPage01)
	}
}

func TestSff8636StructSize(t *testing.T) {
	var sff Sff8636
	size := unsafe.Sizeof(sff)
	t.Logf("Sff8636 struct size: %d bytes", size)

	// Check individual field offsets
	t.Logf("Identifier offset: %d", unsafe.Offsetof(sff.Identifier))
	t.Logf("RevisionCompliance offset: %d", unsafe.Offsetof(sff.RevisionCompliance))
	t.Logf("ChannelMonitoring offset: %d", unsafe.Offsetof(sff.ChannelMonitoring))
	t.Logf("ControlStatus offset: %d", unsafe.Offsetof(sff.ControlStatus))
	t.Logf("IdentifierPage01 offset: %d", unsafe.Offsetof(sff.IdentifierPage01))
}
