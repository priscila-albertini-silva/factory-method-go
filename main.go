package main

import "github.com/priscila-albertini/factory-method-go/internal"

func main() {
	// Criando uma fábrica de documentos PDF
	pdfFactory := &internal.PDFDocumentFactory{}
	pdfDocument := pdfFactory.CreateDocument()
	pdfDocument.Open()
	pdfDocument.Close()

	// Criando uma fábrica de planilhas
	spreadsheetFactory := &internal.SpreadsheetDocumentFactory{}
	spreadsheedDocument := spreadsheetFactory.CreateDocument()
	spreadsheedDocument.Open()
	spreadsheedDocument.Close()
}
