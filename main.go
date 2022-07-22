package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func main()  {
  router := gin.Default();

  router.GET("/contract/:applicant_id", getContract)

  router.Run("localhost:3030")
}

func getContract(c *gin.Context) {
  id := c.Param("applicant_id")
  applicant, err := getApplicant(id)

  if err != nil {
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Can not find applicant with id %s", id)})
  } else {
    c.IndentedJSON(http.StatusOK, applicant)
  }
}

func getApplicant(id string) (*Applicant, error) {
  for _, applicant := range applicants {
    if applicant.ID == id {
      return &applicant, nil 
    }
  }
  return nil, errors.New("Can not find applicant")
}
