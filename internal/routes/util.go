package routes

import (
	"regexp"

	"github.com/google/uuid"
)

func ValidateID (id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func ValidateEmail (email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(regex, email)
	if err != nil {
		return false
	}
	return matched
}
