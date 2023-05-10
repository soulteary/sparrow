package conversation

import (
	"regexp"
	"strings"
	"time"

	eb "github.com/soulteary/sparrow/components/event-broker"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
	"github.com/soulteary/sparrow/internal/mock"
)

func StreamBuilder(data datatypes.Conversation, broker *eb.Broker, prompt string) bool {
	sequences := MakeMessageSequence(data.ParentMessageID, data.ConversationID, prompt)

	count := len(sequences)
	if count == 0 {
		return false
	}

	simulateDelay := 800
	if define.DEV_MODE {
		simulateDelay = 10
	}

	go func() {
		lastThreeBefore := count - 3
		for id, sequence := range sequences {
			if id <= 2 {
				time.Sleep(time.Millisecond * time.Duration(simulateDelay))
			}

			broker.Event <- eb.Event{
				Name:    data.ParentMessageID,
				Payload: sequence,
			}

			if id < lastThreeBefore {
				RandomSleep()
			} else {
				// Acceleration end output
				time.Sleep(time.Millisecond * time.Duration(50))
			}
		}
	}()
	return true
}

func MakeMessageSequence(parentMessageID string, conversationID string, userInput string) (ret []string) {
	return MakeStreamingMessage("The administrator has disabled the export capability of this model.", conversationID)
}

func MakeResponseMessage(text string, conversationID string, newMessageID string, endTurn bool) datatypes.ConversationMessageGenerated {
	data := datatypes.ConversationMessageGenerated{
		ConversationID: conversationID,
		Message: datatypes.ConversationMessage{
			ID:      newMessageID,
			Author:  mock.CreateMessageUser("assistant"),
			Content: mock.CreateMessageContent("text", text),
			Metadata: datatypes.ConversationMessageGeneratedMetaBody{
				MessageType: "variant",
				ModelSlug:   "text-davinci-002-render-sha",
			},
			Weight:    1,
			Recipient: "all",
			EndTurn:   nil,
		},
		Error: nil,
	}

	if endTurn {
		data.Message.Metadata = datatypes.ConversationMessageGeneratedMessageMeta{
			MessageType: "variant",
			ModelSlug:   "text-davinci-002-render-sha",
			FinishDetails: datatypes.ConversationMessageGeneratedMessageMetaType{
				Type: "stop",
			},
		}
		data.Message.EndTurn = false
	}
	return data
}

func MakeStreamingMessage(text string, conversationID string) (ret []string) {
	newConversationID := define.GenerateUUID()

	var messages []datatypes.ConversationMessageGenerated

	// Simulate the feeling of waiting for a response
	message := MakeResponseMessage("", conversationID, newConversationID, false)
	for i := 0; i < 3; i++ {
		messages = append(messages, message)
	}

	s := ""
	lines := strings.Split(text, "\n")
	lastLineId := len(lines) - 1
	for lineId, line := range lines {
		r, _ := regexp.Compile(`!\[[^\]]*\]\([^\)]+\)|\[[^\]]+\]\([^\)]+\)|[^!\[\]\(\)\n]+|\n`)
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			if ContainMarkdownImage(match) || ContainMarkdownLink(match) {
				s += match
				messages = append(messages, MakeResponseMessage(s, conversationID, newConversationID, false))
			} else {
				context := strings.Split(match, "")
				for i := 0; i < len(context); i++ {
					s += context[i]
					messages = append(messages, MakeResponseMessage(s, conversationID, newConversationID, false))
				}
			}
		}

		// before last line
		if lineId != lastLineId {
			s += "\n"
			messages = append(messages, MakeResponseMessage(s, conversationID, newConversationID, false))
		} else {
			messages = append(messages, MakeResponseMessage(s, conversationID, newConversationID, true))
		}
	}

	for _, message := range messages {
		text, err := MakeJSON(message)
		if err == nil {
			ret = append(ret, " "+text)
		}
	}
	ret = append(ret, "[DONE]")
	return ret
}
