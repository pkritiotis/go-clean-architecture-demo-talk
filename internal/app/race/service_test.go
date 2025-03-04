package race

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/race"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type mockRaceRepository struct {
	mock.Mock
}

func (m *mockRaceRepository) SaveRace(race race.Race) error {
	args := m.Called(race)
	return args.Error(0)
}

func (m *mockRaceRepository) GetRace(raceID uuid.UUID) (race.Race, error) {
	args := m.Called(raceID)
	return args.Get(0).(race.Race), args.Error(1)
}

func (m *mockRaceRepository) SaveRaceResult(raceLog race.Result) error {
	args := m.Called(raceLog)
	return args.Error(0)
}

func (m *mockRaceRepository) GetRaceResults(runnerID uuid.UUID) ([]race.Result, error) {
	args := m.Called(runnerID)
	return args.Get(0).([]race.Result), args.Error(1)
}

func TestService_LogRace(t *testing.T) {
	mockRepo := new(mockRaceRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name       string
		runnerID   uuid.UUID
		raceID     uuid.UUID
		finishTime time.Duration
		avgHR      int
		notes      string
		mockSetup  func()
		wantErr    error
	}{
		{
			name:       "valid input",
			runnerID:   uuid.New(),
			raceID:     uuid.New(),
			finishTime: 30 * time.Minute,
			avgHR:      150,
			notes:      "Good race",
			mockSetup: func() {
				r, _ := race.NewRace("a", "l", time.Now(), 1.0, 1.0)
				mockRepo.On("GetRace", mock.Anything).Return(r, nil)
				mockRepo.On("SaveRaceResult", mock.Anything).Return(nil)
			},
			wantErr: nil,
		},
		{
			name:       "empty RunnerID",
			runnerID:   uuid.Nil,
			raceID:     uuid.New(),
			finishTime: 30 * time.Minute,
			avgHR:      150,
			notes:      "Good race",
			mockSetup:  func() {},
			wantErr:    ErrEmptyRunnerID,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			_, err := service.AddResult(tt.runnerID, tt.raceID, tt.finishTime, tt.avgHR, tt.notes)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestService_GetRaceLogs(t *testing.T) {
	mockRepo := new(mockRaceRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name      string
		runnerID  uuid.UUID
		mockSetup func()
		wantErr   error
	}{
		{
			name:     "valid input",
			runnerID: uuid.New(),
			mockSetup: func() {
				mockRepo.On("GetRaceResults", mock.Anything).Return([]race.Result{}, nil)
			},
			wantErr: nil,
		},
		{
			name:      "empty RunnerID",
			runnerID:  uuid.Nil,
			mockSetup: func() {},
			wantErr:   ErrEmptyRunnerID,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			_, err := service.GetResults(tt.runnerID)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
