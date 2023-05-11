package StreamResponser

import (
	"regexp"
	"strings"
	"time"

	eb "github.com/soulteary/sparrow/components/event-broker"
	OpenaiAPI "github.com/soulteary/sparrow/connectors/openai-api"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
	"github.com/soulteary/sparrow/internal/mock"
)

func StreamBuilder(data datatypes.Conversation, modelSlug string, broker *eb.Broker, input string) bool {
	var sequences []string
	messageID := define.GenerateUUID()
	if modelSlug == "" {
		modelSlug = "text-davinci-002-render-sha"
	}
	if define.ENABLE_OPENAI_API {
		sequences = MakeStreamingMessage(OpenaiAPI.Get(input), modelSlug, data.ConversationID, messageID)
		return MakeStreamingResponse(data, broker, sequences)
	}
	sequences = MakeStreamingMessage("The administrator has disabled the export capability of this model.\nProject: [soulteary/sparrow](https://github.com/soulteary/sparrow).\nTalk is Cheap, Let's coding together.", modelSlug, data.ConversationID, messageID)
	return MakeStreamingResponse(data, broker, sequences)
}

func MakeStreamingResponse(data datatypes.Conversation, broker *eb.Broker, sequences []string) bool {
	count := len(sequences)
	if count == 0 {
		return false
	}

	simulateDelay := 800 / define.RESPONSE_SPEED
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
				time.Sleep(time.Millisecond * time.Duration(RandomResponseTime(40, 120)))
			} else {
				// Acceleration end output
				time.Sleep(time.Millisecond * time.Duration(50))
			}
		}
	}()
	return true
}

func MakeResponseMessage(text string, modelSlug, conversationID string, messageID string, endTurn bool) datatypes.ConversationMessageGenerated {
	data := datatypes.ConversationMessageGenerated{
		ConversationID: conversationID,
		Message: datatypes.ConversationMessage{
			ID:      messageID,
			Author:  mock.CreateMessageUser("assistant"),
			Content: mock.CreateMessageContent("text", text),
			Metadata: datatypes.ConversationMessageGeneratedMetaBody{
				MessageType: "variant",
				ModelSlug:   modelSlug,
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
			ModelSlug:   modelSlug,
			FinishDetails: datatypes.ConversationMessageGeneratedMessageMetaType{
				Type: "stop",
			},
		}
		data.Message.EndTurn = false
	}
	return data
}

func MakeStreamingMessage(text string, modelSlug string, conversationID string, messageID string) (ret []string) {
	var messages []datatypes.ConversationMessageGenerated

	// Simulate the feeling of waiting for a response
	message := MakeResponseMessage("", modelSlug, conversationID, messageID, false)
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
				messages = append(messages, MakeResponseMessage(s, modelSlug, conversationID, messageID, false))
			} else {
				context := strings.Split(match, "")
				for i := 0; i < len(context); i++ {
					s += context[i]
					messages = append(messages, MakeResponseMessage(s, modelSlug, conversationID, messageID, false))
				}
			}
		}

		// before last line
		if lineId != lastLineId {
			s += "\n"
			messages = append(messages, MakeResponseMessage(s, modelSlug, conversationID, messageID, false))
		} else {
			messages = append(messages, MakeResponseMessage(s, modelSlug, conversationID, messageID, true))
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
