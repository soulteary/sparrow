package define_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/soulteary/sparrow/internal/define"
)

func TestGetBool(t *testing.T) {
	tests := []struct {
		name     string
		envKey   string
		value    string
		def      bool
		expected bool
	}{
		{
			name:     "EMPTY",
			envKey:   "EMPTY_ENV",
			value:    "",
			def:      true,
			expected: true,
		},
		{
			name:     "EMPTY",
			envKey:   "EMPTY_ENV",
			value:    "",
			def:      false,
			expected: false,
		},
		{
			name:     "on",
			envKey:   "ON_ENV",
			value:    "on",
			def:      false,
			expected: true,
		},
		{
			name:     "on",
			envKey:   "ON_ENV",
			value:    "on",
			def:      true,
			expected: true,
		},
		{
			name:     "true",
			envKey:   "TRUE_ENV",
			value:    "true",
			def:      false,
			expected: true,
		},
		{
			name:     "1",
			envKey:   "ONE_ENV",
			value:    "1",
			def:      false,
			expected: true,
		},
		{
			name:     "off",
			envKey:   "OFF_ENV",
			value:    "off",
			def:      true,
			expected: false,
		},
		{
			name:     "false",
			envKey:   "FALSE_ENV",
			value:    "false",
			def:      true,
			expected: false,
		},
		{
			name:     "0",
			envKey:   "ZERO_ENV",
			value:    "0",
			def:      true,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Setenv(test.envKey, test.value)
			if got := define.GetBool(test.envKey, test.def); got != test.expected {
				t.Errorf("GetBool(%s,%t) = %t; expected %t", test.envKey, test.def, got, test.expected)
			}
			defer os.Unsetenv(test.envKey)
		})
	}
}

func TestGetPortString(t *testing.T) {
	tests := []struct {
		name     string
		envKey   string
		value    string
		def      int
		expected string
	}{
		{
			name:     "empty env",
			envKey:   "EMPTY_ENV",
			value:    "",
			def:      8080,
			expected: ":8080",
		},
		{
			name:     "valid port env",
			envKey:   "PORT_VALID_ENV",
			value:    "abcd",
			def:      8080,
			expected: ":8080",
		},
		{
			name:     "valid port env",
			envKey:   "PORT_VALID_ENV",
			value:    "-1234",
			def:      8080,
			expected: ":8080",
		},
		{
			name:     "valid port env",
			envKey:   "PORT_VALID_ENV",
			value:    "123456789",
			def:      8080,
			expected: ":8080",
		},
		{
			name:     "valid port env",
			envKey:   "PORT_VALID_ENV",
			value:    "0",
			def:      8080,
			expected: ":8080",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Setenv(test.envKey, test.value)
			if got := define.GetPortString(test.envKey, test.def); got != test.expected {
				t.Errorf("GetPortString(%s,%d) = %s; expected %s", test.envKey, test.def, got, test.expected)
			}
			defer os.Unsetenv(test.envKey)
		})
	}
}

func TestGetHostName(t *testing.T) {
	tests := []struct {
		name     string
		envKey   string
		value    string
		def      string
		expected string
	}{
		{
			name:     "empty env",
			envKey:   "EMPTY_ENV",
			value:    "",
			def:      "http://localhost:8080",
			expected: "http://localhost:8080",
		},
		{
			name:     "valid env",
			envKey:   "VALID_ENV",
			value:    "http://www.google.com/api",
			def:      "localhost:8080",
			expected: "http://www.google.com/api",
		},
		{
			name:     "invalid env",
			envKey:   "INVALID_ENV",
			value:    "s!ss!sss..!@#.s",
			def:      "http://localhost:8080/abc",
			expected: "http://localhost:8080/abc",
		},
		{
			name:     "invalid env",
			envKey:   "INVALID_ENV",
			value:    "abcd://abcd.abcd.abcd",
			def:      "http://localhost:8080/abc",
			expected: "http://localhost:8080/abc",
		},
		{
			name:     "valid env",
			envKey:   "VALID_ENV",
			value:    "http://abcd.abcd.abcd/path",
			def:      "http://localhost:8080/abc",
			expected: "http://abcd.abcd.abcd/path",
		},
		{
			name:     "valid env",
			envKey:   "VALID_ENV",
			value:    "ws://abcd.abcd.abcd/ws",
			def:      "http://localhost:8080/abc",
			expected: "ws://abcd.abcd.abcd/ws",
		},
		{
			name:     "valid env",
			envKey:   "VALID_ENV",
			value:    "ws://abcd.abcd.abcd/ws/",
			def:      "http://localhost:8080/abc",
			expected: "ws://abcd.abcd.abcd/ws",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Setenv(test.envKey, test.value)
			if got := define.GetHostName(test.envKey, test.def); got != test.expected {
				fmt.Println(got)
				t.Errorf("GetHostName(%s,%s) = %s; expected %s", test.envKey, test.def, got, test.expected)
			}
			defer os.Unsetenv(test.envKey)
		})
	}
}

func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{
			name:   "length 0",
			length: 0,
		},
		{
			name:   "length 10",
			length: 10,
		},
		{
			name:   "length 20",
			length: 20,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := define.GenerateRandomString(test.length)
			if len(got) != test.length {
				t.Errorf("GenerateRandomString(%d) = %s; length mismatch", test.length, got)
			}
		})
	}
}

func TestGenerateUUID(t *testing.T) {
	if define.GenerateUUID() == "" {
		t.Fatal("GenerateUUID() = \"\"; expected UUID")
	}
}

func TestGetMidJourneySecret(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		def    string
		envKey string
		expect string
	}{
		{
			name:   "empty string",
			envKey: "empty-key",
			input:  "",
			def:    "def",
			expect: "def",
		},
		{
			name:   "user input string",
			envKey: "env-key",
			input:  "abc",
			def:    "def",
			expect: "abc",
		},
	}

	for _, test := range tests {
		os.Setenv(test.envKey, test.input)
		defer os.Unsetenv(test.envKey)
		t.Run(test.name, func(t *testing.T) {
			got := define.GetMidJourneySecret(test.envKey, test.def)
			if got != test.expect {
				t.Errorf("GetMidJourneySecret(%s, %s) = %s; expect %s", test.envKey, test.def, got, test.expect)
			}
		})
	}
}

func TestGetRandomNumber(t *testing.T) {
	for i := 0; i < 10; i++ {
		randomNumber := define.GetRandomNumber(1, 100)
		if randomNumber < 1 || randomNumber > 100 {
			t.Errorf("GetRandomNumber(%d, %d) = %d; want a number between 1 and 100", 1, 100, randomNumber)
		}
	}
}

func TestMakeJSON(t *testing.T) {
	type test struct {
		input interface{}
		want  string
		err   bool
	}

	tests := []test{
		{input: "hello", want: "\"hello\""},
		{input: 123, want: "123"},
		{input: []string{"a", "b", "c"}, want: "[\"a\",\"b\",\"c\"]"},
		{input: map[string]int{"a": 1, "b": 2}, want: "{\"a\":1,\"b\":2}"},
		{input: func() {}, err: true},
	}

	for _, tc := range tests {
		got, err := define.MakeJSON(tc.input)
		if tc.err {
			if err == nil {
				t.Errorf("MakeJSON(%v) expected an error but did not get one", tc.input)
			}
			continue
		}
		if err != nil {
			t.Errorf("MakeJSON(%v) got unexpected error: %v", tc.input, err)
			continue
		}
		if got != tc.want {
			t.Errorf("MakeJSON(%v) = %q; want %q", tc.input, got, tc.want)
		}
	}
}
