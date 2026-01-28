package utils

import (
	"time"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

type PDFContent struct {
	Title           string
	InstitutionName string
	Address         string
	Phone           string
	Content         string
	LogoPath        string
}


func GenerateFullPDF(filePath string, data PDFContent) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AliasNbPages("")

	// HEADER
	pdf.SetHeaderFunc(func() {
		// LOGO (LEFT)
		if data.LogoPath != "" {
			pdf.Image(data.LogoPath, 10, 10, 20, 0, false, "", 0, "")
		}

		// INSTITUTION NAME (CENTER)
		pdf.SetFont("Arial", "B", 14)
		pdf.CellFormat(0, 10, data.InstitutionName, "", 1, "C", false, 0, "")

		// ADDRESS & CONTACT
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(0, 6, data.Address+" | "+data.Phone, "", 1, "C", false, 0, "")
		pdf.Ln(5)
	})


	// FOOTER
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "", 8)
		pdf.CellFormat(0, 10,
			"Page "+strconv.Itoa(pdf.PageNo())+" of {nb} | Generated at "+time.Now().Format(time.RFC3339),
			"", 0, "C", false, 0, "",
		)
	})

	pdf.AddPage()

	// TITLE
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, data.Title)
	pdf.Ln(12)

	// DATE
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 8, "Date: "+time.Now().Format("02 January 2006"))
	pdf.Ln(12)

	// CONTENT
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 8, data.Content, "", "", false)

	return pdf.OutputFileAndClose(filePath)
}