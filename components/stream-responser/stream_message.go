package StreamResponser

import (
	"regexp"
	"strings"

	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
	"github.com/soulteary/sparrow/internal/mock"
)

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

func MakeStreamingMessage(text string, modelSlug string, conversationID string, messageID string, mode StreamMessageMode) (ret []string) {
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
			if mode == MSG_STATUS_AUTO_MODE || mode == MSG_STATUS_DONE {
				messages = append(messages, MakeResponseMessage(s, modelSlug, conversationID, messageID, true))
			}
		}
	}

	for _, message := range messages {
		text, err := define.MakeJSON(message)
		if err == nil {
			ret = append(ret, " "+text)
		}
	}
	if mode == MSG_STATUS_AUTO_MODE || mode == MSG_STATUS_DONE {
		ret = append(ret, "[DONE]")
	}
	return ret
}
