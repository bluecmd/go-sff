package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bluecmd/go-sff"
)

func main() {
	var (
		devicePath = flag.String("device", "/dev/i2c-0", "I2C device path")
		outputJSON = flag.Bool("json", false, "Output in JSON format")
		outputCol  = flag.Bool("color", false, "Output with colors")
		help       = flag.Bool("help", false, "Show help")
	)
	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s -device /dev/i2c-1\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -json\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -color\n", os.Args[0])
		os.Exit(0)
	}

	// Read transceiver data
	module, err := sff.Read(*devicePath)
	if err != nil {
		log.Fatalf("Failed to read transceiver: %v", err)
	}

	// Output based on flags
	if *outputJSON {
		// JSON output
		data, err := json.MarshalIndent(module, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal JSON: %v", err)
		}
		fmt.Println(string(data))
	} else if *outputCol {
		// Colored output
		fmt.Println(module.StringCol())
	} else {
		// Standard output
		fmt.Println(module.String())
	}

	// Print summary information
	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Module Type: %s\n", module.Type)

	switch module.Type {
	case sff.TypeSff8079:
		sfp := module.Sff8079
		fmt.Printf("Vendor: %s\n", sfp.Vendor)
		fmt.Printf("Part Number: %s\n", sfp.VendorPn)
		fmt.Printf("Serial Number: %s\n", sfp.VendorSn)
		fmt.Printf("Date Code: %s\n", sfp.DateCode)
		fmt.Printf("Connector: %s\n", sfp.Connector)
		fmt.Printf("Bit Rate: %s\n", sfp.BrNominal)

	case sff.TypeSff8636:
		qsfp := module.Sff8636
		fmt.Printf("Vendor: %s\n", qsfp.Vendor)
		fmt.Printf("Part Number: %s\n", qsfp.VendorPn)
		fmt.Printf("Serial Number: %s\n", qsfp.VendorSn)
		fmt.Printf("Date Code: %s\n", qsfp.DateCode)
		fmt.Printf("Connector: %s\n", qsfp.Connector)
		fmt.Printf("Bit Rate: %s\n", qsfp.BrNominal)

	default:
		fmt.Printf("Unknown module type\n")
	}
}
