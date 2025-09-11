package sff

import (
	"errors"
	"fmt"
	"os"

	"github.com/bluecmd/go-sff/sff8079"
	"github.com/bluecmd/go-sff/sff8636"
)

// Reader interface defines how to read SFF EEPROM data
type Reader interface {
	Read() ([]byte, error)
}

// Type of eeprom module.
type Type string

const (
	TypeUnknown = Type("Unknown")
	TypeSff8079 = Type("SFF-8079")
	TypeSff8636 = Type("SFF-8636")
)

var ErrUnknownType = errors.New("unknown type")

type Module struct {
	Type Type
	*sff8079.Sff8079
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

	// 0xB is "DWDM-SFP/SFP+ (not using SFF-8472)" so technically it shouldn't be
	// compatible with SFF-8079 but in reality it seems to be.
	if (eeprom[0] == 2 || eeprom[0] == 3 || eeprom[0] == 0xb) && eeprom[1] == 4 {
		return TypeSff8079, nil
	}

	if eeprom[128] == 12 || eeprom[128] == 13 || eeprom[128] == 17 {
		return TypeSff8636, nil
	}

	return TypeUnknown, fmt.Errorf("eeprom unknown type")
}

// I2CReader implements Reader interface for I2C devices
type I2CReader struct {
	path string
}

// NewI2CReader creates a new I2CReader for the given device path
func NewI2CReader(path string) *I2CReader {
	return &I2CReader{path: path}
}

// Read implements the Reader interface for I2C devices
func (r *I2CReader) Read() ([]byte, error) {
	// 0x50 and 0x51 SFP port
	i, err := NewI2C(r.path, 0x50)
	if err != nil {
		return nil, err
	}
	defer i.Close()
	i.Write([]byte{0x00})
	b := make([]byte, 256)
	i.Read(b)
	return b, nil
}

// FileReader implements Reader interface for file-based reading
type FileReader struct {
	path string
}

// NewFileReader creates a new FileReader for the given file path
func NewFileReader(path string) *FileReader {
	return &FileReader{path: path}
}

// Read implements the Reader interface for file-based reading
func (r *FileReader) Read() ([]byte, error) {
	file, err := os.Open(r.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b := make([]byte, 256)
	n, err := file.Read(b)
	if err != nil {
		return nil, err
	}

	// If we read less than 256 bytes, pad with zeros
	if n < 256 {
		for i := n; i < 256; i++ {
			b[i] = 0
		}
	}

	return b, nil
}

// Read reads SFF EEPROM data using the provided Reader interface
func Read(reader Reader) (*Module, error) {
	eeprom, err := reader.Read()
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
