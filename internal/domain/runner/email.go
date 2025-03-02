package runner

import (
	"errors"
	"regexp"
)

type emailAddress string

var (
	emailValidationRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// ErrInvalidEmail Error when the email address is invalid
	ErrInvalidEmail = errors.New("invalid email address")
)

// newEmailAddress Creates a new emailAddress
func newEmailAddress(email string) (emailAddress, error) {
	//validate the email address with a regex
	if !emailValidationRegex.MatchString(email) {
		return "", ErrInvalidEmail
	}

	return emailAddress(email), nil
}

func (e emailAddress) String() string {
	return string(e)
}
