package mock

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func createAssistantMessage(createTime int64, uuid string, input string) (conversation datatypes.ConversationHistory) {
	if uuid == "" {
		conversation.ID = define.GenerateUUID()
	} else {
		conversation.ID = uuid
	}
	conversation.Parent = ""
	conversation.Children = []string{}

	var message datatypes.ConversationMessage
	message.ID = conversation.ID
	message.Author = CreateMessageUser("assistant")
	message.CreateTime = float64(createTime)
	message.Content = CreateMessageContent("text", input)
	message.Metadata = createModelMessageMeta("text-davinci-002-render-sha", "absolute")

	message.EndTurn = true
	message.Weight = 1.0
	message.Recipient = "all"

	conversation.Message = message
	return conversation
}

func createPluginAssistantMessage(createTime int64, input string, isEnd bool, recipient string, stopWord string) (conversation datatypes.ConversationHistory) {
	conversation.ID = define.GenerateUUID()
	conversation.Parent = ""
	conversation.Children = []string{}

	var message datatypes.ConversationMessage
	message.ID = conversation.ID
	message.Author = CreateMessageUser("assistant")
	message.CreateTime = float64(createTime)
	message.Content = CreateMessageContent("text", input)
	message.Metadata = createPluginModelMessageMeta("text-davinci-002-plugins", "absolute", stopWord)

	message.EndTurn = isEnd
	message.EndTurn = true // TODO double check

	message.Weight = 1.0
	message.Recipient = recipient

	conversation.Message = message
	return conversation
}
