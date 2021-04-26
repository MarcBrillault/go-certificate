package main

import (
	"fmt"
	"os"

	"example.com/go-cert/cert"
	"example.com/go-cert/cert/pdf"
)

func main() {
	c, err := cert.New("Golang Programing", "Bob Dylan", "2018-06-21")
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}

	var saver cert.Saver
	saver, err = pdf.New("output")
	if err != nil {
		fmt.Printf("Error during pdf generation: %v\n", err)
		os.Exit(1)
	}

	saver.Save(*c)
}
