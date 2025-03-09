package runner

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
	"time"
)

// Repo Implements the Repository Interface to provide a MySQL storage provider
type Repo struct {
	db *sql.DB
}

// NewRepository Constructor
func NewRepository(db *sql.DB) Repo {
	return Repo{db}
}

// GetByID Returns the runner with the provided id
func (m Repo) GetByID(id uuid.UUID) (*runner.Runner, error) {
	var r struct {
		id           uuid.UUID
		name         string
		emailAddress string
		createdAt    time.Time
	}
	query := "SELECT id, name, email_address, created_at FROM runners WHERE id = ?"
	row := m.db.QueryRow(query, id)
	err := row.Scan(&r.id, &r.name, &r.emailAddress, &r.createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	domainRunner, err := runner.LoadRunner(r.id, r.name, r.emailAddress, r.createdAt)
	if err != nil {
		return nil, err
	}
	return domainRunner, nil
}

// GetAll Returns all stored runners
func (m Repo) GetAll() ([]*runner.Runner, error) {
	query := "SELECT id, name, email_address, created_at FROM runners"
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var runners []*runner.Runner
	for rows.Next() {
		var r struct {
			id           uuid.UUID
			name         string
			emailAddress string
			createdAt    time.Time
		}
		err := rows.Scan(&r.id, &r.name, &r.emailAddress, &r.createdAt)
		if err != nil {
			return nil, err
		}
		domainRunner, err := runner.LoadRunner(r.id, r.name, r.emailAddress, r.createdAt)
		if err != nil {
			return nil, err
		}
		runners = append(runners, domainRunner)
	}
	return runners, nil
}

// Add the provided runner
func (m Repo) Add(runner *runner.Runner) error {
	query := "INSERT INTO runners (id, name, email_address, created_at) VALUES (?, ?, ?, ?)"
	_, err := m.db.Exec(query, runner.ID(), runner.Name(), runner.EmailAddress(), runner.CreatedAt())
	return err
}

// Update the provided runner
func (m Repo) Update(runner *runner.Runner) error {
	query := "UPDATE runners SET name = ?, email_address = ?, created_at = ? WHERE id = ?"
	_, err := m.db.Exec(query, runner.Name(), runner.EmailAddress(), runner.CreatedAt(), runner.ID())
	return err
}

// Delete the runner with the provided id
func (m Repo) Delete(id uuid.UUID) error {
	query := "DELETE FROM runners WHERE id = ?"
	result, err := m.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("id %v not found", id.String())
	}
	return nil
}
