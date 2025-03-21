// Package console contains the console implementation of the notification service
package console

import (
	"encoding/json"
	"fmt"

	"github.com/pkritiotis/go-clean-architecture-example/internal/app/notification"
)

// NotificationService provides a console implementation of the Service
type NotificationService struct{}

// NewNotificationService constructor for NotificationService
func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

// Notify prints out the notifications in console
func (NotificationService) Notify(notification notification.Notification) error {
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	fmt.Printf("Notification Received: %v\n", string(jsonNotification))
	return nil
}
