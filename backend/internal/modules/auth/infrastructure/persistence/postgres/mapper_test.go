package userpg

import (
	"testing"
	"time"

	"uuid"

	"github.com/chishkin-afk/interval-learning-algo/backend/internal/modules/auth/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToRecord(t *testing.T) {
	id := uuid.New()
	now := time.Now().UTC()

	u := user.From(
		id,
		"testuser",
		user.Email("test@example.com"),
		user.PasswordHash("super_secret_hash"),
		now,
		now,
	)

	record := toRecord(u)

	assert.Equal(t, id, record.ID)
	assert.Equal(t, "testuser", record.Username)
	assert.Equal(t, "test@example.com", record.Email)
	assert.Equal(t, "super_secret_hash", record.PasswordHash)
	assert.Equal(t, now, record.CreatedAt)
	assert.Equal(t, now, record.UpdatedAt)
}

func TestToUser(t *testing.T) {
	id := uuid.New()
	now := time.Now().UTC()

	record := &userRecord{
		ID:           id,
		Username:     "db_user",
		Email:        "db@example.com",
		PasswordHash: "db_hash_value",
		CreatedAt:    now,
		UpdatedAt:    now.Add(time.Hour),
	}

	u := toUser(record)

	require.NotNil(t, u)
	assert.Equal(t, id, u.ID())
	assert.Equal(t, "db_user", u.Username())
	assert.Equal(t, "db@example.com", string(u.Email()))
	assert.Equal(t, "db_hash_value", string(u.PasswordHash()))
	assert.Equal(t, now, u.CreatedAt())
	assert.Equal(t, now.Add(time.Hour), u.UpdatedAt())
}

func TestMappingRoundTrip(t *testing.T) {
	id := uuid.New()
	now := time.Now().UTC()

	original := user.From(id, "roundtrip", user.Email("rt@test.com"), user.PasswordHash("hash"), now, now)

	rec := toRecord(original)

	restored := toUser(rec)

	assert.Equal(t, original.ID(), restored.ID())
	assert.Equal(t, original.Username(), restored.Username())
	assert.Equal(t, string(original.Email()), string(restored.Email()))
	assert.Equal(t, string(original.PasswordHash()), string(restored.PasswordHash()))
	assert.Equal(t, original.CreatedAt(), restored.CreatedAt())
	assert.Equal(t, original.UpdatedAt(), restored.UpdatedAt())
}
