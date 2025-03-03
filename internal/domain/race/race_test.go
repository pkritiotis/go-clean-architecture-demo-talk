package race

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRace(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name          string
		raceName      string
		location      string
		date          time.Time
		distanceKm    float64
		elevationGain float64
		expectError   bool
		errorMessage  string
	}{
		{
			name:          "Valid race",
			raceName:      "Marathon",
			location:      "Athens",
			date:          now,
			distanceKm:    42.195,
			elevationGain: 100,
			expectError:   false,
		},
		{
			name:          "Empty name",
			raceName:      "",
			location:      "Athens",
			date:          now,
			distanceKm:    42.195,
			elevationGain: 100,
			expectError:   true,
			errorMessage:  "name cannot be empty",
		},
		{
			name:          "Empty location",
			raceName:      "Marathon",
			location:      "",
			date:          now,
			distanceKm:    42.195,
			elevationGain: 100,
			expectError:   true,
			errorMessage:  "location cannot be empty",
		},
		{
			name:          "Zero distance",
			raceName:      "Marathon",
			location:      "Athens",
			date:          now,
			distanceKm:    0,
			elevationGain: 100,
			expectError:   true,
			errorMessage:  "distanceKm must be greater than 0",
		},
		{
			name:          "Negative distance",
			raceName:      "Marathon",
			location:      "Athens",
			date:          now,
			distanceKm:    -10,
			elevationGain: 100,
			expectError:   true,
			errorMessage:  "distanceKm must be greater than 0",
		},
		{
			name:          "Negative elevation",
			raceName:      "Marathon",
			location:      "Athens",
			date:          now,
			distanceKm:    42.195,
			elevationGain: -100,
			expectError:   true,
			errorMessage:  "elevationGain cannot be negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			race, err := NewRace(tt.raceName, tt.location, tt.date, tt.distanceKm, tt.elevationGain)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMessage, err.Error())
				assert.Equal(t, Race{}, race)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, uuid.Nil, race.ID())
				assert.Equal(t, tt.raceName, race.Name())
				assert.Equal(t, tt.location, race.Location())
				assert.Equal(t, tt.date, race.Date())
				assert.Equal(t, tt.distanceKm, race.DistanceKm())
				assert.Equal(t, tt.elevationGain, race.ElevationGain())
			}
		})
	}
}

func TestRaceGetters(t *testing.T) {
	now := time.Now()
	id := uuid.New()
	race := Race{
		id:            id,
		name:          "Test Race",
		location:      "Test Location",
		date:          now,
		distanceKm:    30.5,
		elevationGain: 500,
	}

	assert.Equal(t, id, race.ID())
	assert.Equal(t, "Test Race", race.Name())
	assert.Equal(t, "Test Location", race.Location())
	assert.Equal(t, now, race.Date())
	assert.Equal(t, 30.5, race.DistanceKm())
	assert.Equal(t, 500.0, race.ElevationGain())
}
