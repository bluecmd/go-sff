# go-sff

[![Container Tests](https://github.com/bluecmd/go-sff/workflows/Container%20Tests/badge.svg)](https://github.com/bluecmd/go-sff/actions)

A Go library for reading and parsing SFF (Small Form Factor) transceiver EEPROM data from network devices. This library supports both SFF-8079 (SFP) and SFF-8636 (QSFP) standards.

## Overview

The `go-sff` library provides functionality to:
- Read EEPROM data from SFF transceivers via I2C interface
- Automatically detect transceiver type (SFF-8079 or SFF-8636)
- Parse and decode EEPROM data according to industry standards
- Extract detailed information about transceivers including vendor details, specifications, and capabilities
- Output data in both human-readable and JSON formats

## Supported Standards

### SFF-8079 (SFP)
- SFP (Small Form-factor Pluggable) transceivers
- 1G and 10G Ethernet modules
- Fiber and copper modules
- InfiniBand modules
- SONET/SDH modules
- Fibre Channel modules

### SFF-8636 (QSFP)
- QSFP (Quad Small Form-factor Pluggable) transceivers
- QSFP+ modules
- QSFP28 modules
- High-density 40G and 100G Ethernet modules

## Features

- Automatically identifies transceiver type from EEPROM data
- Direct I2C interface for reading EEPROM data

## Installation

```bash
go get github.com/bluecmd/go-sff
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/bluecmd/go-sff"
)

func main() {
    // Read transceiver data from I2C device
	module, err := sff.Read(sff.NewI2CReader("/dev/i2c-0"))
    if err != nil {
        log.Fatal(err)
    }

    // Print human-readable output
    fmt.Println(module.String())

    // Print colored output for terminal
    fmt.Println(module.StringCol())

    // Access specific fields
    switch module.Type {
    case sff.TypeSff8079:
        fmt.Printf("Vendor: %s\n", module.Sff8079.Vendor)
        fmt.Printf("Part Number: %s\n", module.Sff8079.VendorPn)
        fmt.Printf("Serial Number: %s\n", module.Sff8079.VendorSn)
    case sff.TypeSff8636:
        fmt.Printf("Vendor: %s\n", module.Sff8636.Vendor)
        fmt.Printf("Part Number: %s\n", module.Sff8636.VendorPn)
        fmt.Printf("Serial Number: %s\n", module.Sff8636.VendorSn)
    }
}
```

### Manual Type Detection

```go
// If you already have EEPROM data
eepromData := []byte{...} // 256+ bytes of EEPROM data

// Detect the type
moduleType, err := sff.GetType(eepromData)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Module type: %s\n", moduleType)
```

### Working with Specific Standards

```go
import (
    "github.com/bluecmd/go-sff/sff8079"
    "github.com/bluecmd/go-sff/sff8636"
)

// For SFF-8079 modules
if module.Type == sff.TypeSff8079 {
    sfp := module.Sff8079
    fmt.Printf("Transceiver: %s\n", sfp.Transceiver)
    fmt.Printf("Encoding: %s\n", sfp.Encoding)
    fmt.Printf("Bit Rate: %s\n", sfp.BrNominal)
}

// For SFF-8636 modules
if module.Type == sff.TypeSff8636 {
    qsfp := module.Sff8636
    fmt.Printf("Transceiver: %s\n", qsfp.Transceiver)
    fmt.Printf("Encoding: %s\n", qsfp.Encoding)
    fmt.Printf("Bit Rate: %s\n", qsfp.BrNominal)
}
```

## Reading SFF EEPROM Data

The library provides a flexible interface-based approach for reading SFF EEPROM data. You can implement your own reader or use the built-in I2C reader.

### Using the Reader Interface

```go
// Define a custom reader
type CustomReader struct {
    // your fields
}

func (r *CustomReader) Read() ([]byte, error) {
    // implement your reading logic
    return eepromData, nil
}

// Use your custom reader
reader := &CustomReader{}
module, err := sff.Read(reader)
```

### Using the Built-in I2C Reader

```go
// Create an I2C reader
reader := sff.NewI2CReader("/dev/i2c-0")
module, err := sff.Read(reader)

// Or use the convenience function for backward compatibility
module, err := sff.ReadFromPath("/dev/i2c-0")
```

### I2C Interface

The library includes a low-level I2C interface for reading EEPROM data:

```go
// Create I2C connection
i2c, err := sff.NewI2C("/dev/i2c-0", 0x50) // 0x50 is typical SFP address
if err != nil {
    log.Fatal(err)
}
defer i2c.Close()

// Read data
data := make([]byte, 256)
i2c.Write([]byte{0x00}) // Set address pointer to 0
i2c.Read(data)
```

## Requirements

- Go 1.19 or later

## Building the sfpdiag

```bash
make sfpdiag
```

## Running Tests

```bash
make test
make test-container
```

## License

This project is licensed under the terms specified in the LICENSE file.

## Standards glossary

| Standard | Description | Status |
|----------|-------------|---------|
| SFF-8079 | SFP Management Interface | Supported |
| SFF-8636 | QSFP Management Interface | Supported |
| SFF-8472 | Diagnostic Monitoring Interface for Optical Transceivers | Referenced |
| SFF-8690 | Tunable SFP+ Memory Map for ITU Frequencies | Not yet supported |
