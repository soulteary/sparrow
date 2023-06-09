package StreamResponser_test

import (
	"strings"
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

func TestParseConversationBody(t *testing.T) {
	body := `{"action":"next","messages":[{"id":"aaa262a7-7b7e-4484-9b04-89417c8679a5","author":{"role":"user"},"content":{"content_type":"text","parts":["123"]}}],"parent_message_id":"aaa1ca9c-563a-43b3-ad39-c10eb1bceaa1","model":"no-models","timezone_offset_min":-480,"history_and_training_disabled":false}`

	r := strings.NewReader(body)
	originBody, result, err := sr.ParseConversationBody(r)

	if err != nil {
		t.Errorf("error parsing conversation body: %v", err)
	}

	if string(originBody) != body {
		t.Errorf("unexpected origin body. expected=%s, got=%s", body, string(originBody))
	}

	expectedAction := "next"
	if result.Action != expectedAction {
		t.Errorf("unexpected action. expected=%s, got=%s", expectedAction, result.Action)
	}

	expectedHistoryAndTrainingDisabled := false
	if result.HistoryAndTrainingDisabled != expectedHistoryAndTrainingDisabled {
		t.Errorf("unexpected history_and_training_disabled. expected=%t, got=%t", expectedHistoryAndTrainingDisabled, result.HistoryAndTrainingDisabled)
	}

	expectedModel := "no-models"
	if result.Model != expectedModel {
		t.Errorf("unexpected model. expected=%s, got=%s", expectedModel, result.Model)
	}

	expectedParentMessageId := "aaa1ca9c-563a-43b3-ad39-c10eb1bceaa1"
	if result.ParentMessageID != expectedParentMessageId {
		t.Errorf("unexpected parent_message_id. expected=%s, got=%s", expectedParentMessageId, result.ParentMessageID)
	}

	expectedTimezoneOffsetMin := -480
	if result.TimezoneOffsetMin != expectedTimezoneOffsetMin {
		t.Errorf("unexpected timezone_offset_min. expected=%d, got=%d", expectedTimezoneOffsetMin, result.TimezoneOffsetMin)
	}

	expectedMessageId := "aaa262a7-7b7e-4484-9b04-89417c8679a5"
	if result.Messages[0].ID != expectedMessageId {
		t.Errorf("unexpected message_id. expected=%s, got=%s", expectedMessageId, result.Messages[0].ID)
	}

	expectedParts := "123"
	if result.Messages[0].Content.Parts[0] != expectedParts {
		t.Errorf("unexpected parts. expected=%s, got=%s", expectedParts, result.Messages[0].Content.Parts[0])
	}
}

func TestContainMarkdownImage(t *testing.T) {
	// Test for false case
	str := "This is a plain text"
	if sr.ContainMarkdownImage(str) {
		t.Errorf("Failed for string: %s", str)
	}

	// Test for true case
	str = "This is a markdown image ![alt text](image.jpg)"
	if !sr.ContainMarkdownImage(str) {
		t.Errorf("Failed for string: %s", str)
	}
}

func TestContainMarkdownLink(t *testing.T) {
	// Test for false case
	str := "This is a plain text"
	if sr.ContainMarkdownLink(str) {
		t.Errorf("Failed for string: %s", str)
	}

	// Test for true case
	str = "This is a markdown link [link text](https://www.example.com)"
	if !sr.ContainMarkdownLink(str) {
		t.Errorf("Failed for string: %s", str)
	}
}
