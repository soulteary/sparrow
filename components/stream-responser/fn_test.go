package StreamResponser_test

import (
	"testing"

	sr "github.com/soulteary/sparrow/components/stream-responser"
	"github.com/soulteary/sparrow/internal/define"
)

func TestRandomResponseTime(t *testing.T) {
	tests := []struct {
		name     string
		devMode  bool
		min      int
		max      int
		expected int
		mode     string
		speed    int
	}{
		{
			name:     "under dev mode should be defalt value",
			devMode:  true,
			min:      20,
			max:      50,
			expected: 10,
			mode:     "=",
			speed:    1,
		},
		{
			name:     "under dev mode should be defalt value",
			devMode:  true,
			min:      0,
			max:      0,
			expected: 10,
			mode:     "=",
			speed:    1,
		},
		{
			name:     "under dev mode should be defalt value",
			devMode:  true,
			min:      -10,
			max:      0,
			expected: 10,
			mode:     "=",
			speed:    1,
		},
		{
			name:     "error input should be default use default value",
			devMode:  false,
			min:      -20,
			max:      50,
			expected: 40,
			mode:     ">=",
			speed:    1,
		},
		{
			name:     "error input should be default use default value",
			devMode:  false,
			min:      20,
			max:      0,
			expected: 40,
			mode:     ">=",
			speed:    1,
		},

		{
			name:     "normal test",
			devMode:  false,
			min:      100,
			max:      200,
			expected: 100,
			mode:     ">=",
			speed:    1,
		},
		{
			name:     "test speed 10x",
			devMode:  false,
			min:      100,
			max:      200,
			expected: 10,
			mode:     ">=",
			speed:    10,
		},
		{
			name:     "excessive multiple",
			devMode:  false,
			min:      100,
			max:      200,
			expected: 10,
			mode:     "=",
			speed:    10000,
		},
		{
			name:     "error speed",
			devMode:  false,
			min:      100,
			max:      200,
			expected: 100,
			mode:     ">=",
			speed:    -123,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			define.RESPONSE_SPEED = test.speed
			define.DEV_MODE = test.devMode
			got := sr.RandomResponseTime(test.min, test.max)
			switch test.mode {
			case "=":
				if got != test.expected {
					t.Errorf("RandomResponseTime(%d,%d) = %d; expected %s %d", test.min, test.max, got, test.mode, test.expected)
				}
			case ">=":
				if got < test.expected {
					t.Errorf("RandomResponseTime(%d,%d) = %d; expected %s %d", test.min, test.max, got, test.mode, test.expected)
				}
			case "<=":
				if got > test.expected {
					t.Errorf("RandomResponseTime(%d,%d) = %d; expected %s %d", test.min, test.max, got, test.mode, test.expected)
				}
			}
		})
	}
}
