package savers

import (
	"fmt"
	"strings"

	"example.com/go-cert/cert"
	"example.com/go-cert/cert/savers/html"
	"example.com/go-cert/cert/savers/pdf"
)

func GetSaver(outputType string, path string) (cert.Saver, error) {
	var err error
	var s cert.Saver

	switch strings.ToLower(outputType) {
	case "pdf":
		s, err = pdf.New(path)
	case "html":
		s, err = html.New(path)
	default:
		err = fmt.Errorf("Unknown output type, got '%v'", outputType)
	}
	if err != nil {
		return nil, err
	}

	return s, err
}
