package templates

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/valenber/gen-con/modules/applicants"
)

func BuidTemplate(applicant applicants.Applicant) (string, error) {

  templatePath := fmt.Sprintf("modules/templates/collection/%s.html", applicant.ContractType)

  t, err := template.ParseFiles(fmt.Sprintf(templatePath))

  if err != nil {
    return "", err
  }

  var page bytes.Buffer

  if err := t.Execute(&page, applicant); err != nil {
    return "", err
  }

  return page.String(), nil
}
