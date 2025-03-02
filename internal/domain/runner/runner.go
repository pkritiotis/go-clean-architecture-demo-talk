// Package runner contains the Runner model.
package runner

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type EmailAddress string

var (
	emailValidationRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	ErrInvalidEmail            = errors.New("invalid email address")
	ErrRunnerNameCannotBeEmpty = errors.New("name cannot be empty")
)

// NewEmailAddress Creates a new EmailAddress
func NewEmailAddress(email string) (EmailAddress, error) {
	//validate the email address with a regex
	if !emailValidationRegex.MatchString(email) {
		return "", ErrInvalidEmail
	}

	return EmailAddress(email), nil
}

// Runner Model that represents the Runner
type Runner struct {
	id           uuid.UUID
	name         string
	emailAddress EmailAddress
	createdAt    time.Time
}

// NewRunner Creates a new Runner
func NewRunner(name, emailAddress string) (Runner, error) {

	//validate the name
	if name == "" {
		return Runner{}, ErrRunnerNameCannotBeEmpty
	}

	//validate the email address
	email, err := NewEmailAddress(emailAddress)
	if err != nil {
		return Runner{}, err
	}

	//create a new UUID for the runner
	id := uuid.New()

	return Runner{
		id:           id,
		name:         name,
		emailAddress: email,
		createdAt:    time.Now().UTC(),
	}, nil
}
