package validator

import (
	"fmt"

	"github.com/hasona23/workis/api/models"
)

func ValidateQualification(qualification models.Qualification) error {
	if qualification.WorkerId < 0 {
		return fmt.Errorf("Invalid worker id")
	}
	err := isValidString(qualification.CertName, 3, 32, "Qualification Certificate Name")
	if err != nil {
		return err
	}
	return nil
}
