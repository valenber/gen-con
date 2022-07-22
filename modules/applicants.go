package applicants

import "errors"

type Applicant struct {
  ID   string `json:"id"`
  Template string `json:"templates"`
  Name string `json:"name"`
  AcceptanceDate string `json:"acceptance_date"`
}

var applicants = []Applicant{
  {ID: "1", Name: "Janise Joplin", Template: "1", AcceptanceDate: "01-07-2022"},
  {ID: "2", Name: "Robin Bobbin", Template: "1", AcceptanceDate: "21-03-2022"},
  {ID: "3", Name: "Laura Maison", Template: "1", AcceptanceDate: ""},
  {ID: "4", Name: "Harriet Tubman", Template: "1", AcceptanceDate: "12-10-2022"},
}

func GetApplicant(id string) (*Applicant, error) {
  for _, applicant := range applicants {
    if applicant.ID == id {
      return &applicant, nil 
    }
  }
  return nil, errors.New("Can not find applicant")
}
