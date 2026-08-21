package userpg

import (
	"log/slog"
	"testing"
	"time"

	"uuid"

	"github.com/Masterminds/squirrel"
	"github.com/chishkin-afk/interval-learning-algo/backend/internal/modules/auth/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestRepo() *userRepository {
	return &userRepository{
		log: slog.Default(),
		sb:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func TestBuildInsert(t *testing.T) {
	repo := newTestRepo()
	now := time.Now().UTC()
	id := uuid.New()

	record := &userRecord{
		ID:           id,
		Username:     "testuser",
		Email:        "test@example.com",
		PasswordHash: "hashed_password",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	query, args, err := repo.buildInsert(record)

	require.NoError(t, err)
	assert.Contains(t, query, "INSERT INTO users")
	assert.Contains(t, query, "(id,username,email,password_hash,created_at,updated_at)")
	assert.Contains(t, query, "VALUES ($1,$2,$3,$4,$5,$6)")

	require.Len(t, args, 6)
	assert.Equal(t, id, args[0])
	assert.Equal(t, "testuser", args[1])
	assert.Equal(t, "test@example.com", args[2])
}

func TestBuildSelectCond(t *testing.T) {
	repo := newTestRepo()

	t.Run("by UUID", func(t *testing.T) {
		id := uuid.New()
		query, args, err := repo.buildSelectCond(id)

		require.NoError(t, err)
		assert.Contains(t, query, "SELECT")
		assert.Contains(t, query, "FROM users")
		assert.Contains(t, query, "WHERE id = $1")
		assert.Contains(t, query, "LIMIT 1")

		require.Len(t, args, 1)
		assert.Equal(t, id, args[0])
	})

	t.Run("by Email", func(t *testing.T) {
		email := user.Email("test@example.com")
		query, args, err := repo.buildSelectCond(email)

		require.NoError(t, err)
		assert.Contains(t, query, "WHERE email = $1")
		assert.Contains(t, query, "LIMIT 1")

		require.Len(t, args, 1)
		assert.Equal(t, email, args[0])
	})
}

func TestBuildSelectForUpdate(t *testing.T) {
	repo := newTestRepo()
	id := uuid.New()

	query, args, err := repo.buildSelectForUpdate(id)

	require.NoError(t, err)
	assert.Contains(t, query, "SELECT")
	assert.Contains(t, query, "WHERE id = $1")
	assert.Contains(t, query, "FOR UPDATE")

	require.Len(t, args, 1)
	assert.Equal(t, id, args[0])
}

func TestBuildUpdate(t *testing.T) {
	repo := newTestRepo()
	now := time.Now().UTC()

	record := &userRecord{
		Username:  "updated_user",
		UpdatedAt: now,
	}

	query, args, err := repo.buildUpdate(record)

	require.NoError(t, err)
	assert.Contains(t, query, "UPDATE users")
	assert.Contains(t, query, "username = $")
	assert.Contains(t, query, "updated_at = $")

	require.Len(t, args, 2)
}

func TestBuildDelete(t *testing.T) {
	repo := newTestRepo()
	id := uuid.New()

	query, args, err := repo.buildDelete(id)

	require.NoError(t, err)
	assert.Contains(t, query, "DELETE FROM users")
	assert.Contains(t, query, "WHERE id = $1")

	require.Len(t, args, 1)
	assert.Equal(t, id, args[0])
}
