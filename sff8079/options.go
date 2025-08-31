package sff8079

import (
	"encoding/json"
	"strings"
)

// Options represents the SFF-8079 options field (bytes 64-65)
// This is a type alias to maintain 1:1 memory mapping with EEPROM
type Options [2]byte

// OptionsByte64 represents byte 64 options according to SFF-8472 Table 8-3
type OptionsByte64 struct {
	Reserved                   bool `json:"reserved"`                   // Bit 7 - Reserved
	HighPowerLevelDeclaration4 bool `json:"highPowerLevelDeclaration4"` // Bit 6 - High Power Level Declaration (Power Level 4)
	HighPowerLevelDeclaration3 bool `json:"highPowerLevelDeclaration3"` // Bit 5 - High Power Level Declaration (Power Level 3)
	PagingImplemented          bool `json:"pagingImplemented"`          // Bit 4 - Paging implemented indicator
	RetimerOrCDR               bool `json:"retimerOrCDR"`               // Bit 3 - Retimer or CDR indicator
	CooledTransceiver          bool `json:"cooledTransceiver"`          // Bit 2 - Cooled Transceiver Declaration
	PowerLevelDeclaration      bool `json:"powerLevelDeclaration"`      // Bit 1 - Power Level Declaration (Power Level 2)
	LinearReceiverOutput       bool `json:"linearReceiverOutput"`       // Bit 0 - Linear Receiver Output Implemented
}

// OptionsByte65 represents byte 65 options according to SFF-8472 Table 8-3
type OptionsByte65 struct {
	ReceiverDecisionThreshold bool `json:"receiverDecisionThreshold"` // Bit 7 - Receiver decision threshold implemented
	TunableTransmitter        bool `json:"tunableTransmitter"`        // Bit 6 - Tunable transmitter technology
	RateSelect                bool `json:"rateSelect"`                // Bit 5 - RATE_SELECT functionality implemented
	TxDisable                 bool `json:"txDisable"`                 // Bit 4 - TX_DISABLE implemented
	TxFault                   bool `json:"txFault"`                   // Bit 3 - TX_FAULT signal implemented
	LossOfSignalInverted      bool `json:"lossOfSignalInverted"`      // Bit 2 - Loss of Signal implemented, signal inverted
	LossOfSignalStandard      bool `json:"lossOfSignalStandard"`      // Bit 1 - Loss of Signal implemented, standard behavior
	AdditionalPages           bool `json:"additionalPages"`           // Bit 0 - Additional pages indicator
}

// DecodeOptions decodes the raw options bytes into structured data
func DecodeOptions(raw [2]byte) Options {
	return Options(raw)
}

// Accessor methods for byte 64 options
func (o *Options) IsReserved() bool                   { return ((*o)[0] & 0x80) != 0 } // Bit 7
func (o *Options) IsHighPowerLevelDeclaration4() bool { return ((*o)[0] & 0x40) != 0 } // Bit 6
func (o *Options) IsHighPowerLevelDeclaration3() bool { return ((*o)[0] & 0x20) != 0 } // Bit 5
func (o *Options) IsPagingImplemented() bool          { return ((*o)[0] & 0x10) != 0 } // Bit 4
func (o *Options) IsRetimerOrCDR() bool               { return ((*o)[0] & 0x08) != 0 } // Bit 3
func (o *Options) IsCooledTransceiver() bool          { return ((*o)[0] & 0x04) != 0 } // Bit 2
func (o *Options) IsPowerLevelDeclaration() bool      { return ((*o)[0] & 0x02) != 0 } // Bit 1
func (o *Options) IsLinearReceiverOutput() bool       { return ((*o)[0] & 0x01) != 0 } // Bit 0

// Accessor methods for byte 65 options
func (o *Options) IsReceiverDecisionThreshold() bool { return ((*o)[1] & 0x80) != 0 } // Bit 7
func (o *Options) IsTunableTransmitter() bool        { return ((*o)[1] & 0x40) != 0 } // Bit 6
func (o *Options) IsRateSelect() bool                { return ((*o)[1] & 0x20) != 0 } // Bit 5
func (o *Options) IsTxDisable() bool                 { return ((*o)[1] & 0x10) != 0 } // Bit 4
func (o *Options) IsTxFault() bool                   { return ((*o)[1] & 0x08) != 0 } // Bit 3
func (o *Options) IsLossOfSignalInverted() bool      { return ((*o)[1] & 0x04) != 0 } // Bit 2
func (o *Options) IsLossOfSignalStandard() bool      { return ((*o)[1] & 0x02) != 0 } // Bit 1
func (o *Options) IsAdditionalPages() bool           { return ((*o)[1] & 0x01) != 0 } // Bit 0

// GetPowerLevel returns the power level based on the options bits
func (o *Options) GetPowerLevel() string {
	if o.IsHighPowerLevelDeclaration4() {
		return "Power Level 4"
	} else if o.IsHighPowerLevelDeclaration3() {
		return "Power Level 3"
	} else if o.IsPowerLevelDeclaration() {
		return "Power Level 2"
	} else {
		return "Power Level 1 (or unspecified)"
	}
}

// String returns a human-readable representation of the options
func (o *Options) String() string {
	var parts []string

	// Byte 64 descriptions
	if o.IsHighPowerLevelDeclaration4() {
		parts = append(parts, "High Power Level 4")
	} else if o.IsHighPowerLevelDeclaration3() {
		parts = append(parts, "High Power Level 3")
	} else if o.IsPowerLevelDeclaration() {
		parts = append(parts, "Power Level 2")
	} else {
		parts = append(parts, "Power Level 1")
	}

	if o.IsPagingImplemented() {
		parts = append(parts, "Paging Implemented")
	}
	if o.IsRetimerOrCDR() {
		parts = append(parts, "Retimer/CDR")
	}
	if o.IsCooledTransceiver() {
		parts = append(parts, "Cooled Transceiver")
	}
	if o.IsLinearReceiverOutput() {
		parts = append(parts, "Linear Receiver Output")
	}

	// Byte 65 descriptions
	if o.IsReceiverDecisionThreshold() {
		parts = append(parts, "RDT Implemented")
	}
	if o.IsTunableTransmitter() {
		parts = append(parts, "Tunable Transmitter")
	}
	if o.IsRateSelect() {
		parts = append(parts, "Rate Select")
	}
	if o.IsTxDisable() {
		parts = append(parts, "TX Disable")
	}
	if o.IsTxFault() {
		parts = append(parts, "TX Fault")
	}
	if o.IsLossOfSignalInverted() {
		parts = append(parts, "Loss of Signal (Inverted)")
	}
	if o.IsLossOfSignalStandard() {
		parts = append(parts, "Loss of Signal (Standard)")
	}
	if o.IsAdditionalPages() {
		parts = append(parts, "Additional Pages")
	}

	if len(parts) == 0 {
		return "No special options"
	}

	return strings.Join(parts, ", ")
}

// MarshalJSON implements custom JSON marshaling
func (o *Options) MarshalJSON() ([]byte, error) {
	type Alias Options
	return json.Marshal(&struct {
		*Alias
		PowerLevel string `json:"powerLevel"`
		Summary    string `json:"summary"`
	}{
		Alias:      (*Alias)(o),
		PowerLevel: o.GetPowerLevel(),
		Summary:    o.String(),
	})
}
