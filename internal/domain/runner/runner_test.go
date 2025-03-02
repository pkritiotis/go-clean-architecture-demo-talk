package runner

import (
	"testing"
)

func TestNewRunner(t *testing.T) {
	tests := []struct {
		name       string
		runnerName string
		email      string
		wantErr    error
	}{
		{
			name:       "Valid data",
			runnerName: "John Doe",
			email:      "john.doe@example.com",
			wantErr:    nil,
		},
		{
			name:       "Empty name",
			runnerName: "",
			email:      "john.doe@example.com",
			wantErr:    ErrRunnerNameCannotBeEmpty,
		},
		{
			name:       "Invalid email",
			runnerName: "John Doe",
			email:      "invalid-email",
			wantErr:    ErrInvalidEmail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner, err := NewRunner(tt.runnerName, tt.email)
			if err != tt.wantErr {
				t.Errorf("NewRunner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if runner.name != tt.runnerName {
					t.Errorf("NewRunner() name = %v, want %v", runner.name, tt.runnerName)
				}
				if runner.emailAddress.String() != tt.email {
					t.Errorf("NewRunner() emailAddress = %v, want %v", runner.emailAddress.String(), tt.email)
				}
			}
		})
	}
}
