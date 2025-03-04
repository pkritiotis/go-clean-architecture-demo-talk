package race

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Result represents a race record for a participant
type Result struct {
	id           uuid.UUID
	runnerID     uuid.UUID
	raceID       uuid.UUID
	finishTime   time.Duration
	pace         float64 // min/km
	heartRateAvg int
	notes        string
	loggedAt     time.Time
}

// NewRecord creates a new Result entity and validates the input
func NewRecord(runnerID, raceID uuid.UUID, finishTime time.Duration, pace float64, heartRateAvg int, notes string) (Result, error) {
	if runnerID == uuid.Nil {
		return Result{}, fmt.Errorf("runnerID cannot be empty")
	}
	if raceID == uuid.Nil {
		return Result{}, fmt.Errorf("raceID cannot be empty")
	}
	if finishTime <= 0 {
		return Result{}, fmt.Errorf("finishTime must be greater than 0")
	}
	if pace <= 0 {
		return Result{}, fmt.Errorf("pace must be greater than 0")
	}
	if heartRateAvg < 0 {
		return Result{}, fmt.Errorf("heartRateAvg cannot be negative")
	}

	return Result{
		id:           uuid.New(),
		runnerID:     runnerID,
		raceID:       raceID,
		finishTime:   finishTime,
		pace:         pace,
		heartRateAvg: heartRateAvg,
		notes:        notes,
		loggedAt:     time.Now(),
	}, nil
}

// ID returns the race log ID
func (r Result) ID() uuid.UUID {
	return r.id
}

// RunnerID returns the runner ID
func (r Result) RunnerID() uuid.UUID {
	return r.runnerID
}

// RaceID returns the race ID
func (r Result) RaceID() uuid.UUID {
	return r.raceID
}

// FinishTime returns the finish time
func (r Result) FinishTime() time.Duration {
	return r.finishTime
}

// Pace returns the pace
func (r Result) Pace() float64 {
	return r.pace
}

// HeartRateAvg returns the average heart rate
func (r Result) HeartRateAvg() int {
	return r.heartRateAvg
}

// Notes returns the notes
func (r Result) Notes() string {
	return r.notes
}

// LoggedAt returns the logged at time
func (r Result) LoggedAt() time.Time {
	return r.loggedAt
}
