package vo

import (
	"fmt"
	"testing"
)

func TestSuccessNewAmount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		value    float64
		expected float64
	}{
		{
			value:    10.01,
			expected: 10.01,
		},
		{
			value:    10.011,
			expected: 10.01,
		},
		{
			value:    10.199,
			expected: 10.20,
		},
		{
			value:    999.9999,
			expected: 1000.00,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Amount test: %f", test.value), func(t *testing.T) {
			amount, _ := NewAmount(test.value)

			if amount.Value() != test.expected {
				t.Errorf("Expected %f got %f", test.expected, amount.Value())
			}
		})
	}
}

func TestErrorNewAmount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		value float64
	}{
		{
			value: -10.0,
		},
		{
			value: -0.00001,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Amount test: %f", test.value), func(t *testing.T) {
			_, err := NewAmount(test.value)

			if err == nil {
				t.Errorf("Expected nil got %s", err.Error())
			}
		})
	}
}
