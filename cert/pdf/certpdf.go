package pdf

import (
	"fmt"
	"os"
	"path"

	"example.com/go-cert/cert"
	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputDir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &PdfSaver{
		OutputDir: outputDir,
	}

	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, true)
	pdf.AddPage()

	background(pdf)

	header(pdf, &cert)
	body(pdf, &cert)
	footer(pdf)

	// save file
	file := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, file)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}

	fmt.Printf("Certificate %v has been generated\n", file)
	return nil
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	width, height := pdf.GetPageSize()
	pdf.ImageOptions("img/background.png", 0, 0, width, height, false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	marginH := 20.0
	imageWidth := 30.0
	pageWidth, _ := pdf.GetPageSize()
	filename := "img/gopher.png"

	pdf.ImageOptions(filename, margin, marginH, imageWidth, 0, false, opts, 0, "")
	pdf.ImageOptions(filename, pageWidth-imageWidth-margin, marginH, imageWidth, 0, false, opts, 0, "")

	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

func body(pdf *gofpdf.Fpdf, c *cert.Cert) {
	pdf.Ln(30)
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelPresented, "C")

	pdf.Ln(30)
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, c.Name, "C")

	pdf.Ln(30)
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelParticipation, "C")

	pdf.Ln(30)
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, c.LabelDate, "C")
}

func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 20.0
	marginH := 10.0
	imageWidth := 50.0
	pageWidth, pageHeight := pdf.GetPageSize()
	filename := "img/stamp.png"

	pdf.ImageOptions(filename, pageWidth-imageWidth-margin, pageHeight-imageWidth-marginH, imageWidth, 0, false, opts, 0, "")
}
