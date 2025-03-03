package racetracker

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/app/racetracker"
	"net/http"
	"time"
)

type raceTrackerService interface {
	CreateRace(name, location string, date time.Time, distanceKm, elevationGain float64) (uuid.UUID, error)
	AddResult(runnerID, raceID uuid.UUID, finishTime time.Duration, pace float64, heartRateAvg int, notes string) (uuid.UUID, error)
}

// Handler RaceTracker http request service
type Handler struct {
	raceTrackerService raceTrackerService
}

// NewHandler Constructor
func NewHandler(service raceTrackerService) Handler {
	return Handler{raceTrackerService: service}
}

// CreateRaceRequestModel represents the request model expected for creating a race
type CreateRaceRequestModel struct {
	Name          string    `json:"name"`
	Location      string    `json:"location"`
	Date          time.Time `json:"date"`
	DistanceKm    float64   `json:"distance_km"`
	ElevationGain float64   `json:"elevation_gain"`
}

// CreateRace handles requests to create a new race
func (h Handler) CreateRace(w http.ResponseWriter, r *http.Request) {
	var raceRequest CreateRaceRequestModel
	decodeErr := json.NewDecoder(r.Body).Decode(&raceRequest)
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, decodeErr.Error())
		return
	}

	id, err := h.raceTrackerService.CreateRace(
		raceRequest.Name,
		raceRequest.Location,
		raceRequest.Date,
		raceRequest.DistanceKm,
		raceRequest.ElevationGain,
	)

	if err != nil {
		switch err {
		case racetracker.ErrEmptyRaceID, racetracker.ErrEmptyRunnerID, racetracker.ErrInvalidAvgHR, racetracker.ErrInvalidFinishTime:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprint(w, err.Error())
		return
	}

	w.Write([]byte(id.String()))
	w.WriteHeader(http.StatusOK)
}

// AddResultRequestModel represents the request model for adding a race result
type AddResultRequestModel struct {
	RunnerID     string  `json:"runner_id"`
	RaceID       string  `json:"race_id"`
	FinishTimeMs int64   `json:"finish_time_ms"`
	Pace         float64 `json:"pace"`
	HeartRateAvg int     `json:"heart_rate_avg"`
	Notes        string  `json:"notes"`
}

// AddResult handles requests to add a new race result
func (h Handler) AddResult(w http.ResponseWriter, r *http.Request) {
	var resultRequest AddResultRequestModel
	decodeErr := json.NewDecoder(r.Body).Decode(&resultRequest)
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, decodeErr.Error())
		return
	}

	runnerID, err := uuid.Parse(resultRequest.RunnerID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid runner ID format")
		return
	}

	raceID, err := uuid.Parse(resultRequest.RaceID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid race ID format")
		return
	}

	finishTime := time.Duration(resultRequest.FinishTimeMs) * time.Millisecond

	id, err := h.raceTrackerService.AddResult(
		runnerID,
		raceID,
		finishTime,
		resultRequest.Pace,
		resultRequest.HeartRateAvg,
		resultRequest.Notes,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Write([]byte(id.String()))
	w.WriteHeader(http.StatusOK)
}
