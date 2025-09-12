package sff8636

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/bluecmd/go-sff/common"
)

const (
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[34m"
	magenta = "\x1b[35m"
	cyan    = "\x1b[36m"
	white   = "\x1b[37m"
	clear   = "\x1b[0m"
)

type Sff8636 struct {
	// Page 00h (Bytes 0-127)
	Identifier         byte                     `json:"identifier"`         // Byte 0: Identifier
	RevisionCompliance RevisionCompliance       `json:"revisionCompliance"` // Byte 1: Revision Compliance
	_                  [20]byte                 `json:"-"`                  // Bytes 2-21
	Temperature        common.TemperatureQ8_8BE `json:"temperature"`        // Bytes 22-23: Temperature
	_                  [2]byte                  `json:"-"`                  // Bytes 24-25
	SupplyVoltage      common.VoltageVoltBE     `json:"supplyVoltage"`      // Bytes 26-27: Supply Voltage
	_                  [6]byte                  `json:"-"`                  // Bytes 28-33
	ChannelMonitoring  ChannelMonitoring        `json:"channelMonitoring"`  // Bytes 34-57: Channel Monitoring
	_                  [35]byte                 `json:"-"`                  // Bytes 58-92
	ControlStatus      ControlStatus            `json:"controlStatus"`      // Byte 93: Control and Status
	_                  [34]byte                 `json:"-"`                  // Bytes 94-127

	// Page 01h (Bytes 128-255)
	IdentifierPage01  common.Identifier            `json:"identifierPage01"`  // 128 - Identifier
	ExtIdentifier     ExtIdentifier                `json:"extIdentifier"`     // 129 - Ext. Identifier
	Connector         common.Connector             `json:"connector"`         // 130 - Connector Type
	Transceiver       Transceiver                  `json:"transceiver"`       // 131-138 - Specification Compliance
	Encoding          Encoding                     `json:"encoding"`          // 139 - Encoding
	BrNominal         common.Value100Mbps          `json:"brNominal"`         // 140 - BR, nominal
	RateIdentifier    byte                         `json:"rateIdentifier"`    // 141 - Extended Rate Select Compliance
	LengthSmf         common.ValueKm               `json:"lengthSmf"`         // 142 - Length (SMF)
	LengthOm3         common.ValueM                `json:"lengthOm3"`         // 143 - Length (OM3 50 um)
	LengthOm2         common.ValueM                `json:"lengthOm2"`         // 144 - Length (OM2 50 um)
	LengthOm1         common.ValueM                `json:"lengthOm1"`         // 145 - Length (OM1 62.5 um) or Copper Cable Attenuation
	LengthCopper      common.ValueM                `json:"lengthCopper"`      // 146 - Length (passive copper or active cable or OM4 50 um)
	DevTech           DeviceTechnology             `json:"devTech"`           // 147 - Device technology
	Vendor            common.String16              `json:"vendor"`            // 148-163 - Vendor name
	ExtModule         byte                         `json:"-"`                 // 164 - Extended Module
	VendorOui         common.VendorOUI             `json:"vendorOui"`         // 165-167 - Vendor OUI
	VendorPn          common.String16              `json:"vendorPn"`          // 168-183 - Vendor PN
	VendorRev         common.String2               `json:"vendorRev"`         // 184-185 - Vendor rev
	LaserWavelen      common.WavelengthNanometerBE `json:"laserWavelen"`      // 186-187 - Wavelength or Copper Cable Attenuation
	LaserWavelenToler common.WavelengthNanometerBE `json:"laserWavelenToler"` // 188-189 - Wavelength tolerance or Copper Cable Attenuation
	MaxCaseTempC      byte                         `json:"-"`                 // 190 - Max case temp.
	CcBase            byte                         `json:"-"`                 // 191 - CC_BASE
	LinkCodes         LinkCodes                    `json:"linkCodes"`         // 192 - Link codes
	Options           Options                      `json:"options"`           // 193-195 - Options
	VendorSn          common.String16              `json:"vendorSn"`          // 196-211 - Vendor SN
	DateCode          common.DateCode              `json:"dateCode"`          // 212-219 - Date Code
	DiagMonType       byte                         `json:"-"`                 // 220 - Diagnostic Monitoring Type
	EnhOptions        byte                         `json:"-"`                 // 221 - Enhanced Options
	BrNominalExt      byte                         `json:"-"`                 // 222 - BR, Nominal
	CcExt             byte                         `json:"-"`                 // 223 - CC_EXT
	VendorSpec        [32]byte                     `json:"-"`                 // 224-255 - Vendor Specific
}

func Decode(eeprom []byte) (*Sff8636, error) {
	if len(eeprom) < 256 {
		return nil, fmt.Errorf("eeprom size to small needs to be 256 bytes or larger got: %d bytes", len(eeprom))
	}

	// Check if this is a valid SFF-8636 EEPROM by checking the identifier at byte 128
	if eeprom[128] == 12 || eeprom[128] == 13 || eeprom[128] == 17 {
		return (*Sff8636)(unsafe.Pointer(&eeprom[0])), nil
	}

	return nil, fmt.Errorf("unknown eeprom standard, identifier: 0x%02x", byte(eeprom[128]))
}

func (s *Sff8636) String() string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("%-50s : 0x%02x\n", "Identifier [0]", s.Identifier))
	result.WriteString(fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Revision Compliance [1]", byte(s.RevisionCompliance), s.RevisionCompliance))
	result.WriteString(fmt.Sprintf("%-50s :\n%s\n", "Channel Monitoring [34-81]", s.ChannelMonitoring.String()))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Temperature [22-23]", s.Temperature))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Supply Voltage [26-27]", s.SupplyVoltage))
	result.WriteString(fmt.Sprintf("%-50s : 0x%02x\n%s\n", "Control Status [93]", byte(s.ControlStatus), s.ControlStatus))
	result.WriteString(fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Identifier [128]", byte(s.IdentifierPage01), s.IdentifierPage01))
	result.WriteString(fmt.Sprintf("%-50s : 0x%02x\n", "Extended Identifier [129]", byte(s.ExtIdentifier)))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Extended Identifier Description", strings.Join(s.ExtIdentifier.List(), fmt.Sprintf("\n%-50s : ", " "))))
	result.WriteString(fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Connector [130]", byte(s.Connector), s.Connector))
	result.WriteString(fmt.Sprintf("%-50s : 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x\n", "Transceiver Codes [131-138]", s.Transceiver[0], s.Transceiver[1], s.Transceiver[2], s.Transceiver[3], s.Transceiver[4], s.Transceiver[5], s.Transceiver[6], s.Transceiver[7]))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Transceiver Type", strings.Join(s.Transceiver.List(), fmt.Sprintf("\n%-50s : ", " "))))
	result.WriteString(fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Encoding [139]", byte(s.Encoding), s.Encoding))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "BR, Nominal [140]", s.BrNominal))
	result.WriteString(fmt.Sprintf("%-50s : 0x%02x\n", "Rate Identifier [141]", s.RateIdentifier))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Length (SMF) [142]", s.LengthSmf))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Length (OM3 50um) [143]", s.LengthOm3))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Length (OM2 50um) [144]", s.LengthOm2))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Length (OM1 62.5um) [145]", s.LengthOm1))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Length (Copper or Active cable) [146]", s.LengthCopper))
	result.WriteString(fmt.Sprintf("%-50s :\n%s\n", "Device Technology [147]", s.DevTech))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Vendor [148-163]", s.Vendor))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Vendor OUI [165-167]", s.VendorOui))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Vendor PN [168-183]", s.VendorPn))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Vendor Rev [184-185]", s.VendorRev))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Wavelength [186-187]", s.LaserWavelen))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "  Tolerance [188-189]", s.LaserWavelenToler))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Option Values [193-195]", s.Options.String()))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Vendor SN [196-211]", s.VendorSn))
	result.WriteString(fmt.Sprintf("%-50s : %s\n", "Date Code [212-219]", s.DateCode))

	return result.String()
}

func strCol(k string, v string, c1 string, c2 string) string {
	return fmt.Sprintf("%s%-50s%s : %s%s%s\n", c1, k, clear, c2, v, clear)
}

func joinStrCol(k string, l []string, c1 string, c2 string) string {
	if len(l) == 0 {
		return strCol(k, "", c1, c2)
	}
	r := strCol(k, l[0], c1, c2)
	for _, s := range l[1:] {
		r += strCol("", s, c1, c2)
	}
	return r
}

func (s *Sff8636) StringCol() string {
	var result strings.Builder
	result.WriteString(strCol("Identifier [0]", fmt.Sprintf("0x%02x", s.Identifier), cyan, green))
	result.WriteString(strCol("Revision Compliance [1]", fmt.Sprintf("0x%02x (%s)", byte(s.RevisionCompliance), s.RevisionCompliance), cyan, green))
	result.WriteString(strCol("Channel Monitoring [34-81]", "", cyan, yellow))
	result.WriteString(strCol("  Rx1 Power", s.ChannelMonitoring.Rx1Power.String(), cyan, green))
	result.WriteString(strCol("  Rx2 Power", s.ChannelMonitoring.Rx2Power.String(), cyan, green))
	result.WriteString(strCol("  Rx3 Power", s.ChannelMonitoring.Rx3Power.String(), cyan, green))
	result.WriteString(strCol("  Rx4 Power", s.ChannelMonitoring.Rx4Power.String(), cyan, green))
	result.WriteString(strCol("  Tx1 Bias", s.ChannelMonitoring.Tx1Bias.String(), cyan, green))
	result.WriteString(strCol("  Tx2 Bias", s.ChannelMonitoring.Tx2Bias.String(), cyan, green))
	result.WriteString(strCol("  Tx3 Bias", s.ChannelMonitoring.Tx3Bias.String(), cyan, green))
	result.WriteString(strCol("  Tx4 Bias", s.ChannelMonitoring.Tx4Bias.String(), cyan, green))
	result.WriteString(strCol("  Tx1 Power", s.ChannelMonitoring.Tx1Power.String(), cyan, green))
	result.WriteString(strCol("  Tx2 Power", s.ChannelMonitoring.Tx2Power.String(), cyan, green))
	result.WriteString(strCol("  Tx3 Power", s.ChannelMonitoring.Tx3Power.String(), cyan, green))
	result.WriteString(strCol("  Tx4 Power", s.ChannelMonitoring.Tx4Power.String(), cyan, green))
	result.WriteString(strCol("Temperature [22-23]", s.Temperature.String(), cyan, green))
	result.WriteString(strCol("Supply Voltage [26-27]", s.SupplyVoltage.String(), cyan, green))
	result.WriteString(strCol("Control Status [93]", fmt.Sprintf("0x%02x", byte(s.ControlStatus)), cyan, green))
	result.WriteString(strCol("  Software Reset", fmt.Sprintf("%t", s.ControlStatus.IsSoftwareReset()), cyan, green))
	result.WriteString(strCol("  High Power Class 8", fmt.Sprintf("%t", s.ControlStatus.IsHighPowerClass8Enabled()), cyan, green))
	result.WriteString(strCol("  High Power Classes 5-7", fmt.Sprintf("%t", s.ControlStatus.IsHighPowerClass5to7Enabled()), cyan, green))
	result.WriteString(strCol("  Low Power Mode", fmt.Sprintf("%t", s.ControlStatus.IsLowPowerMode()), cyan, green))
	result.WriteString(strCol("Identifier [128]", fmt.Sprintf("0x%02x (%s)", byte(s.IdentifierPage01), s.IdentifierPage01), cyan, green))
	result.WriteString(strCol("Extended Identifier [129]", fmt.Sprintf("0x%02x", byte(s.ExtIdentifier)), cyan, green))
	result.WriteString(strCol("Extended Identifier Description", strings.Join(s.ExtIdentifier.List(), fmt.Sprintf("\n%-50s : ", " ")), cyan, green))
	result.WriteString(strCol("Connector [130]", fmt.Sprintf("0x%02x (%s)", byte(s.Connector), s.Connector), cyan, green))
	result.WriteString(strCol("Transceiver Codes [131-138]", fmt.Sprintf("0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x", s.Transceiver[0], s.Transceiver[1], s.Transceiver[2], s.Transceiver[3], s.Transceiver[4], s.Transceiver[5], s.Transceiver[6], s.Transceiver[7]), cyan, green))
	result.WriteString(joinStrCol("Transceiver Type", s.Transceiver.List(), cyan, yellow))
	result.WriteString(strCol("Encoding [139]", fmt.Sprintf("0x%02x (%s)", byte(s.Encoding), s.Encoding), cyan, green))
	result.WriteString(strCol("BR, Nominal [140]", s.BrNominal.String(), cyan, green))
	result.WriteString(strCol("Rate Identifier [141]", fmt.Sprintf("0x%02x", s.RateIdentifier), cyan, green))
	result.WriteString(strCol("Length (SMF) [142]", s.LengthSmf.String(), cyan, green))
	result.WriteString(strCol("Length (OM3 50um) [143]", s.LengthOm3.String(), cyan, green))
	result.WriteString(strCol("Length (OM2 50um) [144]", s.LengthOm2.String(), cyan, green))
	result.WriteString(strCol("Length (OM1 62.5um) [145]", s.LengthOm1.String(), cyan, green))
	result.WriteString(strCol("Length (Copper or Active cable) [146]", s.LengthCopper.String(), cyan, green))
	result.WriteString(strCol("Device Technology [147]", "", cyan, yellow))
	result.WriteString(strCol("  Active wavelength control (bit 3)", fmt.Sprintf("%t", s.DevTech.HasActiveWavelengthControl()), cyan, green))
	result.WriteString(strCol("  Cooled Transmitter (bit 2)", fmt.Sprintf("%t", s.DevTech.HasCooledTransmitter()), cyan, green))
	result.WriteString(strCol("  Detector Type (bit 1)", s.DevTech.GetDetectorType(), cyan, green))
	result.WriteString(strCol("  Transmitter Type (bits 7-4)", s.DevTech.GetTransmitterTechnologyName(), cyan, green))
	result.WriteString(strCol("  Tunable Transmitter (bit 0)", fmt.Sprintf("%t", s.DevTech.IsTunableTransmitter()), cyan, green))
	result.WriteString(strCol("Vendor [148-163]", s.Vendor.String(), cyan, green))
	result.WriteString(strCol("Vendor OUI [165-167]", s.VendorOui.String(), cyan, green))
	result.WriteString(strCol("Vendor PN [168-183]", s.VendorPn.String(), cyan, green))
	result.WriteString(strCol("Vendor Rev [184-185]", s.VendorRev.String(), cyan, green))
	result.WriteString(strCol("Wavelength [186-187]", s.LaserWavelen.String(), cyan, green))
	result.WriteString(strCol("  Tolerance [188-189]", s.LaserWavelenToler.String(), cyan, green))
	result.WriteString(strCol("Option Values [193-195]", s.Options.String(), cyan, green))
	result.WriteString(strCol("Vendor SN [196-211]", s.VendorSn.String(), cyan, green))
	result.WriteString(strCol("Date Code [212-219]", s.DateCode.String(), cyan, green))

	return result.String()
}
