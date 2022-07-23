package templates

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"

	"github.com/valenber/gen-con/modules/applicants"
)

func BuidTemplate(applicant applicants.Applicant) (string, error) {
  t, err := template.ParseFiles("modules/templates/test-template.html")

  if err != nil {
    fmt.Println(err)
    return "", errors.New("Failed to parse template file")
  }

  var page bytes.Buffer

  if err := t.Execute(&page, applicant); err != nil {
    fmt.Println(err)
    return "", errors.New("Failed to populate template with data")
  }

  return page.String(), nil
}
