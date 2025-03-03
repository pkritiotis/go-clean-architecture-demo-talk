// Package memory implements the Repository Interface to provide an in-memory storage provider
package memory

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
)

// Repo Implements the Repository Interface to provide an in-memory storage provider
type Repo struct {
	runners map[uuid.UUID]runner.Runner
}

// NewRepo Constructor
func NewRepo() Repo {
	runners := make(map[uuid.UUID]runner.Runner)
	return Repo{runners}
}

// GetByID Returns the runner with the provided id
func (m Repo) GetByID(id uuid.UUID) (*runner.Runner, error) {
	r, ok := m.runners[id]
	if !ok {
		return nil, nil
	}
	return &r, nil
}

// GetAll Returns all stored runners
func (m Repo) GetAll() ([]runner.Runner, error) {
	keys := make([]uuid.UUID, 0)

	for key := range m.runners {
		keys = append(keys, key)
	}

	var values []runner.Runner
	for _, value := range m.runners {
		values = append(values, value)
	}
	return values, nil
}

// Add the provided runner
func (m Repo) Add(runner runner.Runner) error {
	m.runners[runner.ID()] = runner
	return nil
}

// Update the provided runner
func (m Repo) Update(runner runner.Runner) error {
	m.runners[runner.ID()] = runner
	return nil
}

// Delete the runner with the provided id
func (m Repo) Delete(id uuid.UUID) error {
	_, exists := m.runners[id]
	if !exists {
		return fmt.Errorf("id %v not found", id.String())
	}
	delete(m.runners, id)
	return nil
}
