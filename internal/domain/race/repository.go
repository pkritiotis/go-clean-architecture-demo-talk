package race

import "github.com/google/uuid"

// Repository defines the storage interface for race
type Repository interface {
	SaveRace(Race) error
	GetRace(raceID uuid.UUID) (Race, error)
	SaveRaceResult(raceLog Result) error
	GetRaceResults(runnerID uuid.UUID) ([]Result, error)
}
