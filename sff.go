package sff

import (
	"errors"
	"fmt"

	"github.com/bluecmd/go-sff/sff8079"
	"github.com/bluecmd/go-sff/sff8636"
)

// Type of eeprom module.
type Type string

const (
	TypeUnknown = Type("Unknown")
	TypeSff8079 = Type("SFF-8079")
	TypeSff8636 = Type("SFF-8636")
)

var ErrUnknownType = errors.New("unknown type")

type Module struct {
	Type             Type
	*sff8079.Sff8079
	*sff8636.Sff8636
}

type module Module

type moduleSff8079 struct {
	Type Type
	*sff8079.Sff8079
}

type moduleSff8636 struct {
	Type Type
	*sff8636.Sff8636
}

func (m *Module) String() string {
	switch m.Type {
	case TypeSff8079:
		return m.Sff8079.String()
	case TypeSff8636:
		return m.Sff8636.String()
	}
	return ""
}

func (m *Module) StringCol() string {
	switch m.Type {
	case TypeSff8079:
		return m.Sff8079.StringCol()
	case TypeSff8636:
		return m.Sff8636.StringCol()
	}
	return ""
}

func GetType(eeprom []byte) (Type, error) {
	if len(eeprom) < 256 {
		return TypeUnknown, fmt.Errorf("eeprom size to small needs to be 256 bytes or larger got: %d bytes", len(eeprom))
	}

	if (eeprom[0] == 2 || eeprom[0] == 3) && eeprom[1] == 4 {
		return TypeSff8079, nil
	}

	if eeprom[128] == 12 || eeprom[128] == 13 || eeprom[128] == 17 {
		return TypeSff8636, nil
	}

	return TypeUnknown, fmt.Errorf("eeprom unknown type")
}

func readI2C(f string) ([]byte, error) {
	// 0x50 and 0x51 SFP port
	i, err := NewI2C(f, 0x50)
	if err != nil {
		return nil, err
	}
	defer i.Close()
	i.Write([]byte{0x00})
	b := make([]byte, 256)
	i.Read(b)
	return b, nil
}

func Read(f string) (*Module, error) {
	eeprom, err := readI2C(f)
	if err != nil {
		return nil, err
	}

	t, err := GetType(eeprom)
	if err != nil {
		return nil, err
	}

	switch t {
	case TypeSff8079:
		m, err := sff8079.Decode(eeprom)
		if err != nil {
			return nil, err
		}
		return &Module{Type: TypeSff8079, Sff8079: m}, nil
	case TypeSff8636:
		m, err := sff8636.Decode(eeprom)
		if err != nil {
			return nil, err
		}
		return &Module{Type: TypeSff8636, Sff8636: m}, nil
	}
	return nil, ErrUnknownType
}
