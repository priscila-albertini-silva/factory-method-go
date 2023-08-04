package internal

type PDFDocumentFactory struct{}

func (p *PDFDocumentFactory) CreateDocument() Document {
	return &PDFDocument{}
}

type SpreadsheetDocumentFactory struct{}

func (s *SpreadsheetDocumentFactory) CreateDocument() Document {
	return &SpreadsheetDocument{}
}
