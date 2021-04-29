package main

import (
	"flag"
	"fmt"
	"os"

	"example.com/go-cert/cert"
	"example.com/go-cert/cert/savers"
)

func main() {
	outputType := flag.String("type", "pdf", "Output type of the certificate")
	flag.Parse()

	var err error

	csvName := "students.csv"
	certs, err := cert.ParseCsv(csvName)
	if err != nil {
		fmt.Printf("Unable to open CSV file '%v': %v", csvName, err)
		os.Exit(1)
	}

	for _, cert := range certs {
		saver, err := savers.GetSaver(*outputType, "output")
		if err != nil {
			fmt.Printf("Could not create generator: %v", err)
			os.Exit(1)
		}

		saver.Save(*cert)
	}
}
