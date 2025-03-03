package racetracker

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/app/racetracker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_CreateRace(t *testing.T) {
	testDate := time.Date(2023, 5, 1, 10, 0, 0, 0, time.UTC)

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		mockSetup      func(m *mockRaceTrackerService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful race creation",
			requestBody: map[string]interface{}{
				"name":           "Marathon 2023",
				"location":       "Berlin",
				"date":           testDate,
				"distance_km":    42.195,
				"elevation_gain": 350.5,
			},
			mockSetup: func(m *mockRaceTrackerService) {
				expectedID := uuid.New()
				m.On("CreateRace", "Marathon 2023", "Berlin", testDate, 42.195, 350.5).Return(expectedID, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   uuid.New().String(), // Will be replaced in test with actual mock return
		},
		{
			name: "service error",
			requestBody: map[string]interface{}{
				"name":           "Marathon 2023",
				"location":       "Berlin",
				"date":           testDate,
				"distance_km":    42.195,
				"elevation_gain": 350.5,
			},
			mockSetup: func(m *mockRaceTrackerService) {
				m.On("CreateRace", "Marathon 2023", "Berlin", testDate, 42.195, 350.5).Return(uuid.UUID{}, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "service error",
		},
		{
			name:        "invalid request body",
			requestBody: map[string]interface{}{},
			mockSetup: func(m *mockRaceTrackerService) {
				m.On("CreateRace", "", "", time.Time{}, 0.0, 0.0).Return(uuid.UUID{}, racetracker.ErrEmptyRaceID)
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   racetracker.ErrEmptyRaceID.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mockRaceTrackerService)
			tt.mockSetup(mockService)

			handler := NewHandler(mockService)

			bodyJSON, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/races", bytes.NewBuffer(bodyJSON))
			w := httptest.NewRecorder()

			handler.CreateRace(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.name == "successful race creation" {
				// Special case for successful UUID response
				assert.NotEmpty(t, w.Body.String())
				_, err := uuid.Parse(w.Body.String())
				assert.NoError(t, err)
			} else if tt.expectedBody != "" {
				assert.Contains(t, w.Body.String(), tt.expectedBody)
			}
		})
	}
}

func TestHandler_AddResult(t *testing.T) {
	validRunnerID := uuid.New()
	validRaceID := uuid.New()

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		mockSetup      func(m *mockRaceTrackerService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful result addition",
			requestBody: map[string]interface{}{
				"runner_id":      validRunnerID.String(),
				"race_id":        validRaceID.String(),
				"finish_time_ms": int64(7200000), // 2 hours
				"pace":           5.3,
				"heart_rate_avg": 155,
				"notes":          "Great race",
			},
			mockSetup: func(m *mockRaceTrackerService) {
				expectedID := uuid.New()
				m.On("AddResult", validRunnerID, validRaceID, 2*time.Hour, 5.3, 155, "Great race").Return(expectedID, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   uuid.New().String(), // Will be replaced in test with actual mock return
		},
		{
			name: "service error",
			requestBody: map[string]interface{}{
				"runner_id":      validRunnerID.String(),
				"race_id":        validRaceID.String(),
				"finish_time_ms": int64(7200000),
				"pace":           5.3,
				"heart_rate_avg": 155,
				"notes":          "Great race",
			},
			mockSetup: func(m *mockRaceTrackerService) {
				m.On("AddResult", validRunnerID, validRaceID, 2*time.Hour, 5.3, 155, "Great race").Return(uuid.UUID{}, errors.New("service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "service error",
		},
		{
			name: "invalid runner ID",
			requestBody: map[string]interface{}{
				"runner_id":      "not-a-uuid",
				"race_id":        validRaceID.String(),
				"finish_time_ms": int64(7200000),
				"pace":           5.3,
				"heart_rate_avg": 155,
				"notes":          "Great race",
			},
			mockSetup:      func(m *mockRaceTrackerService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid runner ID format",
		},
		{
			name: "invalid race ID",
			requestBody: map[string]interface{}{
				"runner_id":      validRunnerID.String(),
				"race_id":        "not-a-uuid",
				"finish_time_ms": int64(7200000),
				"pace":           5.3,
				"heart_rate_avg": 155,
				"notes":          "Great race",
			},
			mockSetup:      func(m *mockRaceTrackerService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid race ID format",
		},
		{
			name:           "invalid request body",
			requestBody:    map[string]interface{}{},
			mockSetup:      func(m *mockRaceTrackerService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mockRaceTrackerService)
			tt.mockSetup(mockService)

			handler := NewHandler(mockService)

			bodyJSON, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/results", bytes.NewBuffer(bodyJSON))
			w := httptest.NewRecorder()

			handler.AddResult(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.name == "successful result addition" {
				// Special case for successful UUID response
				assert.NotEmpty(t, w.Body.String())
				_, err := uuid.Parse(w.Body.String())
				assert.NoError(t, err)
			} else if tt.expectedBody != "" {
				assert.Contains(t, w.Body.String(), tt.expectedBody)
			}
		})
	}
}

type mockRaceTrackerService struct {
	mock.Mock
}

func (m *mockRaceTrackerService) CreateRace(name, location string, date time.Time, distanceKm, elevationGain float64) (uuid.UUID, error) {
	args := m.Called(name, location, date, distanceKm, elevationGain)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (m *mockRaceTrackerService) AddResult(runnerID, raceID uuid.UUID, finishTime time.Duration, pace float64, heartRateAvg int, notes string) (uuid.UUID, error) {
	args := m.Called(runnerID, raceID, finishTime, pace, heartRateAvg, notes)
	return args.Get(0).(uuid.UUID), args.Error(1)
}
