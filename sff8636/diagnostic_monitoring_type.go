package sff8636

import (
	"fmt"
	"strings"
)

// DiagnosticMonitoringType represents the Diagnostic Monitoring Type field from SFF-8636 specification
// Table 6-24 Diagnostic Monitoring Type (Page 00h Byte 220)
type DiagnosticMonitoringType byte

// Decode converts the raw byte to structured bit fields
func (d DiagnosticMonitoringType) Decode() DiagnosticMonitoringTypeBits {
	return DiagnosticMonitoringTypeBits{
		Reserved:                           (byte(d) & 0xC0) != 0, // Bits 7-6: Reserved
		TemperatureMonitoringImplemented:   (byte(d) & 0x20) != 0, // Bit 5: Temperature monitoring implemented
		SupplyVoltageMonitoringImplemented: (byte(d) & 0x10) != 0, // Bit 4: Supply voltage monitoring implemented
		ReceivedPowerMeasurementsType:      (byte(d) & 0x08) != 0, // Bit 3: Received power measurements type (0=OMA, 1=Average Power)
		TransmitterPowerMeasurement:        (byte(d) & 0x04) != 0, // Bit 2: Transmitter power measurement (0=Not supported, 1=Supported)
		ReservedBits:                       (byte(d) & 0x03) != 0, // Bits 1-0: Reserved
	}
}

// DiagnosticMonitoringTypeBits represents the individual bit fields of the Diagnostic Monitoring Type byte
type DiagnosticMonitoringTypeBits struct {
	Reserved                           bool // Bits 7-6: Reserved
	TemperatureMonitoringImplemented   bool // Bit 5: Temperature monitoring implemented
	SupplyVoltageMonitoringImplemented bool // Bit 4: Supply voltage monitoring implemented
	ReceivedPowerMeasurementsType      bool // Bit 3: Received power measurements type (0=OMA, 1=Average Power)
	TransmitterPowerMeasurement        bool // Bit 2: Transmitter power measurement (0=Not supported, 1=Supported)
	ReservedBits                       bool // Bits 1-0: Reserved
}

// String returns a human-readable representation of the Diagnostic Monitoring Type field
func (d DiagnosticMonitoringType) String() string {
	bits := d.Decode()
	var parts []string

	if bits.TemperatureMonitoringImplemented {
		parts = append(parts, "Temperature")
	}
	if bits.SupplyVoltageMonitoringImplemented {
		parts = append(parts, "Supply voltage")
	}
	if bits.ReceivedPowerMeasurementsType {
		parts = append(parts, "Received power measurements type: Average Power")
	} else {
		parts = append(parts, "Received power measurements type: OMA")
	}
	if bits.TransmitterPowerMeasurement {
		parts = append(parts, "Transmitter power")
	}

	if len(parts) == 0 {
		return "No diagnostic monitoring features implemented"
	}

	return strings.Join(parts, ", ")
}

// List returns a slice of all implemented diagnostic monitoring features as strings
func (d DiagnosticMonitoringType) List() []string {
	bits := d.Decode()
	var features []string

	if bits.TemperatureMonitoringImplemented {
		features = append(features, "Temperature")
	}
	if bits.SupplyVoltageMonitoringImplemented {
		features = append(features, "Supply voltage")
	}
	if bits.ReceivedPowerMeasurementsType {
		features = append(features, "Received power measurements type: Average Power")
	} else {
		features = append(features, "Received power measurements type: OMA")
	}
	if bits.TransmitterPowerMeasurement {
		features = append(features, "Transmitter power")
	}

	return features
}

// StringCol returns a colored string representation for terminal output
func (d DiagnosticMonitoringType) StringCol() string {
	bits := d.Decode()

	var result strings.Builder

	// Header
	result.WriteString(fmt.Sprintf("%s%-50s%s : %s0x%02x%s\n",
		"\x1b[36m", "Diagnostic Monitoring Type [220]", "\x1b[0m",
		"\x1b[32m", byte(d), "\x1b[0m"))

	// Individual bit details
	if bits.TemperatureMonitoringImplemented || bits.SupplyVoltageMonitoringImplemented ||
		bits.ReceivedPowerMeasurementsType || bits.TransmitterPowerMeasurement {
		result.WriteString(fmt.Sprintf("%s%-50s%s : %sDiagnostic Monitoring Details%s\n",
			"\x1b[33m", " ", "\x1b[0m", "\x1b[33m", "\x1b[0m"))

		if bits.TemperatureMonitoringImplemented {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTemperature monitoring implemented%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if bits.SupplyVoltageMonitoringImplemented {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sSupply voltage monitoring implemented%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if bits.ReceivedPowerMeasurementsType {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sReceived power measurements type: Average Power%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		} else {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sReceived power measurements type: OMA%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if bits.TransmitterPowerMeasurement {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTransmitter power measurement: Supported%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		} else {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTransmitter power measurement: Not supported%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
	}

	return result.String()
}

// IsTemperatureMonitoringImplemented returns true if temperature monitoring is implemented
func (d DiagnosticMonitoringType) IsTemperatureMonitoringImplemented() bool {
	return (byte(d) & 0x20) != 0
}

// IsSupplyVoltageMonitoringImplemented returns true if supply voltage monitoring is implemented
func (d DiagnosticMonitoringType) IsSupplyVoltageMonitoringImplemented() bool {
	return (byte(d) & 0x10) != 0
}

// IsReceivedPowerMeasurementsTypeAveragePower returns true if received power measurements type is Average Power
func (d DiagnosticMonitoringType) IsReceivedPowerMeasurementsTypeAveragePower() bool {
	return (byte(d) & 0x08) != 0
}

// IsTransmitterPowerMeasurementSupported returns true if transmitter power measurement is supported
func (d DiagnosticMonitoringType) IsTransmitterPowerMeasurementSupported() bool {
	return (byte(d) & 0x04) != 0
}
