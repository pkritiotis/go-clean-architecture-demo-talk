// Package runner contains the runner service of the domain
package runner

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	domainRunner "github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
	"net/http"
)

type runnerService interface {
	CreateRunner(name, email string) (uuid.UUID, error)
}

// Handler Runner http request service
type Handler struct {
	runnerService runnerService
}

// NewHandler Constructor
func NewHandler(service runnerService) Handler {
	return Handler{runnerService: service}
}

// CreateRunnerRequestModel represents the request model expected for Add request
type CreateRunnerRequestModel struct {
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
}

// Create Adds the provides runner
func (c Handler) Create(w http.ResponseWriter, r *http.Request) {
	var runnerToAdd CreateRunnerRequestModel
	decodeErr := json.NewDecoder(r.Body).Decode(&runnerToAdd)
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, decodeErr.Error())
		return
	}
	id, err := c.runnerService.CreateRunner(runnerToAdd.Name, runnerToAdd.EmailAddress)
	if err != nil {
		if errors.Is(err, domainRunner.ErrInvalidEmail) || errors.Is(err, domainRunner.ErrRunnerNameCannotBeEmpty) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprint(w, err.Error())
		return
	}
	w.Write([]byte(id.String()))
}
