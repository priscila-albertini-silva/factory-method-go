package internal

import "fmt"

type PDFDocument struct{}

func (p *PDFDocument) Open() {
	fmt.Println("Abrindo documento PDF")
}

func (p *PDFDocument) Close() {
	fmt.Println("Fechando documento PDF")
}

type SpreadsheetDocument struct{}

func (s *SpreadsheetDocument) Open() {
	fmt.Println("Abrindo planilha")
}

func (s *SpreadsheetDocument) Close() {
	fmt.Println("Fechando planilha")
}
