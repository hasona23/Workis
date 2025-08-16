package validator

import (
	"fmt"
	"strings"
	"time"

	"github.com/hasona23/workis/api/models"
)

func ValidateWorker(worker models.Worker) error {
	err := isValidString(worker.Name, 4, 32, "worker name")
	if err != nil {
		return err
	}

	err = isValidString(worker.Position, 1, 32, "worker position")
	if err != nil {
		return err
	}
	if worker.Salary <= 0 {
		return fmt.Errorf("salary cant be less than or equal to zero")
	}
	if (time.Now().UTC().Unix() - worker.BirthDate.Unix()) < 60*60*24*12*16 {
		return fmt.Errorf("worker birth date must imply he is older than 16 years old")
	}
	if worker.Gender != models.MALE && worker.Gender != models.FEMALE {
		return fmt.Errorf("worker must be %v or %v", models.MALE, models.FEMALE)
	}
	err = isValidString(worker.Degree, 2, 32, "worker degree")
	if err != nil {
		return err
	}
	err = isValidString(worker.JobDescription, 16, 128, "worker job description")
	if err != nil {
		return err
	}
	return nil
}

func isStringWhiteSpaces(str string) bool {
	trimName := strings.Replace(str, " ", "", -1)
	return len(trimName) == 0
}
func isValidStringLength(str string, min int, max int, msgTitle string) error {
	if len(str) < min {
		return fmt.Errorf("%v {%v} must be %v or more characters length", msgTitle, str, min)
	}
	if len(str) > max+1 {
		return fmt.Errorf("%v {%v} must be less than %v characters", msgTitle, str, max)
	}
	return nil
}
func isValidString(str string, min int, max int, msgTitle string) error {
	err := isValidStringLength(str, min, max, msgTitle)
	if err != nil {
		return err
	}
	if isStringWhiteSpaces(str) {
		return fmt.Errorf("%v must not be only whitespaces", msgTitle)
	}
	return nil
}
