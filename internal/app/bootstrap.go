// Package app contains the bootstrap logic of the application layer
package app

import (
	"github.com/pkritiotis/go-clean-architecture-example/internal/app/notification"
	"github.com/pkritiotis/go-clean-architecture-example/internal/app/runner"
	domainRunner "github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
)

// Services contains the exposed services of the application layer
type Services struct {
	RunnerService runner.Service
}

// NewServices creates a new application services
func NewServices(runnerRepo domainRunner.Repository, notificationService notification.Service) Services {
	rs := runner.NewService(runnerRepo, notificationService)
	return Services{RunnerService: rs}
}
