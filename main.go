package main

import (
	"fmt"
	"net/http"
  applicants "github.com/valenber/gen-con/modules"

	"github.com/gin-gonic/gin"
)

func main()  {
  router := gin.Default();

  router.GET("/contract/:applicant_id", getContract)

  router.Run("localhost:3030")
}

func getContract(c *gin.Context) {
  id := c.Param("applicant_id")
  applicant, err := applicants.GetApplicant(id)

  if err != nil {
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Can not find applicant with id %s", id)})
  } else {
    c.IndentedJSON(http.StatusOK, applicant)
  }
}
