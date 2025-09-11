package sff8079

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

type Sff8079 struct {
	Identifier      common.Identifier   `json:"identifier"`     // 0 - Identifier
	ExtIdentifier   ExtIdentifier       `json:"extIdentifier"`  // 1 - Ext. Identifier
	Connector       common.Connector    `json:"connector"`      // 2 - Connector
	Transceiver     Transceiver         `json:"transceiver"`    // 3-10 - Transceiver
	Encoding        Encoding            `json:"encoding"`       // 11 - Encoding
	BrNominal       common.Value100Mbps `json:"brNominal"`      // 12 - BR Nominal
	RateIdentifier  byte                `json:"rateIdentifier"` // 13 - Rate ID
	LengthSmfKm     common.ValueKm      `json:"lengthSmfKm"`    // 14 - Length(9μm) - km - (SMF)?
	LengthSmfM      common.ValueM       `json:"lengthSmfM"`     // 15 - Length (9μm) - (SMF)?
	Length50umM     common.ValueM       `json:"length50umM"`    // 16 - Length (50μm)
	Length625umM    common.ValueM       `json:"length625umM"`   // 17 - Length (62.5um)
	LengthCopper    common.ValueM       `json:"lengthCopper"`   // 18 - Length (Copper)
	LengthOm3       common.ValueM       `json:"lengthOm3"`      // 19 - Length (50μm)
	Vendor          common.String16     `json:"vendor"`         // 20-35 - Vendor name
	TranscComp      byte                `json:"-"`              // 36 - Transciever
	VendorOui       common.VendorOUI    `json:"vendorOui"`      // 37-39 - Vendor OUI
	VendorPn        common.String16     `json:"vendorPn"`       // 40-55 - Vendor PN
	VendorRev       common.String4      `json:"vendorRev"`      // 56-59 - Vendor rev
	LaserWavelength [2]byte             `json:"-"`              // 60-61 - Laser wavelength
	Unallocated     byte                `json:"-"`              // 62 - Unallocated
	CcBase          byte                `json:"-"`              // 63 - CC_BASE
	Options         Options             `json:"options"`        // 64-65 - Options
	BrMax           common.ValuePerc    `json:"brMax"`          // 66 - BR, max
	BrMin           common.ValuePerc    `json:"brMin"`          // 67 - BR, min
	VendorSn        common.String16     `json:"vendorSn"`       // 68-83 - Vendor SN
	DateCode        common.DateCode     `json:"dateCode"`       // 84-91 - Date code
	DiagMonitType   byte                `json:"-"`              // 92 - Diagnostic Monitoring Type
	EnhancedOpts    byte                `json:"-"`              // 93 - Enhanced Options
	Sff8472Comp     byte                `json:"-"`              // 94 - SFF-8472 Compliance
	CcExt           byte                `json:"-"`              // 95 - CC_EXT
	VendorSpec1     [24]byte            `json:"-"`              // 96-119 - Vendor Specific 1
	VendorAristaSa  byte                `json:"vendorSa"`       // 120 Vendor Arista SA
	VendorSpec2     [7]byte             `json:"-"`              // 121-127 - Vendor Specific 2
	Reserved        [128]byte           `json:"-"`              // 128-255 - Reserved
	// Address A2h
	A2hReserved0 [96]byte                 `json:"-"`           // 0-95 - Reserved
	Temperature  common.TemperatureQ8_8BE `json:"temperature"` // 96-97 - Internally measured module temperature
	Vcc          common.VoltageVoltBE     `json:"vcc"`         // 98-99 - Internally measured supply voltage in transceiver
	TxBias       common.CurrentMilliAmpBE `json:"txBias"`      // 100-101 - Internally measured TX Bias Current
	TxPower      common.PowerMilliWattBE  `json:"txPower"`     // 102-103 - Measured TX output power
	RxPower      common.PowerMilliWattBE  `json:"rxPower"`     // 104-105 - Measured RX input power
	A2hReserved1 [22]byte                 `json:"-"`           // 106-127 - Reserved
}

func Decode(eeprom []byte) (*Sff8079, error) {
	if len(eeprom) < 512 {
		return nil, fmt.Errorf("eeprom size to small needs to be 512 bytes or larger got: %d bytes", len(eeprom))
	}

	if (eeprom[0] == 2 || eeprom[0] == 3 || eeprom[0] == 0xb) && eeprom[1] == 4 {
		sff := (*Sff8079)(unsafe.Pointer(&eeprom[0]))
		return sff, nil
	}

	return nil, fmt.Errorf("unknown eeprom standard, identifier: 0x%02x", byte(eeprom[0]))
}

func (s *Sff8079) String() string {
	str := fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Identifier [0]", byte(s.Identifier), s.Identifier) +
		fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Extended Identifier [1]", byte(s.ExtIdentifier), s.ExtIdentifier) +
		fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Connector [2]", byte(s.Connector), s.Connector) +
		fmt.Sprintf("%-50s : 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x\n", "Transceiver Codes [3-10]", s.Transceiver[0], s.Transceiver[1], s.Transceiver[2], s.Transceiver[3], s.Transceiver[4], s.Transceiver[5], s.Transceiver[6], s.Transceiver[7]) +
		fmt.Sprintf("%-50s : %s\n", "Transceiver Type", strings.Join(s.Transceiver.List(), fmt.Sprintf("\n%-50s : ", " "))) +
		fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Encoding [11]", byte(s.Encoding), s.Encoding) +
		fmt.Sprintf("%-50s : %s\n", "BR, Nominal [12]", s.BrNominal) +
		fmt.Sprintf("%-50s : 0x%02x\n", "Rate Identifier [13]", s.RateIdentifier) +
		fmt.Sprintf("%-50s : %s\n", "Length (SMF) [14]", s.LengthSmfKm.String()) +
		fmt.Sprintf("%-50s : %s\n", "Length (SMF) [15]", s.LengthSmfM) +
		fmt.Sprintf("%-50s : %s\n", "Length (50um) [16]", s.Length50umM) +
		fmt.Sprintf("%-50s : %s\n", "Length (62.5um) [17]", s.Length625umM) +
		fmt.Sprintf("%-50s : %s\n", "Length (Copper) [18]", s.LengthCopper) +
		fmt.Sprintf("%-50s : %s\n", "Length (OM3) [19]", s.LengthOm3) +
		fmt.Sprintf("%-50s : %s\n", "Vendor [20-35]", s.Vendor) +
		fmt.Sprintf("%-50s : %s\n", "Vendor OUI [37-39]", s.VendorOui) +
		fmt.Sprintf("%-50s : %s\n", "Vendor PN [40-55]", s.VendorPn) +
		fmt.Sprintf("%-50s : %s\n", "Vendor Rev [56-59]", s.VendorRev) +
		fmt.Sprintf("%-50s : %s\n", "Option Values [64-65]", s.Options.String()) +
		fmt.Sprintf("%-50s : %s\n", "BR Margin, Max [66]", s.BrMax) +
		fmt.Sprintf("%-50s : %s\n", "BR Margin, Min [67]", s.BrMin) +
		fmt.Sprintf("%-50s : %s\n", "Vendor SN [68-83]", s.VendorSn) +
		fmt.Sprintf("%-50s : %s\n", "Date Code [84-91]", s.DateCode)

	if s.Vendor.String() == "Arista Networks" && strings.HasPrefix(s.VendorPn.String(), "CAB-Q-S-") {
		str += fmt.Sprintf("%-50s : %x\n", "Vendor SA [120]", s.VendorAristaSa)
	}

	// Address A2h diagnostics
	str += fmt.Sprintf("%-50s : %s\n", "Temperature [A2h 96-97]", s.Temperature) +
		fmt.Sprintf("%-50s : %s\n", "Vcc [A2h 98-99]", s.Vcc) +
		fmt.Sprintf("%-50s : %s\n", "TX Bias [A2h 100-101]", s.TxBias) +
		fmt.Sprintf("%-50s : %s\n", "TX Power [A2h 102-103]", s.TxPower) +
		fmt.Sprintf("%-50s : %s\n", "RX Power [A2h 104-105]", s.RxPower)

	return str
}

func strCol(k string, v string, c1 string, c2 string) string {
	return fmt.Sprintf("%s%-50s%s : %s%s%s\n", c1, k, clear, c2, v, clear)
}

func joinStrCol(k string, l []string, c1 string, c2 string) string {
	if len(l) < 1 {
		return ""
	}

	r := strCol(k, l[0], c1, c2)
	for _, s := range l[1:] {
		r += strCol("", s, c1, c2)
	}
	return r
}

func (s *Sff8079) StringCol() string {
	str := strCol("Identifier [0]", fmt.Sprintf("0x%02x (%s)", byte(s.Identifier), s.Identifier), cyan, green) +
		strCol("Extended Identifier [1]", fmt.Sprintf("0x%02x (%s)", byte(s.ExtIdentifier), s.ExtIdentifier), cyan, green) +
		strCol("Connector [2]", fmt.Sprintf("0x%02x (%s)", byte(s.Connector), s.Connector), cyan, green) +
		strCol("Transceiver Codes [3-10]", fmt.Sprintf("0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x", s.Transceiver[0], s.Transceiver[1], s.Transceiver[2], s.Transceiver[3], s.Transceiver[4], s.Transceiver[5], s.Transceiver[6], s.Transceiver[7]), cyan, green) +
		joinStrCol("Transceiver Type", s.Transceiver.List(), cyan, yellow) +
		strCol("Encoding [11]", fmt.Sprintf("0x%02x (%s)", byte(s.Encoding), s.Encoding), cyan, green) +
		strCol("BR, Nominal [12]", s.BrNominal.String(), cyan, green) +
		strCol("Rate Identifier [13]", fmt.Sprintf("0x%02x", s.RateIdentifier), cyan, green) +
		strCol("Length (SMF) [14]", s.LengthSmfKm.String(), cyan, green) +
		strCol("Length (SMF) [15]", s.LengthSmfM.String(), cyan, green) +
		strCol("Length (50um) [16]", s.Length50umM.String(), cyan, green) +
		strCol("Length (62.5um) [17]", s.Length625umM.String(), cyan, green) +
		strCol("Length (Copper) [18]", s.LengthCopper.String(), cyan, green) +
		strCol("Length (OM3) [19]", s.LengthOm3.String(), cyan, green) +
		strCol("Vendor [20-35]", s.Vendor.String(), cyan, green) +
		strCol("Vendor OUI [37-39]", s.VendorOui.String(), cyan, green) +
		strCol("Vendor PN [40-55]", s.VendorPn.String(), cyan, green) +
		strCol("Vendor Rev [56-59]", s.VendorRev.String(), cyan, green) +
		strCol("Option Values [64-65]", s.Options.String(), cyan, green) +
		strCol("BR Margin, Max [66]", s.BrMax.String(), cyan, green) +
		strCol("BR Margin, Min [67]", s.BrMin.String(), cyan, green) +
		strCol("Vendor SN [68-83]", s.VendorSn.String(), cyan, green) +
		strCol("Date Code [84-91]", s.DateCode.String(), cyan, green)

	if s.Vendor.String() == "Arista Networks" && strings.HasPrefix(s.VendorPn.String(), "CAB-Q-S-") {
		str += strCol("Vendor SA [120]", fmt.Sprintf("%x", s.VendorAristaSa), cyan, green)
	}

	// Address A2h diagnostics
	str += strCol("Temperature [A2h 96-97]", s.Temperature.String(), cyan, green) +
		strCol("Vcc [A2h 98-99]", s.Vcc.String(), cyan, green) +
		strCol("TX Bias [A2h 100-101]", s.TxBias.String(), cyan, green) +
		strCol("TX Power [A2h 102-103]", s.TxPower.String(), cyan, green) +
		strCol("RX Power [A2h 104-105]", s.RxPower.String(), cyan, green)

	return str
}
