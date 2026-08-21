package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmail_Success(t *testing.T) {
	raw := "mail@example.com"
	email := Email(raw)

	require.NoError(t, email.Validate())
	assert.Equal(t, email.String(), raw)
	assert.Equal(t, email.Norm().String(), raw)

}

func TestNewEmail_Invalid(t *testing.T) {
	testCases := []struct{
		name string
		input string
		expected error
	}{
		{
			name: "empty_email",
			input: "",
			expected: ErrInvalidEmail,
		},
		{
			name: "only_spaced_email",
			input: "     \n\t",
			expected: ErrInvalidEmail,
		},
		{
			name: "invalid_mail",
			input: "@example.com",
			expected: ErrInvalidEmail,
		},
		{
			name: "invalid_domain",
			input: "mail@",
			expected: ErrInvalidEmail,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := Email(tc.input).Validate()

			require.Error(t, err)
			if !errors.Is(err, tc.expected) {
				t.Errorf("want %s, got %s", tc.expected.Error(), err.Error())
			}
		})
	}
}