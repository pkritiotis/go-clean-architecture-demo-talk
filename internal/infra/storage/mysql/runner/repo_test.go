//go:build integration

package runner

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	dsn = "user:password@tcp(localhost:3306)/dbname"
)

func TestRepo_GetByID(t *testing.T) {
	db, err := sql.Open("mysql", dsn)
	require.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)
	id := uuid.New()
	createdAt := time.Now()

	_, err = db.Exec("INSERT INTO runners (id, name, email_address, created_at) VALUES (?, ?, ?, ?)", id, "John Doe", "john.doe@example.com", createdAt)
	require.NoError(t, err)

	r, err := repo.GetByID(id)
	require.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, id, r.ID())
	assert.Equal(t, "John Doe", r.Name())
	assert.Equal(t, "john.doe@example.com", r.EmailAddress())
	assert.Equal(t, createdAt, r.CreatedAt())
}

func TestRepo_GetAll(t *testing.T) {
	db, err := sql.Open("mysql", dsn)
	require.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)
	id1 := uuid.New()
	id2 := uuid.New()
	createdAt := time.Now()

	_, err = db.Exec("INSERT INTO runners (id, name, email_address, created_at) VALUES (?, ?, ?, ?)", id1, "John Doe", "john.doe@example.com", createdAt)
	require.NoError(t, err)
	_, err = db.Exec("INSERT INTO runners (id, name, email_address, created_at) VALUES (?, ?, ?, ?)", id2, "Jane Doe", "jane.doe@example.com", createdAt)
	require.NoError(t, err)

	runners, err := repo.GetAll()
	require.NoError(t, err)
	assert.Len(t, runners, 2)
}

func TestRepo_Add(t *testing.T) {
	db, err := sql.Open("mysql", dsn)
	require.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)
	r, err := runner.NewRunner("John Doe", "john.doe@example.com")
	require.NoError(t, err)

	err = repo.Add(r)
	require.NoError(t, err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM runners WHERE id = ?", r.ID()).Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count)

	err = repo.Add(r)
	require.NoError(t, err)
}

func TestRepo_Update(t *testing.T) {
	db, err := sql.Open("mysql", dsn)
	require.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)
	r, err := runner.NewRunner("John Doe", "john.doe@example.com")
	require.NoError(t, err)
	err = repo.Add(r)
	require.NoError(t, err)

	r.Rename("John Smith")
	err = repo.Update(r)
	require.NoError(t, err)

	var name string
	err = db.QueryRow("SELECT name FROM runners WHERE id = ?", r.ID()).Scan(&name)
	require.NoError(t, err)
	assert.Equal(t, "John Smith", name)

	err = repo.Update(r)
	require.NoError(t, err)
}

func TestRepo_Delete(t *testing.T) {
	db, err := sql.Open("mysql", dsn)
	require.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)
	r, err := runner.NewRunner("John Doe", "john.doe@example.com")
	require.NoError(t, err)
	err = repo.Add(r)
	require.NoError(t, err)
	err = repo.Delete(r.ID())
	require.NoError(t, err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM runners WHERE id = ?", r.ID()).Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 0, count)

	err = repo.Delete(id)
	require.NoError(t, err)
}
