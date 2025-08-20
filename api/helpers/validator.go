package helpers

import (
	"fmt"
	"regexp"
)

const EMAIL_REGEX = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
const PHONE_REGEX = "^[+]{1}[0-9]{10,15}$"

func IsEmpty(str string, propertyName string) error {
	if len(str) == 0 {
		return fmt.Errorf("%v cannot be empty", propertyName)
	}
	return nil
}

func CheckStringBounds(str string, propertyName string, min int, max int) error {
	if len(str) < min {
		return fmt.Errorf("%v must be more than %v length", propertyName, min)
	}
	if len(str) > max {
		return fmt.Errorf("%v must be less than %v length", propertyName, max)
	}
	return nil
}

func IsEmail(str string) error {
	match, _ := regexp.MatchString(str, EMAIL_REGEX)
	if !match {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func IsValidPhoneNumber(str string) error {
	match, _ := regexp.MatchString(str, PHONE_REGEX)
	if !match {
		return fmt.Errorf("invalid phone number. must start with \"+\" and then 10/15 numbers")
	}
	return nil

}

func IsValidNumberRange(n int, propertName string, min int, max int) error {
	if n > max {
		return fmt.Errorf("%v must be less than %v", propertName, max)
	}
	if n < min {
		return fmt.Errorf("%v must be more than %v", propertName, min)
	}
	return nil
}
func IsNumberMoreThanZero(n int, propertyName string) error {
	if n <= 0 {
		return fmt.Errorf("%v must be more than zero", propertyName)
	}
	return nil
}

func Validate(checks ...func() error) (err error) {
	for _, check := range checks {
		err = check()
		if err != nil {
			return err
		}
	}
	return nil
}
