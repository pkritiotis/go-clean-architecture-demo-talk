// Package race contains the domain entities for race tracking purposes
package race

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrEmptyName            = errors.New("name cannot be empty")
	ErrEmptyLocation        = errors.New("location cannot be empty")
	ErrInvalidDistanceKm    = errors.New("distanceKm must be greater than 0")
	ErrInvalidElevationGain = errors.New("elevationGain cannot be negative")
)

// Race represents a race event (e.g., marathon, half-marathon, trail run)
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
		return Race{}, ErrEmptyName
	}
	if location == "" {
		return Race{}, ErrEmptyLocation
	}
	if distanceKm <= 0 {
		return Race{}, ErrInvalidDistanceKm
	}
	if elevationGain < 0 {
		return Race{}, ErrInvalidElevationGain
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
