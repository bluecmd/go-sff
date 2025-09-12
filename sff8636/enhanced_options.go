package sff8636

import (
	"fmt"
	"strings"
)

// EnhancedOptions represents the Enhanced Options field from SFF-8636 specification
// Table 6-25 Enhanced Options (Page 00h Byte 221)
type EnhancedOptions byte

// Decode converts the raw byte to structured bit fields
func (e EnhancedOptions) Decode() EnhancedOptionsBits {
	return EnhancedOptionsBits{
		Reserved:                    (byte(e) & 0xE0) != 0, // Bits 7-5: Reserved
		InitCompleteFlagImplemented: (byte(e) & 0x10) != 0, // Bit 4: Initialization Complete Flag implemented
		RateSelectionImplemented:    (byte(e) & 0x08) != 0, // Bit 3: Rate Selection Declaration
		Sff8079SupportReserved:      (byte(e) & 0x04) != 0, // Bit 2: Reserved (was SFF-8079 support)
		TCReadinessFlagImplemented:  (byte(e) & 0x02) != 0, // Bit 1: TC readiness flag implemented
		SoftwareResetImplemented:    (byte(e) & 0x01) != 0, // Bit 0: Software reset is implemented
	}
}

// EnhancedOptionsBits represents the individual bit fields of the Enhanced Options byte
type EnhancedOptionsBits struct {
	Reserved                    bool // Bits 7-5: Reserved
	InitCompleteFlagImplemented bool // Bit 4: Initialization Complete Flag implemented
	RateSelectionImplemented    bool // Bit 3: Rate Selection Declaration
	Sff8079SupportReserved      bool // Bit 2: Reserved (was SFF-8079 support)
	TCReadinessFlagImplemented  bool // Bit 1: TC readiness flag implemented
	SoftwareResetImplemented    bool // Bit 0: Software reset is implemented
}

// String returns a human-readable representation of the Enhanced Options field
func (e EnhancedOptions) String() string {
	bits := e.Decode()
	var parts []string

	if bits.InitCompleteFlagImplemented {
		parts = append(parts, "Initialization Complete Flag implemented")
	}
	if bits.RateSelectionImplemented {
		parts = append(parts, "Rate Selection implemented")
	}
	if bits.TCReadinessFlagImplemented {
		parts = append(parts, "TC readiness flag implemented")
	}
	if bits.SoftwareResetImplemented {
		parts = append(parts, "Software reset implemented")
	}

	if len(parts) == 0 {
		return "No enhanced options implemented"
	}

	return strings.Join(parts, ", ")
}

// List returns a slice of all implemented enhanced options as strings
func (e EnhancedOptions) List() []string {
	bits := e.Decode()
	var options []string

	if bits.InitCompleteFlagImplemented {
		options = append(options, "Initialization Complete Flag implemented")
	}
	if bits.RateSelectionImplemented {
		options = append(options, "Rate Selection implemented")
	}
	if bits.TCReadinessFlagImplemented {
		options = append(options, "TC readiness flag implemented")
	}
	if bits.SoftwareResetImplemented {
		options = append(options, "Software reset implemented")
	}

	return options
}

// StringCol returns a colored string representation for terminal output
func (e EnhancedOptions) StringCol() string {
	bits := e.Decode()

	var result strings.Builder

	// Header
	result.WriteString(fmt.Sprintf("%s%-50s%s : %s0x%02x%s\n",
		"\x1b[36m", "Enhanced Options [221]", "\x1b[0m",
		"\x1b[32m", byte(e), "\x1b[0m"))

	// Individual bit details
	if bits.InitCompleteFlagImplemented || bits.RateSelectionImplemented ||
		bits.TCReadinessFlagImplemented || bits.SoftwareResetImplemented {
		result.WriteString(fmt.Sprintf("%s%-50s%s : %sEnhanced Options Details%s\n",
			"\x1b[33m", " ", "\x1b[0m", "\x1b[33m", "\x1b[0m"))

		if bits.InitCompleteFlagImplemented {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sInitialization Complete Flag implemented%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if bits.RateSelectionImplemented {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sRate Selection implemented%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if bits.TCReadinessFlagImplemented {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTC readiness flag implemented%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if bits.SoftwareResetImplemented {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sSoftware reset implemented%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
	}

	return result.String()
}

// IsInitCompleteFlagImplemented returns true if the Initialization Complete Flag is implemented
func (e EnhancedOptions) IsInitCompleteFlagImplemented() bool {
	return (byte(e) & 0x10) != 0
}

// IsRateSelectionImplemented returns true if Rate Selection is implemented
func (e EnhancedOptions) IsRateSelectionImplemented() bool {
	return (byte(e) & 0x08) != 0
}

// IsTCReadinessFlagImplemented returns true if TC readiness flag is implemented
func (e EnhancedOptions) IsTCReadinessFlagImplemented() bool {
	return (byte(e) & 0x02) != 0
}

// IsSoftwareResetImplemented returns true if Software reset is implemented
func (e EnhancedOptions) IsSoftwareResetImplemented() bool {
	return (byte(e) & 0x01) != 0
}
