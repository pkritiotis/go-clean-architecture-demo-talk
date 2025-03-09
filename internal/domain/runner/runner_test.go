package runner

import (
	"github.com/google/uuid"
	"testing"
	"time"
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

func TestRename(t *testing.T) {
	tests := []struct {
		name        string
		initialName string
		newName     string
		wantErr     error
	}{
		{
			name:        "Valid rename",
			initialName: "John Doe",
			newName:     "Jane Doe",
			wantErr:     nil,
		},
		{
			name:        "Empty new name",
			initialName: "John Doe",
			newName:     "",
			wantErr:     ErrRunnerNameCannotBeEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner, err := NewRunner(tt.initialName, "john.doe@example.com")
			if err != nil {
				t.Fatalf("NewRunner() error = %v", err)
			}
			err = runner.Rename(tt.newName)
			if err != tt.wantErr {
				t.Errorf("Rename() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && runner.Name() != tt.newName {
				t.Errorf("Rename() name = %v, want %v", runner.Name(), tt.newName)
			}
		})
	}
}

func TestLoadRunner(t *testing.T) {
	tests := []struct {
		name       string
		id         uuid.UUID
		runnerName string
		email      string
		createdAt  time.Time
		wantErr    error
	}{
		{
			name:       "Valid data",
			id:         uuid.New(),
			runnerName: "John Doe",
			email:      "john.doe@example.com",
			createdAt:  time.Now().UTC(),
			wantErr:    nil,
		},
		{
			name:       "Empty name",
			id:         uuid.New(),
			runnerName: "",
			email:      "john.doe@example.com",
			createdAt:  time.Now().UTC(),
			wantErr:    ErrRunnerNameCannotBeEmpty,
		},
		{
			name:       "Invalid email",
			id:         uuid.New(),
			runnerName: "John Doe",
			email:      "invalid-email",
			createdAt:  time.Now().UTC(),
			wantErr:    ErrInvalidEmail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner, err := LoadRunner(tt.id, tt.runnerName, tt.email, tt.createdAt)
			if err != tt.wantErr {
				t.Errorf("LoadRunner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if runner.ID() != tt.id {
					t.Errorf("LoadRunner() id = %v, want %v", runner.ID(), tt.id)
				}
				if runner.Name() != tt.runnerName {
					t.Errorf("LoadRunner() name = %v, want %v", runner.Name(), tt.runnerName)
				}
				if runner.EmailAddress() != tt.email {
					t.Errorf("LoadRunner() emailAddress = %v, want %v", runner.EmailAddress(), tt.email)
				}
				if runner.CreatedAt() != tt.createdAt {
					t.Errorf("LoadRunner() createdAt = %v, want %v", runner.CreatedAt(), tt.createdAt)
				}
			}
		})
	}
}
