package sff8636

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bluecmd/go-sff/common"
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
	Rx1Power common.PowerMilliWattBE  `json:"rx1Power"` // Bytes 34-35: Rx1 Power MSB/LSB
	Rx2Power common.PowerMilliWattBE  `json:"rx2Power"` // Bytes 36-37: Rx2 Power MSB/LSB
	Rx3Power common.PowerMilliWattBE  `json:"rx3Power"` // Bytes 38-39: Rx3 Power MSB/LSB
	Rx4Power common.PowerMilliWattBE  `json:"rx4Power"` // Bytes 40-41: Rx4 Power MSB/LSB
	Tx1Bias  common.CurrentMilliAmpBE `json:"tx1Bias"`  // Bytes 42-43: Tx1 Bias MSB/LSB
	Tx2Bias  common.CurrentMilliAmpBE `json:"tx2Bias"`  // Bytes 44-45: Tx2 Bias MSB/LSB
	Tx3Bias  common.CurrentMilliAmpBE `json:"tx3Bias"`  // Bytes 46-47: Tx3 Bias MSB/LSB
	Tx4Bias  common.CurrentMilliAmpBE `json:"tx4Bias"`  // Bytes 48-49: Tx4 Bias MSB/LSB
	Tx1Power common.PowerMilliWattBE  `json:"tx1Power"` // Bytes 50-51: Tx1 Power MSB/LSB
	Tx2Power common.PowerMilliWattBE  `json:"tx2Power"` // Bytes 52-53: Tx2 Power MSB/LSB
	Tx3Power common.PowerMilliWattBE  `json:"tx3Power"` // Bytes 54-55: Tx3 Power MSB/LSB
	Tx4Power common.PowerMilliWattBE  `json:"tx4Power"` // Bytes 56-57: Tx4 Power MSB/LSB
}

func (c *ChannelMonitoring) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    Rx1 Power:      %s\n", c.Rx1Power.String()))
	sb.WriteString(fmt.Sprintf("    Rx2 Power:      %s\n", c.Rx2Power.String()))
	sb.WriteString(fmt.Sprintf("    Rx3 Power:      %s\n", c.Rx3Power.String()))
	sb.WriteString(fmt.Sprintf("    Rx4 Power:      %s\n", c.Rx4Power.String()))
	sb.WriteString(fmt.Sprintf("    Tx1 Bias:       %s\n", c.Tx1Bias.String()))
	sb.WriteString(fmt.Sprintf("    Tx2 Bias:       %s\n", c.Tx2Bias.String()))
	sb.WriteString(fmt.Sprintf("    Tx3 Bias:       %s\n", c.Tx3Bias.String()))
	sb.WriteString(fmt.Sprintf("    Tx4 Bias:       %s\n", c.Tx4Bias.String()))
	sb.WriteString(fmt.Sprintf("    Tx1 Power:      %s\n", c.Tx1Power.String()))
	sb.WriteString(fmt.Sprintf("    Tx2 Power:      %s\n", c.Tx2Power.String()))
	sb.WriteString(fmt.Sprintf("    Tx3 Power:      %s\n", c.Tx3Power.String()))
	sb.WriteString(fmt.Sprintf("    Tx4 Power:      %s", c.Tx4Power.String()))
	return sb.String()
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
