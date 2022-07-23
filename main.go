package main

import (
	"fmt"
	"net/http"

	"github.com/valenber/gen-con/modules/applicants"
	"github.com/valenber/gen-con/modules/documents"
	"github.com/valenber/gen-con/modules/templates"

	"github.com/gin-gonic/gin"
)

func main()  {
  router := gin.Default();

  router.GET("/contracts/:applicant_id", getContract)

  router.Run("localhost:3030")
}

func getContract(c *gin.Context) {
  id := c.Param("applicant_id")
  applicant, err := applicants.GetApplicant(id)

  if err != nil {
    c.IndentedJSON(http.StatusNotFound, gin.H{
      "message": "Failed to get applicant",
      "error": err.Error(),
      "details": fmt.Sprintf("applicant.ID=%s", id),
    })
    return
  }

  page, err := templates.BuidTemplate(*applicant)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{
      "message": "Failed to build template",
      "error": err.Error(),
      "details": fmt.Sprintf("applicant.ID=%s", applicant.ID),
    })
    return
  }

  pdf, err := documents.MakePdf(page)

  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{
      "message": "Failed to make PDF",
      "details": fmt.Sprintf("applicant.ID=%s", applicant.ID),
      "error": err.Error(),
    })
    return
  }

  c.Writer.Header().Set("Content-Disposition", "attachment; filename=contract.pdf")
  c.Writer.Header().Set("Content-type", "application/pdf")
  c.Data(http.StatusOK, "ok", pdf)
}
