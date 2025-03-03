package race

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestRaceLogValidation(t *testing.T) {
	tests := []struct {
		name         string
		runnerID     uuid.UUID
		raceID       uuid.UUID
		finishTime   time.Duration
		pace         float64
		heartRateAvg int
		notes        string
		wantErr      string
	}{
		{
			name:         "RunnerID cannot be empty",
			runnerID:     uuid.Nil,
			raceID:       uuid.New(),
			finishTime:   time.Hour,
			pace:         5.0,
			heartRateAvg: 150,
			notes:        "Good race",
			wantErr:      "runnerID cannot be empty",
		},
		{
			name:         "RaceID cannot be empty",
			runnerID:     uuid.New(),
			raceID:       uuid.Nil,
			finishTime:   time.Hour,
			pace:         5.0,
			heartRateAvg: 150,
			notes:        "Good race",
			wantErr:      "raceID cannot be empty",
		},
		{
			name:         "FinishTime must be greater than zero",
			runnerID:     uuid.New(),
			raceID:       uuid.New(),
			finishTime:   0,
			pace:         5.0,
			heartRateAvg: 150,
			notes:        "Good race",
			wantErr:      "finishTime must be greater than 0",
		},
		{
			name:         "Pace must be greater than zero",
			runnerID:     uuid.New(),
			raceID:       uuid.New(),
			finishTime:   time.Hour,
			pace:         0,
			heartRateAvg: 150,
			notes:        "Good race",
			wantErr:      "pace must be greater than 0",
		},
		{
			name:         "HeartRateAvg cannot be negative",
			runnerID:     uuid.New(),
			raceID:       uuid.New(),
			finishTime:   time.Hour,
			pace:         5.0,
			heartRateAvg: -1,
			notes:        "Good race",
			wantErr:      "heartRateAvg cannot be negative",
		},
		{
			name:         "Valid RaceLog creation",
			runnerID:     uuid.New(),
			raceID:       uuid.New(),
			finishTime:   time.Hour,
			pace:         5.0,
			heartRateAvg: 150,
			notes:        "Good race",
			wantErr:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raceLog, err := NewRaceLog(tt.runnerID, tt.raceID, tt.finishTime, tt.pace, tt.heartRateAvg, tt.notes)
			if tt.wantErr != "" {
				if err == nil || err.Error() != tt.wantErr {
					t.Errorf("expected error '%v', got %v", tt.wantErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if raceLog.ID() == uuid.Nil {
					t.Errorf("expected valid ID, got %v", raceLog.ID())
				}
				if raceLog.RunnerID() != tt.runnerID {
					t.Errorf("expected runnerID %v, got %v", tt.runnerID, raceLog.RunnerID())
				}
				if raceLog.RaceID() != tt.raceID {
					t.Errorf("expected raceID %v, got %v", tt.raceID, raceLog.RaceID())
				}
				if raceLog.FinishTime() != tt.finishTime {
					t.Errorf("expected finishTime %v, got %v", tt.finishTime, raceLog.FinishTime())
				}
				if raceLog.Pace() != tt.pace {
					t.Errorf("expected pace %v, got %v", tt.pace, raceLog.Pace())
				}
				if raceLog.HeartRateAvg() != tt.heartRateAvg {
					t.Errorf("expected heartRateAvg %v, got %v", tt.heartRateAvg, raceLog.HeartRateAvg())
				}
				if raceLog.Notes() != tt.notes {
					t.Errorf("expected notes '%v', got %v", tt.notes, raceLog.Notes())
				}
				if raceLog.LoggedAt().IsZero() {
					t.Errorf("expected valid loggedAt, got %v", raceLog.LoggedAt())
				}
			}
		})
	}
}
