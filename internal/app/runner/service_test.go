package runner

import (
	"errors"
	"testing"

	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
	"github.com/stretchr/testify/mock"
)

func TestCreateRunner(t *testing.T) {
	tests := []struct {
		name             string
		runnerName       string
		email            string
		repoErr          error
		notificationErr  error
		wantErr          error
		mockRepo         *MockRepository
		mockNotification *MockNotificationService
	}{
		{
			name:            "Valid data",
			runnerName:      "John Doe",
			email:           "john.doe@example.com",
			repoErr:         nil,
			notificationErr: nil,
			wantErr:         nil,
			mockRepo: func() *MockRepository {
				mockRepo := new(MockRepository)
				mockRepo.On("Add", mock.Anything).Return(nil)
				return mockRepo
			}(),
			mockNotification: func() *MockNotificationService {
				mockNotificationService := new(MockNotificationService)
				mockNotificationService.
					On("SendNotification", "john.doe@example.com", "Welcome to the race tracker service!").
					Return(nil)
				return mockNotificationService
			}(),
		},
		{
			name:            "Empty name",
			runnerName:      "",
			email:           "john.doe@example.com",
			repoErr:         nil,
			notificationErr: nil,
			wantErr:         runner.ErrRunnerNameCannotBeEmpty,
			mockRepo: func() *MockRepository {
				mockRepo := new(MockRepository)
				return mockRepo
			}(),
			mockNotification: func() *MockNotificationService {
				mockNotificationService := new(MockNotificationService)
				return mockNotificationService
			}(),
		},
		{
			name:            "Invalid email",
			runnerName:      "John Doe",
			email:           "invalid-email",
			repoErr:         nil,
			notificationErr: nil,
			wantErr:         runner.ErrInvalidEmail,
			mockRepo: func() *MockRepository {
				mockRepo := new(MockRepository)
				mockRepo.On("Add", mock.Anything).Return(nil).
					Times(0)
				return mockRepo
			}(),
			mockNotification: func() *MockNotificationService {
				mockNotificationService := new(MockNotificationService)
				mockNotificationService.
					On("SendNotification", "john.doe@example.com", "Welcome to the race tracker service!").
					Times(0)
				return mockNotificationService
			}(),
		},
		{
			name:            "Repository error",
			runnerName:      "John Doe",
			email:           "john.doe@example.com",
			repoErr:         errors.New("repository error"),
			notificationErr: nil,
			wantErr:         errors.New("repository error"),
			mockRepo: func() *MockRepository {
				mockRepo := new(MockRepository)
				mockRepo.On("Add", mock.Anything).Return(errors.New("repository error"))
				return mockRepo
			}(),
			mockNotification: func() *MockNotificationService {
				mockNotificationService := new(MockNotificationService)
				mockNotificationService.
					On("SendNotification", "john.doe@example.com", "Welcome to the race tracker service!").
					Panic("SendNotification should not be called")
				return mockNotificationService
			}(),
		},
		{
			name:            "Notification error",
			runnerName:      "John Doe",
			email:           "john.doe@example.com",
			repoErr:         nil,
			notificationErr: errors.New("notification error"),
			wantErr:         nil,
			mockRepo: func() *MockRepository {
				mockRepo := new(MockRepository)
				mockRepo.On("Add", mock.Anything).Return(nil)
				return mockRepo
			}(),
			mockNotification: func() *MockNotificationService {
				mockNotificationService := new(MockNotificationService)
				mockNotificationService.
					On("SendNotification", "john.doe@example.com", "Welcome to the race tracker service!").
					Return(errors.New("notification error"))
				return mockNotificationService
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			service := NewService(tt.mockRepo, tt.mockNotification)
			_, err := service.CreateRunner(tt.runnerName, tt.email)

			if (err != nil) && (tt.wantErr == nil || err.Error() != tt.wantErr.Error()) {
				t.Errorf("CreateRunner() error = %v, wantErr %v", err, tt.wantErr)
			}

			tt.mockRepo.AssertExpectations(t)
			tt.mockNotification.AssertExpectations(t)
		})
	}
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Add(r runner.Runner) error {
	args := m.Called(r)
	return args.Error(0)
}

type MockNotificationService struct {
	mock.Mock
}

func (m *MockNotificationService) SendNotification(emailAddress string, message string) error {
	args := m.Called(emailAddress, message)
	return args.Error(0)
}
