package race

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Race represents a race event (e.g., marathon, trail run, climbing competition)
type Race struct {
	id            uuid.UUID
	name          string
	location      string
	date          time.Time
	distanceKm    float64
	elevationGain float64
}

// NewRace creates a new Race entity and validates the input
func NewRace(name, location string, date time.Time, distanceKm, elevationGain float64) (Race, error) {
	if name == "" {
		return Race{}, fmt.Errorf("name cannot be empty")
	}
	if location == "" {
		return Race{}, fmt.Errorf("location cannot be empty")
	}
	if distanceKm <= 0 {
		return Race{}, fmt.Errorf("distanceKm must be greater than 0")
	}
	if elevationGain < 0 {
		return Race{}, fmt.Errorf("elevationGain cannot be negative")
	}

	return Race{
		id:            uuid.New(),
		name:          name,
		location:      location,
		date:          date,
		distanceKm:    distanceKm,
		elevationGain: elevationGain,
	}, nil
}

// ID returns the race ID
func (r Race) ID() uuid.UUID {
	return r.id
}

// Name returns the race name
func (r Race) Name() string {
	return r.name
}

// Location returns the race location
func (r Race) Location() string {
	return r.location
}

// Date returns the race date
func (r Race) Date() time.Time {
	return r.date
}

// DistanceKm returns the race distance in kilometers
func (r Race) DistanceKm() float64 {
	return r.distanceKm
}

// ElevationGain returns the race elevation gain
func (r Race) ElevationGain() float64 {
	return r.elevationGain
}
