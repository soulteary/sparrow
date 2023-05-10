package mock

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func createPluginToolMessage(createTime int64, input string, namespace string) (conversation datatypes.ConversationHistory) {
	conversation.ID = define.GenerateUUID()
	conversation.Parent = ""
	conversation.Children = []string{}

	var message datatypes.ConversationMessage
	message.ID = conversation.ID
	message.Author = createPluginMessageUser("tool", namespace)
	message.CreateTime = float64(createTime)
	message.UpdateTime = float64(createTime + 1)
	message.Content = createMessageContent("text", input)
	message.Metadata = createPluginToolMessageMeta("text-davinci-002-plugins", "absolute", namespace)

	message.EndTurn = true // TODO double check
	message.Weight = 1.0
	message.Recipient = "all"

	conversation.Message = message
	return conversation
}
