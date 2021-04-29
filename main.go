package main

import (
	"flag"
	"fmt"
	"os"

	"example.com/go-cert/cert"
	"example.com/go-cert/cert/html"
	"example.com/go-cert/cert/pdf"
)

func main() {
	outputType := flag.String("type", "pdf", "Output type of the certificate")
	flag.Parse()

	var saver cert.Saver
	var err error

	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		err = fmt.Errorf("Unknown output type, got '%v'", *outputType)
	}
	if err != nil {
		fmt.Printf("Could not create generator: %v\n", err)
		os.Exit(1)
	}

	c, err := cert.New("Golang Programing", "Bob Dylan", "2018-06-21")
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}

	saver.Save(*c)
}
