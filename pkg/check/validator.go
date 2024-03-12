package check

import (
	"errors"
	"strings"
	"time"
)

func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}

func Validategmail(gmail string) error {
	// Check if the email contains "@gmail.com"
	containsGmailCom := strings.Contains(gmail, "@gmail.com")
	// Check if the email contains "@gmail.ru"
	containsGmailRu := strings.Contains(gmail, "@gmail.ru")

	// If the email contains both domains, return an error
	if containsGmailCom && containsGmailRu {
		return errors.New("email address cannot contain both @gmail.com and @gmail.ru domains")
	}

	// If the email is valid (doesn't contain both domains), return nil
	return nil
}

func Validatenumber(number string) error {
	containsNumber := strings.Contains(number, "998")

	if containsNumber {
		return errors.New("number  first three numbers should be: 998for example: 998 700 00 00 ")
	}

	return nil
}
