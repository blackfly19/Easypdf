package easypdf

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	htmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gomarkdown/markdown"
)

type MdToPdf struct {
	Files             []string
	Directory         string
	OutputFilename    string
	CssFilename       string
	CoverPageFileName string
	WatchMode         bool
	Toc
}

func (mdpdf *MdToPdf) ConvertFileToPDF() error {

	var pdfBytes = new(bytes.Buffer)

	pdfg, err := htmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	mdpdf.IntegrateCSS(pdfBytes)

	if mdpdf.Include {
		mdpdf.SetTOCFields(pdfg)
	}

	if mdpdf.CoverPageFileName != "" {
		coverFile, err := mdpdf.AddCoverPage(mdpdf.CoverPageFileName)
		if err != nil {
			return err
		}
		defer os.Remove(coverFile.Name())

		pdfg.Cover.Input = coverFile.Name()
	}

	for _, file := range mdpdf.Files {
		fmt.Println(file, mdpdf.CoverPageFileName)

		fileAbsPath, _ := filepath.Abs(file)
		coverPageAbsolutePath, _ := filepath.Abs(mdpdf.CoverPageFileName)

		if fileAbsPath != coverPageAbsolutePath {
			md, err := os.ReadFile(file)
			if err != nil {
				return err
			}

			pdfBytes.Write(markdown.ToHTML(md, nil, nil))
			pdfBytes.WriteString(`<P style="page-break-before: always">`)
		}
	}

	page := htmltopdf.NewPageReader(pdfBytes)
	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	if filepath.Ext(mdpdf.OutputFilename) == "" {
		mdpdf.OutputFilename = mdpdf.OutputFilename + ".pdf"
	}

	err = pdfg.WriteFile(mdpdf.OutputFilename)
	if err != nil {
		return err
	}

	return nil
}

func (mdpdf *MdToPdf) IntegrateCSS(html *bytes.Buffer) error {

	cssFile, err := os.ReadFile(mdpdf.CssFilename)
	if err != nil {
		return err
	}

	var cssStyleString string = "<style>" + string(cssFile) + "</style>"

	html.WriteString(cssStyleString)
	return nil
}

func (mdpdf *MdToPdf) AddCoverPage(coverPageFileName string) (*os.File, error) {
	coverFile, err := os.CreateTemp("", "*.html")
	if err != nil {
		return nil, err
	}

	coverhtml, err := os.ReadFile(coverPageFileName)
	if err != nil {
		return nil, err
	}

	coverFile.Write(markdown.ToHTML(coverhtml, nil, nil))
	return coverFile, nil
}
