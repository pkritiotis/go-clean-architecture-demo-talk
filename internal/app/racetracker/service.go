// Package racetracker contains the service providing the use cases for race tracking functionality
package racetracker

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/race"
	"time"
)

// RaceTracker defines the methods for logging race data
type RaceTracker interface {
	LogRaceResult(runnerID, raceID uuid.UUID, finishTime time.Duration, avgHR int, notes string) (race.Result, error)
	GetRaceResults(runnerID uuid.UUID) ([]race.Result, error)
}

type raceRepository interface {
	GetRace(raceID uuid.UUID) (race.Race, error)
	SaveRaceResult(raceLog race.Result) error
	GetRaceResults(runnerID uuid.UUID) ([]race.Result, error)
}

// Error variables for input validation
var (
	ErrEmptyRunnerID     = errors.New("runner ID cannot be empty")
	ErrEmptyRaceID       = errors.New("race ID cannot be empty")
	ErrInvalidFinishTime = errors.New("finish time must be greater than zero")
	ErrInvalidAvgHR      = errors.New("average heart rate must be positive")
)

// Service implements the RaceTracker interface
type Service struct {
	repo raceRepository
}

// NewService creates a new Service with the given repository
func NewService(repo raceRepository) *Service {
	return &Service{repo: repo}
}

// SaveRaceResult logs race data for a participant
func (s *Service) SaveRaceResult(runnerID, raceID uuid.UUID, finishTime time.Duration, avgHR int, notes string) (uuid.UUID, error) {

	// Validate inputs
	if runnerID == uuid.Nil {
		return uuid.Nil, ErrEmptyRunnerID
	}
	if raceID == uuid.Nil {
		return uuid.Nil, ErrEmptyRaceID
	}
	if finishTime <= 0 {
		return uuid.Nil, ErrInvalidFinishTime
	}
	if avgHR <= 0 {
		return uuid.Nil, ErrInvalidAvgHR
	}

	// Get race details to calculate pace
	raceDetails, err := s.repo.GetRace(raceID)
	if err != nil {
		return uuid.Nil, err
	}

	// Calculate pace (minutes per km)
	pace := float64(finishTime.Minutes()) / raceDetails.DistanceKm()

	// Create and store the race log
	raceLog, err := race.NewRecord(runnerID, raceID, finishTime, pace, avgHR, notes)
	if err != nil {
		return uuid.Nil, err
	}

	// Save the race log using the repository
	err = s.repo.SaveRaceResult(raceLog)
	if err != nil {
		return uuid.Nil, err
	}

	return raceLog.RaceID(), nil
}

// RaceResult represents
type RaceResult struct {
	id           uuid.UUID
	runnerID     uuid.UUID
	raceID       uuid.UUID
	finishTime   time.Duration
	pace         float64 // min/km
	heartRateAvg int
	notes        string
}

// GetRaceResults retrieves race logs for a participant
func (s *Service) GetRaceResults(runnerID uuid.UUID) ([]RaceResult, error) {
	if runnerID == uuid.Nil {
		return nil, ErrEmptyRunnerID
	}

	logs, err := s.repo.GetRaceResults(runnerID)
	if err != nil {
		return nil, err
	}

	results := make([]RaceResult, len(logs))
	for i, log := range logs {
		results[i] = RaceResult{
			id:           log.ID(),
			runnerID:     log.RunnerID(),
			raceID:       log.RaceID(),
			finishTime:   log.FinishTime(),
			pace:         log.Pace(),
			heartRateAvg: log.HeartRateAvg(),
			notes:        log.Notes(),
		}
	}

	return results, nil
}
