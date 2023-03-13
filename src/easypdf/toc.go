package easypdf

import htmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"

type Toc struct {
	Include             bool
	DisableDottedLines  bool
	DisableTocLinks     bool
	TocHeaderText       string
	TocLevelIndentation uint
	TocTextSizeShrink   float64
	//XslStyleSheet     string
}

func (toc *Toc) SetTOCFields(pdfg *htmltopdf.PDFGenerator) {
	pdfg.TOC.Include = toc.Include
	pdfg.TOC.DisableDottedLines.Set(toc.DisableDottedLines)
	pdfg.TOC.DisableTocLinks.Set(toc.DisableTocLinks)
	pdfg.TOC.TocHeaderText.Set(toc.TocHeaderText)
	pdfg.TOC.TocLevelIndentation.Set(toc.TocLevelIndentation)
	pdfg.TOC.TocTextSizeShrink.Set(toc.TocTextSizeShrink)
}
