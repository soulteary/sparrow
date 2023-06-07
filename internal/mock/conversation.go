package mock

import (
	"time"

	lcs "github.com/soulteary/sparrow/components/local-conversation-storage"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetConversationById(userID, conversationID string) any {
	if define.DEV_MODE {
		// random response conversation type
		if define.GetRandomNumber(1, 2) == 1 {
			return GeneralConversationHistory(conversationID)
		} else {
			return PluginConversationHistory(conversationID)
		}
	}
	return GetConversationHistory(userID, conversationID)
}

func GetConversationHistory(userID, conversationID string) datatypes.GeneralConversationHistory {
	var conversations []*datatypes.ConversationHistory

	firstMessage := lcs.GetConversationInfoByID(userID, conversationID)
	systemConversation := createSystemConversation(firstMessage.CreateTime.Unix(), firstMessage.ID)
	conversations = append(conversations, &systemConversation)

	secondMessage, _ := lcs.GetConversationDataByMessageID(firstMessage.Children[0])
	frontEndMessage := createMapMessage(secondMessage.ID)
	linkMessages(&frontEndMessage, &systemConversation)
	conversations = append(conversations, &frontEndMessage)

	var prevMessage datatypes.ConversationHistory
	prevMessage = systemConversation
	messageID := secondMessage.Children[0]

	for {
		var newMessage datatypes.ConversationHistory
		message, _ := lcs.GetConversationDataByMessageID(messageID)
		if message.IsUser {
			newMessage = createUserMessage(message.CreateTime.Unix(), message.ID, message.Content)
			linkMessages(&prevMessage, &newMessage)
			conversations = append(conversations, &newMessage)
		} else {
			newMessage = createAssistantMessage(message.CreateTime.Unix(), message.ID, message.Content)
			linkMessages(&prevMessage, &newMessage)
			conversations = append(conversations, &newMessage)
		}
		if message.Children != nil && len(message.Children) > 0 {
			messageID = message.Children[0]
			prevMessage = newMessage
		} else {
			break
		}
	}

	mapping := make(map[string]datatypes.ConversationHistory)
	for _, conversation := range conversations {
		mapping[conversation.ID] = *conversation
	}

	return datatypes.GeneralConversationHistory{
		// TODO get conversation title
		Title:             conversations[len(conversations)-1].ID,
		CreateTime:        float64(time.Now().Unix()),
		UpdateTime:        float64(time.Now().Unix()),
		ModerationResults: []string{},
		CurrentNode:       conversations[len(conversations)-1].ID,
		Mapping:           mapping,
	}
}

func GeneralConversationHistory(requestID string) datatypes.GeneralConversationHistory {
	var conversations []*datatypes.ConversationHistory

	// message 1
	now := time.Now().Unix()
	systemConversation := createSystemConversation(now+1, "")
	conversations = append(conversations, &systemConversation)

	// message 2
	frontEndMessage := createMapMessage(requestID)
	linkMessages(&frontEndMessage, &systemConversation)
	conversations = append(conversations, &frontEndMessage)

	// message 3
	userMessage := createUserMessage(now+2, "", "用户输入的内容")
	linkMessages(&systemConversation, &userMessage)
	conversations = append(conversations, &userMessage)

	// message 4
	reply := "\n\n这里是来自接口的内容。 ![](https://images.openai.com/blob/8d14e8f0-e267-4b8b-a9f2-a79120808f5a/chatgpt.jpg?trim=0%2C0%2C0%2C0)"
	replyMessage := createAssistantMessage(now+3, "", reply)
	linkMessages(&userMessage, &replyMessage)
	conversations = append(conversations, &replyMessage)

	// message 5
	userMessage2 := createUserMessage(now+4, "", "还有什么？")
	linkMessages(&replyMessage, &userMessage2)
	conversations = append(conversations, &userMessage2)

	// message 6
	reply = "接口返回的更多内容"
	replyMessage2 := createAssistantMessage(now+5, "", reply)

	linkMessages(&userMessage2, &replyMessage2)
	conversations = append(conversations, &replyMessage2)

	mapping := make(map[string]datatypes.ConversationHistory)
	for _, conversation := range conversations {
		mapping[conversation.ID] = *conversation
	}

	return datatypes.GeneralConversationHistory{
		Title:             "一段简单的对话",
		CreateTime:        float64(time.Now().Unix()),
		UpdateTime:        float64(time.Now().Unix()),
		ModerationResults: []string{},
		CurrentNode:       conversations[len(conversations)-1].ID,
		Mapping:           mapping,
	}
}

func PluginConversationHistory(requestID string) datatypes.PluginConversationHistory {
	var conversations []*datatypes.ConversationHistory

	pluginNS := "KAYAK"

	// message 1
	now := time.Now().Unix()
	systemConversation := createPluginSystemConversation(now+1, pluginNS)
	conversations = append(conversations, &systemConversation)

	// message 2
	frontEndMessage := createMapMessage(requestID)
	linkMessages(&frontEndMessage, &systemConversation)
	conversations = append(conversations, &frontEndMessage)

	// message 3
	userMessage := createUserMessage(now+2, "", "随便聊点啥")
	linkMessages(&systemConversation, &userMessage)
	conversations = append(conversations, &userMessage)

	// message 4
	reply := "{\n  \"origin\": \"ABCD\",\n  \"destination\": \"EFGH\",\n  \"departDate\": \"2099-01-01\"\n}"
	replyMessage1 := createPluginAssistantMessage(now+3, reply, false, pluginNS+".searchFlights", "<|im_end|>")
	linkMessages(&userMessage, &replyMessage1)
	conversations = append(conversations, &replyMessage1)

	// message 5
	reply = "I understood you want to search for flights from ABCD to EFGH from 2099-01-01\r\n\r\nPrices for flights start at $1,234. The majority of flights overall are offered by American Airlines, United Airlines and Alaska Airlines. Most flights that have one or more stops have a layover at Seattle or Los Angeles. The shortest total trip duration is 15h 00m.\r\n\r\nRecommended flights leaving Anchorage, AK are available in the following hours: 0am-2am, 5am-1pm, 3pm, 5pm, 7pm, 9pm and 11pm.\r\n\r\nTo see more flights and book go to https://www.kayak.com/flights/ABCD-EFGH/2023-01-01?a=openai"
	replyMessage2 := createPluginToolMessage(now+3, reply, pluginNS)
	linkMessages(&replyMessage1, &replyMessage2)
	conversations = append(conversations, &replyMessage2)

	// message 6
	reply = "这里展示模型和插件协同工作后，输出的内容。\n\n 这里可以包含链接，图片，以及比如一个包含部分预览卡片的东西，ps：展示文档更合适了: \n - [例子：来自百度的文档](https://www.baidu.com/flights/ABCD-EFGH/2023-01-01?a=openai) \n - [例子：来自微博的文档](https://www.weibo.com/flights/ABCD-EFGH/2023-01-01?a=openai)"
	replyMessage3 := createPluginAssistantMessage(now+3, reply, true, "all", "<|diff_marker|>")
	linkMessages(&replyMessage2, &replyMessage3)
	conversations = append(conversations, &replyMessage3)

	mapping := make(map[string]datatypes.ConversationHistory)
	for _, conversation := range conversations {
		mapping[conversation.ID] = *conversation
	}

	return datatypes.PluginConversationHistory{
		Title:             "插件会话记录展示",
		CreateTime:        float64(time.Now().Unix()),
		UpdateTime:        float64(time.Now().Unix()),
		ModerationResults: []string{},
		CurrentNode:       conversations[len(conversations)-1].ID,
		Mapping:           mapping,
		PluginIds:         []string{"plugin-" + define.GenerateUUID()},
	}
}
