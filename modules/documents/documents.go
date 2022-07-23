package documents

import (
	"bytes"
	"fmt"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func MakePdf(htmlString string) ([]byte, error) {
  pdfg, err := wkhtmltopdf.NewPDFGenerator()

  var nilBuffer []byte

  if err != nil {
    fmt.Println(err)
    return nilBuffer, err
  }

  page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(htmlString)))

  // process styles
  page.EnableLocalFileAccess.Set(true)

  pdfg.AddPage(page)

  err = pdfg.Create()

  if err != nil {
    fmt.Println(err)
    return nilBuffer, err
  }

  return pdfg.Bytes(), nil
}
