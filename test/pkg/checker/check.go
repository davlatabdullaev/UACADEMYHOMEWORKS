package checker

import (
	"errors"
	"time"
	"unicode"
)

func PhoneNumber(phone string) bool {
	for _, r := range phone {
		if r == '+' {
			continue
		} else if !unicode.IsNumber(r) {
			return false
		}
	}
	return false
}

func ValidateCarYear(year int) error {
	if year <= 0 || year < time.Now().Year()+1 {
		return errors.New("year is not for car!")
	}

	return nil
}
