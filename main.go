package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Applicant struct {
  ID string `json:"id"`
  Name string `json:"name"`
  AcceptanceDate string `json:"acceptance_date"`
}

var applicants = []Applicant{
  {ID: "1", Name: "Janise Joplin", AcceptanceDate: "01-07-2022"},
  {ID: "2", Name: "Robin Bobbin", AcceptanceDate: "21-03-2022"},
  {ID: "3", Name: "Laura Maison", AcceptanceDate: ""},
  {ID: "4", Name: "Harriet Tubman", AcceptanceDate: "12-10-2022"},
}

func main()  {
  router := gin.Default();

  router.GET("/contract/:applicant_id", getContract)

  router.Run("localhost:3030")
}

func getContract(c *gin.Context) {
  id := c.Param("applicant_id")

  for _, a := range applicants {
    if a.ID == id {
      c.IndentedJSON(http.StatusOK, a)
      return
    }
  }

  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "applicant no found"})
}
