// Package runner provides an app-level service for runners use-cases
package runner

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/app/notification"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
)

// Service provides runner operations.
type Service struct {
	repo                repository
	notificationService notification.Service
}

type repository interface {
	Add(runner runner.Runner) error
}

// NewService creates a new runner service.
func NewService(repo repository, notificationService notification.Service) Service {
	return Service{repo: repo, notificationService: notificationService}
}

// CreateRunner creates a new runner.
func (s Service) CreateRunner(name, email string) (uuid.UUID, error) {

	runner, err := runner.NewRunner(name, email)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = s.repo.Add(runner)
	if err != nil {
		return uuid.UUID{}, err
	}

	_ = s.notificationService.Notify(
		notification.Notification{
			EmailAddress: runner.EmailAddress(),
			Subject:      fmt.Sprintf("Welcome %s", runner.Name()),
			Message:      "Welcome to the race tracker service!",
		},
	)

	return runner.ID(), nil
}
