package StreamResponser

import (
	"time"

	eb "github.com/soulteary/sparrow/components/event-broker"
	OpenaiAPI "github.com/soulteary/sparrow/connectors/openai-api"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func StreamBuilder(data datatypes.Conversation, modelSlug string, broker *eb.Broker, input string) bool {
	messageID, modelSlug := GetBuilderParams(modelSlug)
	if define.ENABLE_OPENAI_API {
		sequences := MakeStreamingMessage(OpenaiAPI.Get(input), modelSlug, data.ConversationID, messageID)
		return MakeStreamingResponse(data, broker, sequences)
	}
	sequences := MakeStreamingMessage("The administrator has disabled the export capability of this model.\nProject: [soulteary/sparrow](https://github.com/soulteary/sparrow).\nTalk is Cheap, Let's coding together.", modelSlug, data.ConversationID, messageID)
	return MakeStreamingResponse(data, broker, sequences)
}

func StreamBuilderManual(data datatypes.Conversation, modelSlug string, broker *eb.Broker, input string, markAsEnd bool) bool {
	return false
}

func GetBuilderParams(modelSlug string) (string, string) {
	messageID := define.GenerateUUID()
	if modelSlug == "" {
		modelSlug = "text-davinci-002-render-sha"
	}
	return messageID, modelSlug
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
