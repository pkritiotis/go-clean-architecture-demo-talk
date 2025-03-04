// Package app contains the bootstrap logic of the application layer
package app

import (
	"github.com/pkritiotis/go-clean-architecture-example/internal/app/notification"
	"github.com/pkritiotis/go-clean-architecture-example/internal/app/race"
	"github.com/pkritiotis/go-clean-architecture-example/internal/app/runner"
	domainRace "github.com/pkritiotis/go-clean-architecture-example/internal/domain/race"
	domainRunner "github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
)

// Services contains the exposed services of the application layer
type Services struct {
	RunnerService runner.Service
	RaceService   race.Service
}

// NewServices creates a new application services
func NewServices(runnerRepo domainRunner.Repository, raceRepo domainRace.Repository, notificationService notification.Service) Services {
	rs := runner.NewService(runnerRepo, notificationService)
	rts := race.NewService(raceRepo)
	return Services{RunnerService: rs, RaceService: rts}
}
