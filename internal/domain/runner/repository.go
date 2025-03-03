package runner

// Repository Interface for runners
type Repository interface {
	Add(runner Runner) error
}
