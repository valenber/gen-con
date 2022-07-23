package applicants

import (
	"errors"
)

type Applicant struct {
  ID   string `json:"id"`
  Name string `json:"name"`
  ContractType string `json:"contract_type"`
  AcceptanceDate string `json:"acceptance_date"`
}

var applicants = []Applicant{
  {ID: "1", Name: "Janise Joplin", ContractType: "test", AcceptanceDate: "01-07-2022"},
  {ID: "2", Name: "Robin Bobbin", ContractType: "type1", AcceptanceDate: "21-03-2022"},
  {ID: "3", Name: "Laura Maison", ContractType: "type1", AcceptanceDate: ""},
  {ID: "4", Name: "Harriet Tubman", ContractType: "type2", AcceptanceDate: "12-10-2022"},
}

func GetApplicant(id string) (*Applicant, error) {
  for _, applicant := range applicants {
    if applicant.ID == id {
      return &applicant, nil 
    }
  }

  return nil, errors.New("Applicant is not in the list")
}
