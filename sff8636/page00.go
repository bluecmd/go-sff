package sff8636

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

// RevisionCompliance represents the Memory Map Version (Byte 1)
type RevisionCompliance byte

const (
	RevNotSpecified     RevisionCompliance = 0x00
	RevSff8436Rev48     RevisionCompliance = 0x01
	RevSff8436Rev48Plus RevisionCompliance = 0x02
	RevSff8636Rev13     RevisionCompliance = 0x03
	RevSff8636Rev14     RevisionCompliance = 0x04
	RevSff8636Rev15     RevisionCompliance = 0x05
	RevSff8636Rev20     RevisionCompliance = 0x06
	RevSff8636Rev25     RevisionCompliance = 0x07
	RevSff8636Rev28     RevisionCompliance = 0x08
)

var revisionNames = map[RevisionCompliance]string{
	RevNotSpecified:     "Revision not specified. Do not use for SFF-8636 rev 2.5 or higher.",
	RevSff8436Rev48:     "SFF-8436 Rev 4.8 or earlier",
	RevSff8436Rev48Plus: "Includes functionality described in revision 4.8 or earlier of SFF-8436, except that this byte and Bytes 186-189 are as defined in this document",
	RevSff8636Rev13:     "SFF-8636 Rev 1.3 or earlier",
	RevSff8636Rev14:     "SFF-8636 Rev 1.4",
	RevSff8636Rev15:     "SFF-8636 Rev 1.5",
	RevSff8636Rev20:     "SFF-8636 Rev 2.0",
	RevSff8636Rev25:     "SFF-8636 Rev 2.5, 2.6 and 2.7",
	RevSff8636Rev28:     "SFF-8636 Rev 2.8, 2.9 and 2.10",
}

func (r RevisionCompliance) String() string {
	if name, ok := revisionNames[r]; ok {
		return name
	}
	return fmt.Sprintf("Reserved (0x%02x)", byte(r))
}

func (r RevisionCompliance) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": r.String(),
		"hex":   hex.EncodeToString([]byte{byte(r)}),
	}
	return json.Marshal(m)
}

// ChannelMonitoring represents the channel monitoring values (Bytes 34-57)
type ChannelMonitoring struct {
	Rx1Power [2]byte `json:"rx1Power"` // Bytes 34-35: Rx1 Power MSB/LSB
	Rx2Power [2]byte `json:"rx2Power"` // Bytes 36-37: Rx2 Power MSB/LSB
	Rx3Power [2]byte `json:"rx3Power"` // Bytes 38-39: Rx3 Power MSB/LSB
	Rx4Power [2]byte `json:"rx4Power"` // Bytes 40-41: Rx4 Power MSB/LSB
	Tx1Bias  [2]byte `json:"tx1Bias"`  // Bytes 42-43: Tx1 Bias MSB/LSB
	Tx2Bias  [2]byte `json:"tx2Bias"`  // Bytes 44-45: Tx2 Bias MSB/LSB
	Tx3Bias  [2]byte `json:"tx3Bias"`  // Bytes 46-47: Tx3 Bias MSB/LSB
	Tx4Bias  [2]byte `json:"tx4Bias"`  // Bytes 48-49: Tx4 Bias MSB/LSB
	Tx1Power [2]byte `json:"tx1Power"` // Bytes 50-51: Tx1 Power MSB/LSB
	Tx2Power [2]byte `json:"tx2Power"` // Bytes 52-53: Tx2 Power MSB/LSB
	Tx3Power [2]byte `json:"tx3Power"` // Bytes 54-55: Tx3 Power MSB/LSB
	Tx4Power [2]byte `json:"tx4Power"` // Bytes 56-57: Tx4 Power MSB/LSB
}

// GetRxPower returns the Rx power in mW for the specified channel (1-4)
func (c *ChannelMonitoring) GetRxPower(channel int) float64 {
	if channel < 1 || channel > 4 {
		return 0.0
	}

	var power [2]byte
	switch channel {
	case 1:
		power = c.Rx1Power
	case 2:
		power = c.Rx2Power
	case 3:
		power = c.Rx3Power
	case 4:
		power = c.Rx4Power
	}

	// Convert 16-bit value to mW (assuming linear scale)
	// This is a simplified conversion - actual scaling may vary by vendor
	value := uint16(power[0])<<8 | uint16(power[1])
	return float64(value) * 0.0001 // Scale factor may need adjustment
}

// GetTxBias returns the Tx bias current in mA for the specified channel (1-4)
func (c *ChannelMonitoring) GetTxBias(channel int) float64 {
	if channel < 1 || channel > 4 {
		return 0.0
	}

	var bias [2]byte
	switch channel {
	case 1:
		bias = c.Tx1Bias
	case 2:
		bias = c.Tx2Bias
	case 3:
		bias = c.Tx3Bias
	case 4:
		bias = c.Tx4Bias
	}

	// Convert 16-bit value to mA (assuming linear scale)
	// This is a simplified conversion - actual scaling may vary by vendor
	value := uint16(bias[0])<<8 | uint16(bias[1])
	return float64(value) * 0.002 // Scale factor may need adjustment
}

// GetTxPower returns the Tx power in mW for the specified channel (1-4)
func (c *ChannelMonitoring) GetTxPower(channel int) float64 {
	if channel < 1 || channel > 4 {
		return 0.0
	}

	var power [2]byte
	switch channel {
	case 1:
		power = c.Tx1Power
	case 2:
		power = c.Tx2Power
	case 3:
		power = c.Tx3Power
	case 4:
		power = c.Tx4Power
	}

	// Convert 16-bit value to mW (assuming linear scale)
	// This is a simplified conversion - actual scaling may vary by vendor
	value := uint16(power[0])<<8 | uint16(power[1])
	return float64(value) * 0.0001 // Scale factor may need adjustment
}

func (c *ChannelMonitoring) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    Rx1 Power:      %.4f mW\n", c.GetRxPower(1)))
	sb.WriteString(fmt.Sprintf("    Rx2 Power:      %.4f mW\n", c.GetRxPower(2)))
	sb.WriteString(fmt.Sprintf("    Rx3 Power:      %.4f mW\n", c.GetRxPower(3)))
	sb.WriteString(fmt.Sprintf("    Rx4 Power:      %.4f mW\n", c.GetRxPower(4)))
	sb.WriteString(fmt.Sprintf("    Tx1 Bias:       %.4f mA\n", c.GetTxBias(1)))
	sb.WriteString(fmt.Sprintf("    Tx2 Bias:       %.4f mA\n", c.GetTxBias(2)))
	sb.WriteString(fmt.Sprintf("    Tx3 Bias:       %.4f mA\n", c.GetTxBias(3)))
	sb.WriteString(fmt.Sprintf("    Tx4 Bias:       %.4f mA\n", c.GetTxBias(4)))
	sb.WriteString(fmt.Sprintf("    Tx1 Power:      %.4f mW\n", c.GetTxPower(1)))
	sb.WriteString(fmt.Sprintf("    Tx2 Power:      %.4f mW\n", c.GetTxPower(2)))
	sb.WriteString(fmt.Sprintf("    Tx3 Power:      %.4f mW\n", c.GetTxPower(3)))
	sb.WriteString(fmt.Sprintf("    Tx4 Power:      %.4f mW", c.GetTxPower(4)))
	return sb.String()
}

func (c *ChannelMonitoring) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"rx1Power": map[string]interface{}{
			"value": c.GetRxPower(1),
			"hex":   hex.EncodeToString(c.Rx1Power[:]),
		},
		"rx2Power": map[string]interface{}{
			"value": c.GetRxPower(2),
			"hex":   hex.EncodeToString(c.Rx2Power[:]),
		},
		"rx3Power": map[string]interface{}{
			"value": c.GetRxPower(3),
			"hex":   hex.EncodeToString(c.Rx3Power[:]),
		},
		"rx4Power": map[string]interface{}{
			"value": c.GetRxPower(4),
			"hex":   hex.EncodeToString(c.Rx4Power[:]),
		},
		"tx1Bias": map[string]interface{}{
			"value": c.GetTxBias(1),
			"hex":   hex.EncodeToString(c.Tx1Bias[:]),
		},
		"tx2Bias": map[string]interface{}{
			"value": c.GetTxBias(2),
			"hex":   hex.EncodeToString(c.Tx2Bias[:]),
		},
		"tx3Bias": map[string]interface{}{
			"value": c.GetTxBias(3),
			"hex":   hex.EncodeToString(c.Tx3Bias[:]),
		},
		"tx4Bias": map[string]interface{}{
			"value": c.GetTxBias(4),
			"hex":   hex.EncodeToString(c.Tx4Bias[:]),
		},
		"tx1Power": map[string]interface{}{
			"value": c.GetTxPower(1),
			"hex":   hex.EncodeToString(c.Tx1Power[:]),
		},
		"tx2Power": map[string]interface{}{
			"value": c.GetTxPower(2),
			"hex":   hex.EncodeToString(c.Tx2Power[:]),
		},
		"tx3Power": map[string]interface{}{
			"value": c.GetTxPower(3),
			"hex":   hex.EncodeToString(c.Tx3Power[:]),
		},
		"tx4Power": map[string]interface{}{
			"value": c.GetTxPower(4),
			"hex":   hex.EncodeToString(c.Tx4Power[:]),
		},
	}
	return json.Marshal(m)
}

// ControlStatus represents the control and status byte (Byte 93)
type ControlStatus byte

const (
	// Bit 7: Software Reset
	SwResetMask ControlStatus = 0x80
	SwReset     ControlStatus = 0x80

	// Bit 6-4: Reserved

	// Bit 3: High Power Class Enable (Class 8)
	HighPowerClass8Mask ControlStatus = 0x08
	HighPowerClass8     ControlStatus = 0x08

	// Bit 2: High Power Class Enable (Classes 5-7)
	HighPowerClass5to7Mask ControlStatus = 0x04
	HighPowerClass5to7     ControlStatus = 0x04

	// Bit 1: Power set (Low Power Mode)
	PowerSetMask ControlStatus = 0x02
	PowerSet     ControlStatus = 0x02

	// Bit 0: Power override
	PowerOverrideMask ControlStatus = 0x01
	PowerOverride     ControlStatus = 0x01
)

func (c ControlStatus) IsSoftwareReset() bool {
	return c&SwResetMask == SwReset
}

func (c ControlStatus) IsHighPowerClass8Enabled() bool {
	return c&HighPowerClass8Mask == HighPowerClass8
}

func (c ControlStatus) IsHighPowerClass5to7Enabled() bool {
	return c&HighPowerClass5to7Mask == HighPowerClass5to7
}

func (c ControlStatus) IsLowPowerMode() bool {
	return c&PowerSetMask == PowerSet
}

func (c ControlStatus) IsPowerOverride() bool {
	return c&PowerOverrideMask == PowerOverride
}

func (c ControlStatus) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    Software Reset:                 %t\n", c.IsSoftwareReset()))
	sb.WriteString(fmt.Sprintf("    High Power Class 8 Enabled:     %t\n", c.IsHighPowerClass8Enabled()))
	sb.WriteString(fmt.Sprintf("    High Power Classes 5-7 Enabled: %t\n", c.IsHighPowerClass5to7Enabled()))
	sb.WriteString(fmt.Sprintf("    Low Power Mode:                 %t\n", c.IsLowPowerMode()))
	sb.WriteString(fmt.Sprintf("    Power Override:                 %t", c.IsPowerOverride()))
	return sb.String()
}

func (c ControlStatus) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"softwareReset":             c.IsSoftwareReset(),
		"highPowerClass8Enabled":    c.IsHighPowerClass8Enabled(),
		"highPowerClass5to7Enabled": c.IsHighPowerClass5to7Enabled(),
		"lowPowerMode":              c.IsLowPowerMode(),
		"powerOverride":             c.IsPowerOverride(),
		"hex":                       hex.EncodeToString([]byte{byte(c)}),
	}
	return json.Marshal(m)
}
