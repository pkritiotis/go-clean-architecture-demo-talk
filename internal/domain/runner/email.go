package runner

import (
	"errors"
	"regexp"
)

type EmailAddress string

var (
	emailValidationRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	ErrInvalidEmail = errors.New("invalid email address")
)

// NewEmailAddress Creates a new EmailAddress
func NewEmailAddress(email string) (EmailAddress, error) {
	//validate the email address with a regex
	if !emailValidationRegex.MatchString(email) {
		return "", ErrInvalidEmail
	}

	return EmailAddress(email), nil
}

func (e EmailAddress) String() string {
	return string(e)
}
