// Package infra contains the services of the interface adapters
package infra

import (
	"github.com/pkritiotis/go-clean-architecture-example/internal/app"
	"github.com/pkritiotis/go-clean-architecture-example/internal/app/notification"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/race"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
	"github.com/pkritiotis/go-clean-architecture-example/internal/infra/http"
	"github.com/pkritiotis/go-clean-architecture-example/internal/infra/notification/console"
	racememrepo "github.com/pkritiotis/go-clean-architecture-example/internal/infra/storage/memory/race"
	runnermemrep "github.com/pkritiotis/go-clean-architecture-example/internal/infra/storage/memory/runner"
)

// Services contains the exposed services of interface adapters
type Services struct {
	NotificationService notification.Service
	RunnerRepository    runner.Repository
	RaceRepository      race.Repository
	Server              *http.Server
}

// NewInfraProviders Instantiates the infra services
func NewInfraProviders() Services {
	return Services{
		NotificationService: console.NewNotificationService(),
		RaceRepository:      racememrepo.NewRepository(),
		RunnerRepository:    runnermemrep.NewRepository(),
	}
}

// NewHTTPServer creates a new server
func NewHTTPServer(appServices app.Services) *http.Server {
	return http.NewServer(appServices)
}
