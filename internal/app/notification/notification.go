package notification

// Notification provides a struct to send messages via the Service
type Notification struct {
	EmailAddress string
	Subject      string
	Message      string
}

// Service sends Notification
type Service interface {
	Notify(notification Notification) error
}
