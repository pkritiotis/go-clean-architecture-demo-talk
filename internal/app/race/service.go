// Package race contains the service providing the use cases for race tracking functionality
package race

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/race"
	"time"
)

// Error variables for input validation
var (
	ErrEmptyRunnerID     = errors.New("runner ID cannot be empty")
	ErrEmptyRaceID       = errors.New("race ID cannot be empty")
	ErrInvalidFinishTime = errors.New("finish time must be greater than zero")
	ErrInvalidAvgHR      = errors.New("average heart rate must be positive")
)

// Service implements the raceTracker interface
type Service struct {
	repo race.Repository
}

// NewService creates a new Service with the given repository
func NewService(repo race.Repository) Service {
	return Service{repo: repo}
}

// AddResult logs race data for a participant
func (s Service) AddResult(runnerID, raceID uuid.UUID, finishTime time.Duration, avgHR int, notes string) (uuid.UUID, error) {

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

	// Get race details to calculate Pace
	raceDetails, err := s.repo.GetRace(raceID)
	if err != nil {
		return uuid.Nil, err
	}

	// Calculate Pace (minutes per km)
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

// ResultItem represents
type ResultItem struct {
	ID           uuid.UUID
	RunnerID     uuid.UUID
	RaceID       uuid.UUID
	FinishTime   time.Duration
	Pace         float64 // min/km
	HeartRateAvg int
	Notes        string
}

// GetRaceResults retrieves race logs for a participant
func (s Service) GetResults(runnerID uuid.UUID) ([]ResultItem, error) {
	if runnerID == uuid.Nil {
		return nil, ErrEmptyRunnerID
	}

	res, err := s.repo.GetRaceResults(runnerID)
	if err != nil {
		return nil, err
	}

	results := make([]ResultItem, len(res))
	for i, r := range res {
		results[i] = ResultItem{
			ID:           r.ID(),
			RunnerID:     r.RunnerID(),
			RaceID:       r.RaceID(),
			FinishTime:   r.FinishTime(),
			Pace:         r.Pace(),
			HeartRateAvg: r.HeartRateAvg(),
			Notes:        r.Notes(),
		}
	}

	return results, nil
}

func (s Service) CreateRace(name, location string, date time.Time, distanceKm, elevationGain float64) (uuid.UUID, error) {
	if name == "" {
		return uuid.Nil, errors.New("race name cannot be empty")
	}
	if location == "" {
		return uuid.Nil, errors.New("race location cannot be empty")
	}
	if distanceKm <= 0 {
		return uuid.Nil, errors.New("distance must be greater than zero")
	}
	if elevationGain < 0 {
		return uuid.Nil, errors.New("elevation gain cannot be negative")
	}

	race, _ := race.NewRace(name, location, date, distanceKm, elevationGain)

	err := s.repo.SaveRace(race)
	if err != nil {
		return uuid.Nil, err
	}

	return race.ID(), nil
}
