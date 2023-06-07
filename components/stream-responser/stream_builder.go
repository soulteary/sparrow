package StreamResponser

import (
	"time"

	eb "github.com/soulteary/sparrow/components/event-broker"
	lcs "github.com/soulteary/sparrow/components/local-conversation-storage"
	FlagStudio "github.com/soulteary/sparrow/connectors/flag-studio"
	GitHubTop "github.com/soulteary/sparrow/connectors/github-top"
	OpenaiAPI "github.com/soulteary/sparrow/connectors/openai-api"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

type StreamMessageMode int

var (
	MSG_STATUS_AUTO_MODE StreamMessageMode = 0
	MSG_STATUS_CONTINUE  StreamMessageMode = 1
	MSG_STATUS_DONE      StreamMessageMode = 2
)

func StreamBuilder(userID string, conversationID string, parentMessageID string, messageID string, nextMessageID string, modelSlug string, broker *eb.Broker, input string, mode StreamMessageMode) bool {
	if modelSlug == "" {
		modelSlug = "text-davinci-002-render-sha"
	}

	var sequences []string
	var quickMode bool
	var output string

	switch modelSlug {
	case datatypes.MODEL_OPENAI_API_3_5.Slug:
		if define.ENABLE_OPENAI_API {
			if define.OPENAI_API_KEY == "" {
				output = "OpenAI API Key needs to be set correctly."
			} else {
				output = OpenaiAPI.Get(input)
			}
			sequences = MakeStreamingMessage(output, modelSlug, conversationID, nextMessageID, mode)
			quickMode = false
		}
	case datatypes.MODEL_TEXT_DAVINCI_002_PLUGINS.Slug:
	// case datatypes.MODEL_TEXT_DAVINCI_002_RENDER_PAID.Slug:
	case datatypes.MODEL_TEXT_DAVINCI_002_RENDER_SHA.Slug:
	case datatypes.MODEL_GPT4.Slug:
	case datatypes.MODEL_MIDJOURNEY.Slug:
		if define.ENABLE_MIDJOURNEY {
			output = input
			sequences = MakeStreamingMessage(input, modelSlug, conversationID, nextMessageID, mode)
			quickMode = true
		}
	case datatypes.MODEL_FLAGSTUDIO.Slug:
		if define.ENABLE_FLAGSTUDIO {
			output := FlagStudio.GenerateImageByText(input)
			sequences = MakeStreamingMessage(output, modelSlug, conversationID, nextMessageID, mode)
			quickMode = true
		}
	case datatypes.MODEL_CLAUDE.Slug:
		if define.ENABLE_CLAUDE {
			output = input
			sequences = MakeStreamingMessage(input, modelSlug, conversationID, nextMessageID, mode)
			quickMode = true
		}
	case datatypes.MODEL_GITHUB_TOP.Slug:
		if define.ENABLE_GITHUB_TOP {
			output = GitHubTop.HandleUserPrompt(input)
			sequences = MakeStreamingMessage(output, modelSlug, conversationID, nextMessageID, mode)
			quickMode = true
		}
		// case datatypes.MODEL_NO_MODELS.Slug:
		// default:
	}

	if len(sequences) > 0 {
		lcs.SetMessage(userID, messageID, nextMessageID, output, false)
		return MakeStreamingResponse(conversationID, parentMessageID, broker, sequences, quickMode)
	}

	output = "The administrator has disabled the export capability of this model.\nProject: [soulteary/sparrow](https://github.com/soulteary/sparrow).\nTalk is Cheap, Let's coding together."
	sequences = MakeStreamingMessage(output, modelSlug, conversationID, nextMessageID, mode)

	lcs.SetMessage(userID, messageID, nextMessageID, output, false)
	return MakeStreamingResponse(conversationID, parentMessageID, broker, sequences, true)
}

func MakeStreamingResponse(conversationID string, parentMessageID string, broker *eb.Broker, sequences []string, quickMode bool) bool {
	count := len(sequences)
	if count == 0 {
		return false
	}

	simulateDelay := 800 / define.RESPONSE_SPEED
	if define.DEV_MODE || quickMode {
		simulateDelay = 10
	}

	go func() {
		lastThreeBefore := count - 3
		for id, sequence := range sequences {
			if id <= 2 {
				time.Sleep(time.Millisecond * time.Duration(simulateDelay))
			}

			broker.Event <- eb.Event{
				ConversationID:  conversationID,
				ParentMessageID: parentMessageID,
				Payload:         sequence,
			}

			if !quickMode {
				if id < lastThreeBefore {
					time.Sleep(time.Millisecond * time.Duration(RandomResponseTime(40, 120)))
				} else {
					// Acceleration end output
					time.Sleep(time.Millisecond * time.Duration(50))
				}
			} else {
				time.Sleep(time.Millisecond * time.Duration(10))
			}
		}
	}()
	return true
}
