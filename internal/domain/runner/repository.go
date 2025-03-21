package runner

import "github.com/google/uuid"

// Repository Interface for runners
type Repository interface {
	GetByID(id uuid.UUID) (*Runner, error)
	Add(runner *Runner) error
	Update(runner *Runner) error
	GetAll() ([]*Runner, error)
	Delete(id uuid.UUID) error
}
