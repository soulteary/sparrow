package EventBroker_test

import (
	"testing"

	eb "github.com/soulteary/sparrow/components/event-broker"
)

func TestIsLastMessage(t *testing.T) {
	tests := []struct {
		name     string
		payload  interface{}
		expected bool
	}{
		{
			name:     "empty payload",
			payload:  "",
			expected: false,
		},
		{
			name:     "last message",
			payload:  "[DONE]",
			expected: true,
		},
		{
			name:     "last message with spaces",
			payload:  "  [DONE]   ",
			expected: true,
		},
		{
			name:     "non-last message",
			payload:  "Hello World",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := eb.IsLastMessage(test.payload); got != test.expected {
				t.Errorf("IsLastMessage(%v) = %t; expected %t", test.payload, got, test.expected)
			}
		})
	}
}
