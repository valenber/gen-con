package applicants_test

import (
	"testing"

	applicants "github.com/valenber/gen-con/modules"
)

func TestGetApplicantsNotFound(t *testing.T) {
  res, err := applicants.GetApplicant("223")

  if err == nil {
    t.Errorf("Expected to throw with id 233, but got %s", res.Name)
  }
}

func TestGetApplicantFound(t *testing.T) {
  res, _ := applicants.GetApplicant("3")

  if res.Name != "Laura Maison" {
    t.Errorf("Expected Laura Maison with id 3, but got %s", res.Name)
  }
}
