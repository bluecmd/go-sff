package sff8636

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// DeviceTechnology represents the device technology byte (Byte 147)
type DeviceTechnology byte

// Transmitter technology constants (bits 7-4)
const (
	TxTech850nmVCSEL               DeviceTechnology = 0x00 // 0000b: 850 nm VCSEL
	TxTech1310nmVCSEL              DeviceTechnology = 0x10 // 0001b: 1310 nm VCSEL
	TxTech1550nmVCSEL              DeviceTechnology = 0x20 // 0010b: 1550 nm VCSEL
	TxTech1310nmFP                 DeviceTechnology = 0x30 // 0011b: 1310 nm FP
	TxTech1310nmDFB                DeviceTechnology = 0x40 // 0100b: 1310 nm DFB
	TxTech1550nmDFB                DeviceTechnology = 0x50 // 0101b: 1550 nm DFB
	TxTech1310nmEML                DeviceTechnology = 0x60 // 0110b: 1310 nm EML
	TxTech1550nmEML                DeviceTechnology = 0x70 // 0111b: 1550 nm EML
	TxTechOtherUndefined           DeviceTechnology = 0x80 // 1000b: Other / Undefined
	TxTech1490nmDFB                DeviceTechnology = 0x90 // 1001b: 1490 nm DFB
	TxTechCopperUnequalized        DeviceTechnology = 0xA0 // 1010b: Copper cable unequalized
	TxTechCopperPassiveEqualized   DeviceTechnology = 0xB0 // 1011b: Copper cable passive equalized
	TxTechCopperNearFarEndLimiting DeviceTechnology = 0xC0 // 1100b: Copper cable, near and far end limiting active equalizers
	TxTechCopperFarEndLimiting     DeviceTechnology = 0xD0 // 1101b: Copper cable, far end limiting active equalizers
	TxTechCopperNearEndLimiting    DeviceTechnology = 0xE0 // 1110b: Copper cable, near end limiting active equalizers
	TxTechCopperLinearActive       DeviceTechnology = 0xF0 // 1111b: Copper cable, linear active equalizers
)

// Bit masks for individual fields
const (
	TxTechMask         DeviceTechnology = 0xF0 // Bits 7-4: Transmitter technology
	WavelengthCtrlMask DeviceTechnology = 0x08 // Bit 3: Wavelength control
	CooledTxMask       DeviceTechnology = 0x04 // Bit 2: Cooled transmitter
	DetectorMask       DeviceTechnology = 0x02 // Bit 1: Detector type
	TunableTxMask      DeviceTechnology = 0x01 // Bit 0: Tunable transmitter
)

var transmitterTechNames = map[DeviceTechnology]string{
	TxTech850nmVCSEL:               "850 nm VCSEL",
	TxTech1310nmVCSEL:              "1310 nm VCSEL",
	TxTech1550nmVCSEL:              "1550 nm VCSEL",
	TxTech1310nmFP:                 "1310 nm FP",
	TxTech1310nmDFB:                "1310 nm DFB",
	TxTech1550nmDFB:                "1550 nm DFB",
	TxTech1310nmEML:                "1310 nm EML",
	TxTech1550nmEML:                "1550 nm EML",
	TxTechOtherUndefined:           "Other / Undefined",
	TxTech1490nmDFB:                "1490 nm DFB",
	TxTechCopperUnequalized:        "Copper cable unequalized",
	TxTechCopperPassiveEqualized:   "Copper cable passive equalized",
	TxTechCopperNearFarEndLimiting: "Copper cable, near and far end limiting active equalizers",
	TxTechCopperFarEndLimiting:     "Copper cable, far end limiting active equalizers",
	TxTechCopperNearEndLimiting:    "Copper cable, near end limiting active equalizers",
	TxTechCopperLinearActive:       "Copper cable, linear active equalizers",
}

// GetTransmitterTechnology returns the transmitter technology (bits 7-4)
func (d DeviceTechnology) GetTransmitterTechnology() DeviceTechnology {
	return d & TxTechMask
}

// HasActiveWavelengthControl returns true if bit 3 is set (active wavelength control)
func (d DeviceTechnology) HasActiveWavelengthControl() bool {
	return d&WavelengthCtrlMask == WavelengthCtrlMask
}

// HasCooledTransmitter returns true if bit 2 is set (cooled transmitter)
func (d DeviceTechnology) HasCooledTransmitter() bool {
	return d&CooledTxMask == CooledTxMask
}

// GetDetectorType returns the detector type (bit 1)
func (d DeviceTechnology) GetDetectorType() string {
	if d&DetectorMask == DetectorMask {
		return "APD"
	}
	return "Pin"
}

// IsTunableTransmitter returns true if bit 0 is set (tunable transmitter)
func (d DeviceTechnology) IsTunableTransmitter() bool {
	return d&TunableTxMask == TunableTxMask
}

// GetTransmitterTechnologyName returns the human-readable name for the transmitter technology
func (d DeviceTechnology) GetTransmitterTechnologyName() string {
	tech := d.GetTransmitterTechnology()
	if name, ok := transmitterTechNames[tech]; ok {
		return name
	}
	return fmt.Sprintf("Unknown (0x%02x)", byte(tech))
}

// String returns a formatted string representation of the device technology
func (d DeviceTechnology) String() string {
	return fmt.Sprintf("Device Tech (00h:147):\n    Active wavelength control (bit 3): %t\n    Cooled Transmitter (bit 2): %t\n    Detector Type (bit 1): %s\n    Transmitter Type (bits 7-4): %s\n    Tunable Transmitter (bit 0): %t",
		d.HasActiveWavelengthControl(),
		d.HasCooledTransmitter(),
		d.GetDetectorType(),
		d.GetTransmitterTechnologyName(),
		d.IsTunableTransmitter())
}

// MarshalJSON implements json.Marshaler interface
func (d DeviceTechnology) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"activeWavelengthControl": d.HasActiveWavelengthControl(),
		"cooledTransmitter":       d.HasCooledTransmitter(),
		"detectorType":            d.GetDetectorType(),
		"transmitterType":         d.GetTransmitterTechnologyName(),
		"tunableTransmitter":      d.IsTunableTransmitter(),
		"hex":                     hex.EncodeToString([]byte{byte(d)}),
	}
	return json.Marshal(m)
}
