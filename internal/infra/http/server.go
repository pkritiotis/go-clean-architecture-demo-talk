// Package http contains the implementation of the HTTP server.
package http

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/app"
	appRace "github.com/pkritiotis/go-clean-architecture-example/internal/app/race"
	"github.com/pkritiotis/go-clean-architecture-example/internal/infra/http/race"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkritiotis/go-clean-architecture-example/internal/infra/http/runner"
)

type runnerService interface {
	CreateRunner(name, email string) (uuid.UUID, error)
}

type raceService interface {
	CreateRace(name, location string, date time.Time, distanceKm, elevationGain float64) (uuid.UUID, error)
	AddResult(runnerID, raceID uuid.UUID, finishTime time.Duration, heartRateAvg int, notes string) (uuid.UUID, error)
	GetResults(runnerID uuid.UUID) ([]appRace.ResultItem, error)
}

// Server Represents the http server running for this service
type Server struct {
	runnerService runnerService
	raceService   raceService
	router        *mux.Router
}

// NewServer HTTP Server constructor
func NewServer(appServices app.Services) *Server {
	httpServer := &Server{runnerService: appServices.RunnerService, raceService: appServices.RaceService}
	httpServer.router = mux.NewRouter()
	httpServer.AddRunnerHTTPRoutes()
	httpServer.AddRaceHTTPRoutes()
	http.Handle("/", httpServer.router)

	return httpServer
}

// AddRunnerHTTPRoutes registers runner route handlers
func (httpServer *Server) AddRunnerHTTPRoutes() {
	const runnersHTTPRoutePath = "/runners"
	httpServer.router.HandleFunc(runnersHTTPRoutePath, runner.NewHandler(httpServer.runnerService).Create).Methods("POST")
}

// AddRaceHTTPRoutes registers race route handlers
func (httpServer *Server) AddRaceHTTPRoutes() {
	const racesHTTPRoutePath = "/races"
	handler := race.NewHandler(httpServer.raceService)
	httpServer.router.HandleFunc(racesHTTPRoutePath, handler.CreateRace).Methods("POST")
	httpServer.router.HandleFunc(racesHTTPRoutePath, handler.GetRaceResults).Methods("GET")
	httpServer.router.HandleFunc(racesHTTPRoutePath+"/{raceID}/results", handler.AddResult).Methods("POST")
}

// ListenAndServe Starts listening for requests
func (httpServer *Server) ListenAndServe(port string) {
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
