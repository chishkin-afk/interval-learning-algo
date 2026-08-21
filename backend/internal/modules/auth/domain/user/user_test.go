package user

import (
	"errors"
	"testing"
	"time"

	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser_Success(t *testing.T) {
	username := "username"
	email := Email("mail@example.com")
	password := "password"

	now := time.Now().UTC()
	user, err := New(
		username,
		email,
		password,
	)

	require.NoError(t, err)
	require.NotEmpty(t, user.ID())

	assert.Equal(t, user.Username(), username)
	assert.Equal(t, user.Email(), email)
	assert.True(t, user.PasswordHash().Compare(password))
	assert.WithinDuration(t, user.CreatedAt(), now, 100*time.Millisecond)
	assert.WithinDuration(t, user.UpdatedAt(), now, 100*time.Millisecond)
}

func TestNewUser_Invalid(t *testing.T) {
	empties := make([]rune, 256)
	type data struct {
		username string
		email    Email
		password string
	}

	testCases := []struct {
		name     string
		input    *data
		expected error
	}{
		{
			name: "empty_username",
			input: &data{
				username: "",
				email:    Email("mail@example.com"),
				password: "password",
			},
			expected: ErrInvalidUsername,
		},
		{
			name: "too_long_username",
			input: &data{
				username: string(empties),
				email:    Email("mail@example.com"),
				password: "password",
			},
			expected: ErrInvalidUsername,
		},
		{
			name: "invalid_email",
			input: &data{
				username: "username",
				email:    Email(""),
				password: "password",
			},
			expected: ErrInvalidEmail,
		},
		{
			name: "invalid_password",
			input: &data{
				username: "username",
				email:    Email("mail@example.com"),
				password: "",
			},
			expected: ErrInvalidPassword,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := New(
				tc.input.username,
				tc.input.email,
				tc.input.password,
			)

			require.Error(t, err)
			if !errors.Is(err, tc.expected) {
				t.Errorf("want %s, got %s", tc.expected.Error(), err.Error())
			}
		})
	}
}

func TestFormUser(t *testing.T) {
	id := uuid.MustParse("a5704f3a-1cde-44cd-af4e-798ecff8aee0")
	username := "username"
	email := Email("mail@example.com")
	passwordHash := PasswordHash("some_hash_123")

	now := time.Now()
	createdAt := now
	updatedAt := now

	user := From(
		id,
		username,
		email,
		passwordHash,
		createdAt,
		updatedAt,
	)

	require.NotNil(t, user)
	assert.Equal(t, user.ID(), id)
	assert.Equal(t, user.Username(), username)
	assert.Equal(t, user.Email(), email)
	assert.Equal(t, user.PasswordHash(), passwordHash)
	assert.Equal(t, user.CreatedAt(), createdAt)
	assert.Equal(t, user.UpdatedAt(), updatedAt)
}
