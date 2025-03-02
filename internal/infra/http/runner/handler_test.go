package runner

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRunningService struct {
	Handler func(name, email string) (uuid.UUID, error)
}

func (m MockRunningService) CreateRunner(name string, email string) (uuid.UUID, error) {
	return m.Handler(name, email)
}

func TestRunnerHandler_AddRunner(t *testing.T) {
	testUUID := uuid.New()
	tests := []struct {
		name               string
		service            runnerService
		reqVars            map[string]interface{}
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should add runner successfully",
			service: MockRunningService{Handler: func(name, email string) (uuid.UUID, error) {
				if name != "test" || email != "name@example.com" {
					return uuid.UUID{}, errors.New("objects not matching")
				}
				return testUUID, nil
			}},
			reqVars: map[string]interface{}{},
			Body: CreateRunnerRequestModel{
				Name:         "test",
				EmailAddress: "name@example.com",
			},
			ResultBodyContains: testUUID.String(),
			ResultStatus:       http.StatusOK,
		},
		{
			name: "should return error",
			service: MockRunningService{Handler: func(name, email string) (uuid.UUID, error) {
				if name != "test" || email != "name@example.com" {
					return uuid.UUID{}, errors.New("objects not matching")
				}
				return uuid.UUID{}, errors.New("test error")
			}},
			reqVars: map[string]interface{}{},
			Body: CreateRunnerRequestModel{
				Name:         "test",
				EmailAddress: "name@example.com",
			},
			ResultBodyContains: errors.New("test error").Error(),
			ResultStatus:       http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewHandler(tt.service)
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(tt.Body)
			req, _ := http.NewRequest("POST", "", buf)
			rsp := httptest.NewRecorder()
			c.Create(rsp, req)
			assert.Contains(t, tt.ResultBodyContains, rsp.Body.String())
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}
