// Package http contains the implementation of the HTTP server.
package http

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkritiotis/go-clean-architecture-example/internal/infra/http/runner"
)

type runnerService interface {
	CreateRunner(name, email string) (uuid.UUID, error)
}

// Server Represents the http server running for this service
type Server struct {
	runnerService runnerService
	router        *mux.Router
}

// NewServer HTTP Server constructor
func NewServer(runnerService runnerService) *Server {
	httpServer := &Server{runnerService: runnerService}
	httpServer.router = mux.NewRouter()
	httpServer.AddRunnerHTTPRoutes()
	http.Handle("/", httpServer.router)

	return httpServer
}

// AddRunnerHTTPRoutes registers runner route handlers
func (httpServer *Server) AddRunnerHTTPRoutes() {
	const runnersHTTPRoutePath = "/runners"
	httpServer.router.HandleFunc(runnersHTTPRoutePath, runner.NewHandler(httpServer.runnerService).Create).Methods("POST")
}

// ListenAndServe Starts listening for requests
func (httpServer *Server) ListenAndServe(port string) {
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
