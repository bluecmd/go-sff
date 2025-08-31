package sff8636

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	PwrClassMask = 0xC0
	PwrClass1    = (0 << 6)
	PwrClass2    = (1 << 6)
	PwrClass3    = (2 << 6)
	PwrClass4    = (3 << 6)

	ExtPwrClass8Mask = 0x20
	ExtPwrClass8     = (1 << 5)

	ClieCodeMask = 0x10
	NoClieCode   = (0 << 4)
	ClieCode     = (1 << 4)

	CdrInTxMask = 0x08
	NoCdrInTx   = (0 << 3)
	CdrInTx     = (1 << 3)

	CdrInRxMask = 0x04
	NoCdrInRx   = (0 << 2)
	CdrInRx     = (1 << 2)

	ExtPwrClassMask   = 0x03
	ExtPwrClassUnused = 0
	ExtPwrClass5      = 1
	ExtPwrClass6      = 2
	ExtPwrClass7      = 3
)

var pwrClassNames = map[byte]string{
	PwrClass1: "1.5 W max. power consumption",
	PwrClass2: "2.0 W max. power consumption",
	PwrClass3: "2.5 W max. power consumption",
	PwrClass4: "3.5 W max. power consumption (or Power Classes 5, 6 or 7)",
}

var clieCodeNames = map[byte]string{
	NoClieCode: "No CLEI code present",
	ClieCode:   "CLEI code present",
}

var cdrInTxNames = map[byte]string{
	NoCdrInTx: "No CDR in TX",
	CdrInTx:   "CDR in TX",
}

var cdrInRxNames = map[byte]string{
	NoCdrInRx: "No CDR in RX",
	CdrInRx:   "CDR in RX",
}

var extPwrClassNames = map[byte]string{
	ExtPwrClassUnused: "unused (legacy setting)",
	ExtPwrClass5:      "4.0 W max. power consumption",
	ExtPwrClass6:      "4.5 W max. power consumption",
	ExtPwrClass7:      "5.0 W max. power consumption",
}

type ExtIdentifier byte

// GetPowerClass returns a single power class (1-8) based on the SFF-8636 specification
func (e ExtIdentifier) GetPowerClass() string {
	b := byte(e)

	// Check for Power Class 8 first (highest priority)
	if b&ExtPwrClass8Mask == ExtPwrClass8 {
		return "Power Class 8"
	}

	// Check extended power classes (5, 6, 7)
	extPwrClass := b & ExtPwrClassMask
	if extPwrClass != ExtPwrClassUnused {
		switch extPwrClass {
		case ExtPwrClass5:
			return "Power Class 5"
		case ExtPwrClass6:
			return "Power Class 6"
		case ExtPwrClass7:
			return "Power Class 7"
		}
	}

	// Base power classes (1, 2, 3, 4)
	basePwrClass := b & PwrClassMask
	switch basePwrClass {
	case PwrClass1:
		return "Power Class 1"
	case PwrClass2:
		return "Power Class 2"
	case PwrClass3:
		return "Power Class 3"
	case PwrClass4:
		// If ExtPwrClassUnused, translate to Power Class 4
		if extPwrClass == ExtPwrClassUnused {
			return "Power Class 4"
		}
		// Otherwise, the extended power class takes precedence
		return "Power Class 4"
	}

	return "Power Class Unknown"
}

func (e ExtIdentifier) List() []string {
	b := byte(e)
	s := []string{
		e.GetPowerClass(),
		clieCodeNames[b&ClieCodeMask],
		fmt.Sprintf("%s, %s", cdrInTxNames[b&CdrInTxMask], cdrInRxNames[b&CdrInRxMask]),
	}
	return s
}

func (e ExtIdentifier) String() string {
	return strings.Join(e.List(), "\n")
}

func (e ExtIdentifier) MarshalJSON() ([]byte, error) {
	b := byte(e)
	m := map[string]interface{}{
		"values": e.List(),
		"hex":    hex.EncodeToString([]byte{b}),
	}
	return json.Marshal(m)
}

func (e *ExtIdentifier) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*e = ExtIdentifier(b[0])
	return nil
}
