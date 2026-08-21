package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPasswordHash_Success(t *testing.T) {
	rawPassword := "very_secret_password"
	passwordHash, err := NewPasswordHash(rawPassword)

	require.NoError(t, err)
	assert.NotEqual(t, rawPassword, passwordHash)
	assert.True(t, passwordHash.Compare(rawPassword))
}

func TestNewPassword_Invalid(t *testing.T) {
	empties := make([]rune, 64)
	testCases := []struct{
		name string
		input string
		expected error
	}{
		{
			name: "empty_password",
			input: "",
			expected: ErrInvalidPassword,
		},
		{
			name: "too_long_password",
			input: string(empties),
			expected: ErrInvalidPassword,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewPasswordHash(tc.input)
			
			require.Error(t, err)
			if !errors.Is(err, tc.expected) {
				t.Errorf("want %s, got %s", err.Error(), tc.expected.Error())
			}
		})
	}
}

func TestInvalidPasswordError(t *testing.T) {
	err := &InvalidPasswordError{
		original: ErrInvalidPassword,
		msg: "msg",
	}

	if !errors.Is(err, ErrInvalidPassword) {
		t.Errorf("want %s, got %s", ErrInvalidPassword.Error(), err.Error())
	}

	assert.EqualError(t, err, "invalid password: msg")
}