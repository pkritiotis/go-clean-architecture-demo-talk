package runner

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
)

// Service provides runner operations.
type Service struct {
	repo                repository
	notificationService notificationService
}

type repository interface {
	Add(runner runner.Runner) error
}

type notificationService interface {
	SendNotification(emailAddress string, message string) error
}

// NewService creates a new runner service.
func NewService(repo repository, notificationService notificationService) Service {
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

	err = s.notificationService.SendNotification(runner.EmailAddress(), "Welcome to the race tracker service!")
	if err != nil {
		// log the error
	}

	return runner.ID(), nil
}
