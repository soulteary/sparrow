package mock

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func createSystemConversation(createTime int64, uuid string) (conversation datatypes.ConversationHistory) {
	if uuid == "" {
		conversation.ID = define.GenerateUUID()
	} else {
		conversation.ID = uuid
	}
	conversation.Parent = ""
	conversation.Children = []string{}

	var message datatypes.ConversationMessage
	message.ID = conversation.ID
	message.Author = CreateMessageUser("system")
	message.CreateTime = float64(createTime)
	message.Content = CreateMessageContent("text", "")
	message.Metadata = createEmptyMessageMeta()

	message.EndTurn = true
	message.Weight = 1.0
	message.Recipient = "all"

	conversation.Message = message
	return conversation
}

func createPluginSystemConversation(createTime int64, namespace string) (conversation datatypes.ConversationHistory) {
	conversation.ID = define.GenerateUUID()
	conversation.Parent = ""
	conversation.Children = []string{}

	var message datatypes.ConversationMessage
	message.ID = conversation.ID
	message.Author = CreateMessageUser("system")
	message.CreateTime = float64(createTime)
	message.Content = createPluginSystemMessageContent(namespace)
	message.Metadata = createEmptyMessageMeta()

	message.EndTurn = true // TODO double check
	message.Weight = 1.0
	message.Recipient = "all"

	conversation.Message = message
	return conversation
}
