// Package runner contains the Runner model.
package runner

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	// ErrRunnerNameCannotBeEmpty Error when the name of the runner is empty
	ErrRunnerNameCannotBeEmpty = errors.New("name cannot be empty")
)

// Runner Model that represents the Runner
type Runner struct {
	id           uuid.UUID
	name         string
	emailAddress emailAddress
	createdAt    time.Time
}

// NewRunner Creates a new Runner
func NewRunner(name, emailAddress string) (*Runner, error) {

	//validate the name
	if name == "" {
		return nil, ErrRunnerNameCannotBeEmpty
	}

	//validate the email address
	email, err := newEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	//create a new UUID for the runner
	id := uuid.New()

	return &Runner{
		id:           id,
		name:         name,
		emailAddress: email,
		createdAt:    time.Now().UTC(),
	}, nil
}

// LoadRunner Loads an existing Runner
func LoadRunner(id uuid.UUID, name, emailAddress string, createdAt time.Time) (*Runner, error) {

	//validate the name
	if name == "" {
		return nil, ErrRunnerNameCannotBeEmpty
	}

	//validate the email address
	email, err := newEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	return &Runner{
		id:           id,
		name:         name,
		emailAddress: email,
		createdAt:    createdAt,
	}, nil
}

// Rename Runner
func (r *Runner) Rename(name string) error {
	if name == "" {
		return ErrRunnerNameCannotBeEmpty
	}
	r.name = name
	return nil
}

// ID Returns the ID of the runner
func (r *Runner) ID() uuid.UUID {
	return r.id
}

// Name Returns the name of the runner
func (r *Runner) Name() string {
	return r.name
}

// EmailAddress Returns the email address of the runner
func (r *Runner) EmailAddress() string {
	return r.emailAddress.String()
}

// CreatedAt Returns the creation date of the runner
func (r *Runner) CreatedAt() any {
	return r.createdAt
}
