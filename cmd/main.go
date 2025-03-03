// package main contains the entry point of the application.
package main

import (
	"github.com/pkritiotis/go-clean-architecture-example/internal/app"
	"github.com/pkritiotis/go-clean-architecture-example/internal/infra"
)

func main() {
	//Initialize the infrastructure providers
	infraProviders := infra.NewInfraProviders()

	//Initialize the application services using the infrastructure provider implementations
	appServices := app.NewServices(infraProviders.RunnerRepository, infraProviders.NotificationService)

	//Initialize the HTTP server that calls the application services
	infraHTTPServer := infra.NewHTTPServer(appServices)
	infraHTTPServer.ListenAndServe(":8080")
}
