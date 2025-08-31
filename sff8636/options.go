package sff8636

import (
	"fmt"
	"strings"
)

// Options represents the Options field from SFF-8636 specification
// Table 6-22 Option Values (Page 00h Bytes 193-195)
type Options [3]byte

// OptionsByte0 represents the first byte (193) of the Options field
type OptionsByte0 struct {
	Reserved                  bool // Bit 7 - Reserved
	LPModeTxDisConfigurable   bool // Bit 6 - LPMode/TxDis input signal is configurable
	IntLRxLOSLConfigurable    bool // Bit 5 - IntL/RxLOSL output signal is configurable
	TxInputAdaptiveEqualizers bool // Bit 4 - Tx input adaptive equalizers freeze capable
	TxInputEqualizersAuto     bool // Bit 3 - Tx input equalizers auto-adaptive capable
	TxInputEqualizersFixed    bool // Bit 2 - Tx input equalizers fixed-programmable settings
	RxOutputEmphasisFixed     bool // Bit 1 - Rx output emphasis fixed-programmable settings
	RxOutputAmplitudeFixed    bool // Bit 0 - Rx output amplitude fixed-programmable settings
}

// OptionsByte1 represents the second byte (194) of the Options field
type OptionsByte1 struct {
	TxCDROnOffControl bool // Bit 7 - Tx CDR On/Off Control implemented
	RxCDROnOffControl bool // Bit 6 - Rx CDR On/Off Control implemented
	TxCDRLossOfLock   bool // Bit 5 - Tx CDR Loss of Lock (LOL) flag implemented
	RxCDRLossOfLock   bool // Bit 4 - Rx CDR Loss of Lock (LOL) flag implemented
	RxSquelchDisable  bool // Bit 3 - Rx Squelch Disable implemented
	RxOutputDisable   bool // Bit 2 - Rx Output Disable implemented
	TxSquelchDisable  bool // Bit 1 - Tx Squelch Disable implemented
	TxSquelch         bool // Bit 0 - Tx Squelch implemented
}

// OptionsByte2 represents the third byte (195) of the Options field
type OptionsByte2 struct {
	MemoryPage02          bool // Bit 7 - Memory Page 02 provided
	MemoryPage01          bool // Bit 6 - Memory Page 01h provided
	RateSelectImplemented bool // Bit 5 - Rate select is implemented
	TxDisableImplemented  bool // Bit 4 - Tx_Disable is implemented
	TxFaultImplemented    bool // Bit 3 - Tx_Fault signal implemented
	TxSquelchType         bool // Bit 2 - Tx Squelch type (0=OMA, 1=Pave)
	TxLossOfSignal        bool // Bit 1 - Tx Loss of Signal implemented
	Pages20_21Implemented bool // Bit 0 - Pages 20-21h implemented
}

// Decode converts the raw [3]byte Options to structured OptionsByte types
func (o Options) Decode() (OptionsByte0, OptionsByte1, OptionsByte2) {
	var byte0 OptionsByte0
	var byte1 OptionsByte1
	var byte2 OptionsByte2

	// Decode byte 0 (193)
	byte0.Reserved = (o[0] & 0x80) != 0
	byte0.LPModeTxDisConfigurable = (o[0] & 0x40) != 0
	byte0.IntLRxLOSLConfigurable = (o[0] & 0x20) != 0
	byte0.TxInputAdaptiveEqualizers = (o[0] & 0x10) != 0
	byte0.TxInputEqualizersAuto = (o[0] & 0x08) != 0
	byte0.TxInputEqualizersFixed = (o[0] & 0x04) != 0
	byte0.RxOutputEmphasisFixed = (o[0] & 0x02) != 0
	byte0.RxOutputAmplitudeFixed = (o[0] & 0x01) != 0

	// Decode byte 1 (194)
	byte1.TxCDROnOffControl = (o[1] & 0x80) != 0
	byte1.RxCDROnOffControl = (o[1] & 0x40) != 0
	byte1.TxCDRLossOfLock = (o[1] & 0x20) != 0
	byte1.RxCDRLossOfLock = (o[1] & 0x10) != 0
	byte1.RxSquelchDisable = (o[1] & 0x08) != 0
	byte1.RxOutputDisable = (o[1] & 0x04) != 0
	byte1.TxSquelchDisable = (o[1] & 0x02) != 0
	byte1.TxSquelch = (o[1] & 0x01) != 0

	// Decode byte 2 (195)
	byte2.MemoryPage02 = (o[2] & 0x80) != 0
	byte2.MemoryPage01 = (o[2] & 0x40) != 0
	byte2.RateSelectImplemented = (o[2] & 0x20) != 0
	byte2.TxDisableImplemented = (o[2] & 0x10) != 0
	byte2.TxFaultImplemented = (o[2] & 0x08) != 0
	byte2.TxSquelchType = (o[2] & 0x04) != 0
	byte2.TxLossOfSignal = (o[2] & 0x02) != 0
	byte2.Pages20_21Implemented = (o[2] & 0x01) != 0

	return byte0, byte1, byte2
}

// String returns a human-readable representation of the Options field
func (o Options) String() string {
	byte0, byte1, byte2 := o.Decode()

	var parts []string

	// Byte 0 descriptions
	if byte0.LPModeTxDisConfigurable {
		parts = append(parts, "LPMode/TxDis configurable")
	}
	if byte0.IntLRxLOSLConfigurable {
		parts = append(parts, "IntL/RxLOSL configurable")
	}
	if byte0.TxInputAdaptiveEqualizers {
		parts = append(parts, "Tx input adaptive equalizers freeze capable")
	}
	if byte0.TxInputEqualizersAuto {
		parts = append(parts, "Tx input equalizers auto-adaptive capable")
	}
	if byte0.TxInputEqualizersFixed {
		parts = append(parts, "Tx input equalizers fixed-programmable")
	}
	if byte0.RxOutputEmphasisFixed {
		parts = append(parts, "Rx output emphasis fixed-programmable")
	}
	if byte0.RxOutputAmplitudeFixed {
		parts = append(parts, "Rx output amplitude fixed-programmable")
	}

	// Byte 1 descriptions
	if byte1.TxCDROnOffControl {
		parts = append(parts, "Tx CDR On/Off Control")
	}
	if byte1.RxCDROnOffControl {
		parts = append(parts, "Rx CDR On/Off Control")
	}
	if byte1.TxCDRLossOfLock {
		parts = append(parts, "Tx CDR Loss of Lock flag")
	}
	if byte1.RxCDRLossOfLock {
		parts = append(parts, "Rx CDR Loss of Lock flag")
	}
	if byte1.RxSquelchDisable {
		parts = append(parts, "Rx Squelch Disable")
	}
	if byte1.RxOutputDisable {
		parts = append(parts, "Rx Output Disable")
	}
	if byte1.TxSquelchDisable {
		parts = append(parts, "Tx Squelch Disable")
	}
	if byte1.TxSquelch {
		parts = append(parts, "Tx Squelch")
	}

	// Byte 2 descriptions
	if byte2.MemoryPage02 {
		parts = append(parts, "Memory Page 02 provided")
	}
	if byte2.MemoryPage01 {
		parts = append(parts, "Memory Page 01h provided")
	}
	if byte2.RateSelectImplemented {
		parts = append(parts, "Rate select implemented")
	}
	if byte2.TxDisableImplemented {
		parts = append(parts, "Tx_Disable implemented")
	}
	if byte2.TxFaultImplemented {
		parts = append(parts, "Tx_Fault signal")
	}
	if byte2.TxSquelchType {
		parts = append(parts, "Tx Squelch reduces Pave")
	} else if byte1.TxSquelch {
		parts = append(parts, "Tx Squelch reduces OMA")
	}
	if byte2.TxLossOfSignal {
		parts = append(parts, "Tx Loss of Signal")
	}
	if byte2.Pages20_21Implemented {
		parts = append(parts, "Pages 20-21h implemented")
	}

	if len(parts) == 0 {
		return "No options implemented"
	}

	return strings.Join(parts, ", ")
}

// List returns a slice of all implemented options as strings
func (o Options) List() []string {
	byte0, byte1, byte2 := o.Decode()

	var options []string

	// Byte 0 options
	if byte0.LPModeTxDisConfigurable {
		options = append(options, "LPMode/TxDis configurable")
	}
	if byte0.IntLRxLOSLConfigurable {
		options = append(options, "IntL/RxLOSL configurable")
	}
	if byte0.TxInputAdaptiveEqualizers {
		options = append(options, "Tx input adaptive equalizers freeze capable")
	}
	if byte0.TxInputEqualizersAuto {
		options = append(options, "Tx input equalizers auto-adaptive capable")
	}
	if byte0.TxInputEqualizersFixed {
		options = append(options, "Tx input equalizers fixed-programmable")
	}
	if byte0.RxOutputEmphasisFixed {
		options = append(options, "Rx output emphasis fixed-programmable")
	}
	if byte0.RxOutputAmplitudeFixed {
		options = append(options, "Rx output amplitude fixed-programmable")
	}

	// Byte 1 options
	if byte1.TxCDROnOffControl {
		options = append(options, "Tx CDR On/Off Control")
	}
	if byte1.RxCDROnOffControl {
		options = append(options, "Rx CDR On/Off Control")
	}
	if byte1.TxCDRLossOfLock {
		options = append(options, "Tx CDR Loss of Lock flag")
	}
	if byte1.RxCDRLossOfLock {
		options = append(options, "Rx CDR Loss of Lock flag")
	}
	if byte1.RxSquelchDisable {
		options = append(options, "Rx Squelch Disable")
	}
	if byte1.RxOutputDisable {
		options = append(options, "Rx Output Disable")
	}
	if byte1.TxSquelchDisable {
		options = append(options, "Tx Squelch Disable")
	}
	if byte1.TxSquelch {
		options = append(options, "Tx Squelch")
	}

	// Byte 2 options
	if byte2.MemoryPage02 {
		options = append(options, "Memory Page 02 provided")
	}
	if byte2.MemoryPage01 {
		options = append(options, "Memory Page 01h provided")
	}
	if byte2.RateSelectImplemented {
		options = append(options, "Rate select implemented")
	}
	if byte2.TxDisableImplemented {
		options = append(options, "Tx_Disable implemented")
	}
	if byte2.TxFaultImplemented {
		options = append(options, "Tx_Fault signal")
	}
	if byte2.TxSquelchType {
		options = append(options, "Tx Squelch reduces Pave")
	} else if byte1.TxSquelch {
		options = append(options, "Tx Squelch reduces OMA")
	}
	if byte2.TxLossOfSignal {
		options = append(options, "Tx Loss of Signal")
	}
	if byte2.Pages20_21Implemented {
		options = append(options, "Pages 20-21h implemented")
	}

	return options
}

// StringCol returns a colored string representation for terminal output
func (o Options) StringCol() string {
	byte0, byte1, byte2 := o.Decode()

	var result strings.Builder

	// Header
	result.WriteString(fmt.Sprintf("%s%-50s%s : %s0x%02x 0x%02x 0x%02x%s\n",
		"\x1b[36m", "Options [193-195]", "\x1b[0m",
		"\x1b[32m", o[0], o[1], o[2], "\x1b[0m"))

	// Byte 0 details
	if byte0.LPModeTxDisConfigurable || byte0.IntLRxLOSLConfigurable ||
		byte0.TxInputAdaptiveEqualizers || byte0.TxInputEqualizersAuto ||
		byte0.TxInputEqualizersFixed || byte0.RxOutputEmphasisFixed ||
		byte0.RxOutputAmplitudeFixed {
		result.WriteString(fmt.Sprintf("%s%-50s%s : %sByte 0 Options%s\n",
			"\x1b[33m", " ", "\x1b[0m", "\x1b[33m", "\x1b[0m"))

		if byte0.LPModeTxDisConfigurable {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sLPMode/TxDis configurable%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte0.IntLRxLOSLConfigurable {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sIntL/RxLOSL configurable%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte0.TxInputAdaptiveEqualizers {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx input adaptive equalizers freeze capable%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte0.TxInputEqualizersAuto {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx input equalizers auto-adaptive capable%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte0.TxInputEqualizersFixed {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx input equalizers fixed-programmable%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte0.RxOutputEmphasisFixed {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sRx output emphasis fixed-programmable%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte0.RxOutputAmplitudeFixed {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sRx output amplitude fixed-programmable%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
	}

	// Byte 1 details
	if byte1.TxCDROnOffControl || byte1.RxCDROnOffControl ||
		byte1.TxCDRLossOfLock || byte1.RxCDRLossOfLock ||
		byte1.RxSquelchDisable || byte1.RxOutputDisable ||
		byte1.TxSquelchDisable || byte1.TxSquelch {
		result.WriteString(fmt.Sprintf("%s%-50s%s : %sByte 1 Options%s\n",
			"\x1b[33m", " ", "\x1b[0m", "\x1b[33m", "\x1b[0m"))

		if byte1.TxCDROnOffControl {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx CDR On/Off Control%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte1.RxCDROnOffControl {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sRx CDR On/Off Control%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte1.TxCDRLossOfLock {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx CDR Loss of Lock flag%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte1.RxCDRLossOfLock {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sRx CDR Loss of Lock flag%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte1.RxSquelchDisable {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sRx Squelch Disable%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte1.RxOutputDisable {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sRx Output Disable%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte1.TxSquelchDisable {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx Squelch Disable%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte1.TxSquelch {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx Squelch%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
	}

	// Byte 2 details
	if byte2.MemoryPage02 || byte2.MemoryPage01 ||
		byte2.RateSelectImplemented || byte2.TxDisableImplemented ||
		byte2.TxFaultImplemented || byte2.TxSquelchType ||
		byte2.TxLossOfSignal || byte2.Pages20_21Implemented {
		result.WriteString(fmt.Sprintf("%s%-50s%s : %sByte 2 Options%s\n",
			"\x1b[33m", " ", "\x1b[0m", "\x1b[33m", "\x1b[0m"))

		if byte2.MemoryPage02 {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sMemory Page 02 provided%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte2.MemoryPage01 {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sMemory Page 01h provided%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte2.RateSelectImplemented {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sRate select implemented%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte2.TxDisableImplemented {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx_Disable implemented%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte2.TxFaultImplemented {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx_Fault signal%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte2.TxSquelchType {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx Squelch reduces Pave%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		} else if byte1.TxSquelch {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx Squelch reduces OMA%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte2.TxLossOfSignal {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sTx Loss of Signal%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
		if byte2.Pages20_21Implemented {
			result.WriteString(fmt.Sprintf("%s%-50s%s : %sPages 20-21h implemented%s\n",
				"\x1b[33m", " ", "\x1b[0m", "\x1b[32m", "\x1b[0m"))
		}
	}

	return result.String()
}
