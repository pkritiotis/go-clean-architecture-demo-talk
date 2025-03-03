package race

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// RaceLog represents a logged race result for a participant
type RaceLog struct {
	id           uuid.UUID
	runnerID     uuid.UUID
	raceID       uuid.UUID
	finishTime   time.Duration
	pace         float64 // min/km
	heartRateAvg int
	notes        string
	loggedAt     time.Time
}

// NewRaceLog creates a new RaceLog entity and validates the input
func NewRaceLog(runnerID, raceID uuid.UUID, finishTime time.Duration, pace float64, heartRateAvg int, notes string) (RaceLog, error) {
	if runnerID == uuid.Nil {
		return RaceLog{}, fmt.Errorf("runnerID cannot be empty")
	}
	if raceID == uuid.Nil {
		return RaceLog{}, fmt.Errorf("raceID cannot be empty")
	}
	if finishTime <= 0 {
		return RaceLog{}, fmt.Errorf("finishTime must be greater than 0")
	}
	if pace <= 0 {
		return RaceLog{}, fmt.Errorf("pace must be greater than 0")
	}
	if heartRateAvg < 0 {
		return RaceLog{}, fmt.Errorf("heartRateAvg cannot be negative")
	}

	return RaceLog{
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
func (r RaceLog) ID() uuid.UUID {
	return r.id
}

// RunnerID returns the runner ID
func (r RaceLog) RunnerID() uuid.UUID {
	return r.runnerID
}

// RaceID returns the race ID
func (r RaceLog) RaceID() uuid.UUID {
	return r.raceID
}

// FinishTime returns the finish time
func (r RaceLog) FinishTime() time.Duration {
	return r.finishTime
}

// Pace returns the pace
func (r RaceLog) Pace() float64 {
	return r.pace
}

// HeartRateAvg returns the average heart rate
func (r RaceLog) HeartRateAvg() int {
	return r.heartRateAvg
}

// Notes returns the notes
func (r RaceLog) Notes() string {
	return r.notes
}

// LoggedAt returns the logged at time
func (r RaceLog) LoggedAt() time.Time {
	return r.loggedAt
}
