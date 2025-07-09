package password

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordEntropy(t *testing.T) {
	tests := []struct {
		password string
		expected string
	}{
		{"1234567890", "Password is too weak"},
		{"abcdefghij", "Password is weak"},
		{"abcdefghij1A", "Normal password"},
		{"abcdefghij1A&%)", "Good password"},
		{"Aj8@fGk$1pZoAj8@fGk$1pZo]})", "Excellent password"},
	}

	for _, testCase := range tests {
		t.Run(testCase.password, func(t *testing.T) {
			result, err := passwordEntropy(testCase.password)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, result,
				fmt.Sprintf("Incorrect result. Expected %s, got %s",
					testCase.expected, result))
		})
	}
}
