// Package race contains the in-memory implementation of the race repository
package race

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/race"
	"sync"
)

// Repo is an in-memory implementation of the race repository
type Repo struct {
	races           map[uuid.UUID]race.Race
	raceResults     map[uuid.UUID]race.Result
	resultsByRunner map[uuid.UUID][]uuid.UUID
	mu              sync.RWMutex
}

// NewRepository creates a new in-memory race repository
func NewRepository() *Repo {
	return &Repo{
		races:           make(map[uuid.UUID]race.Race),
		raceResults:     make(map[uuid.UUID]race.Result),
		resultsByRunner: make(map[uuid.UUID][]uuid.UUID),
	}
}

// GetRace retrieves a race by ID
func (r *Repo) GetRace(raceID uuid.UUID) (race.Race, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	race, exists := r.races[raceID]
	if !exists {
		return race, fmt.Errorf("race with ID %s not found", raceID)
	}

	return race, nil
}

// SaveRace saves a race to the repository
func (r *Repo) SaveRace(race race.Race) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.races[race.ID()] = race
	return nil
}

// SaveRaceResult saves a race result to the repository
func (r *Repo) SaveRaceResult(result race.Result) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Store the race result
	r.raceResults[result.ID()] = result

	// Add to runner's results
	runnerID := result.RunnerID()
	r.resultsByRunner[runnerID] = append(r.resultsByRunner[runnerID], result.ID())

	return nil
}

// GetRaceResults gets all race results for a runner
func (r *Repo) GetRaceResults(runnerID uuid.UUID) ([]race.Result, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	resultIDs, exists := r.resultsByRunner[runnerID]
	if !exists {
		return []race.Result{}, nil
	}

	results := make([]race.Result, 0, len(resultIDs))
	for _, resultID := range resultIDs {
		result, exists := r.raceResults[resultID]
		if exists {
			results = append(results, result)
		}
	}

	return results, nil
}
