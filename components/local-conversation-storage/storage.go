package LocalConversationStorage

import (
	"time"

	"github.com/soulteary/sparrow/internal/define"
)

// Determine whether the message is the root message
func IsRootMessage(userID string, conversationID string, parentMessageID string) bool {
	refs := getRefsByUserID(UserID(userID))
	_, exist := getParentRef(refs, parentMessageID)
	if !exist && conversationID == parentMessageID {
		return true
	}
	return false
}

// Update the message to the message list
func addConversationToUserConversationList(userID string, conversationID string) {
	uid := UserID(userID)
	conversations := getConversationsByUserID(uid)
	if !IsStrInArray(conversations, conversationID) {
		conversations = append(conversations, conversationID)
	}
	updateConversationByUserID(uid, conversations)
}

// Get the message list of the user
func GetConversationListByUserID(userID string) (conversations []string) {
	return getConversationsByUserID(UserID(userID))
}

// Clear the message list of the user
func ClearConversationListByUserID(userID string) {
	uid := UserID(userID)
	updateConversationByUserID(uid, make([]string, 0))
	refs := getRefsByUserID(uid)
	for _, refID := range refs {
		updateConversationDataByMessageID(refID, Message{})
	}
	updateRefsByUserID(uid, make(LinkReferences))
}

// Clear the message id of the user
func ClearConversationIdByUserID(userID string, conversationID string) {
	uid := UserID(userID)
	refs := getRefsByUserID(uid)
	filteredRefs := make(LinkReferences)
	for _, refID := range refs {
		if refID == conversationID {
			updateConversationDataByMessageID(refID, Message{})
		} else {
			filteredRefs[refID] = refs[refID]
		}
	}
	updateRefsByUserID(uid, filteredRefs)
}

// Get the conversation info by conversation id
func GetConversationInfoByID(userID string, conversationID string) (conversationInfo Message) {
	refs := getRefsByUserID(UserID(userID))
	messageID, exist := getParentRef(refs, conversationID)
	if !exist {
		return conversationInfo
	}
	conversationInfo, err := GetConversationDataByMessageID(messageID)
	if err != nil {
		return conversationInfo
	}
	return conversationInfo
}

// Create root message and return the root message id
func CreateRootMessages(userID string, childID string) (rootMessageID string) {
	rootMessageID = define.GenerateUUID()
	updateConversationDataByMessageID(rootMessageID, Message{
		ID:         rootMessageID,
		Children:   []string{childID},
		Content:    "Root Message",
		CreateTime: time.Now(),
	})
	updateMessageParentRefs(UserID(userID), rootMessageID, rootMessageID)
	return rootMessageID
}

// Set the message as roote message
func SetRootMessage(userID string, parentMessageID string, messageID string, content string) (rootMessageID string) {
	rootMessageID = CreateRootMessages(userID, parentMessageID)
	addConversationToUserConversationList(userID, rootMessageID)

	uid := UserID(userID)
	updateMessageParentRefs(uid, messageID, parentMessageID)

	updateConversationDataByMessageID(parentMessageID, Message{
		ID:         parentMessageID,
		Parent:     rootMessageID,
		Children:   []string{messageID},
		CreateTime: time.Now(),
	})
	LinkMessageToOtherAsChild(userID, rootMessageID, parentMessageID)

	updateConversationDataByMessageID(messageID, Message{
		ID:         messageID,
		Parent:     parentMessageID,
		Children:   []string{},
		Content:    content,
		CreateTime: time.Now(),
		IsUser:     true,
	})
	LinkMessageToOtherAsChild(userID, parentMessageID, messageID)
	return rootMessageID
}

// Link a message to other message as child
func LinkMessageToOtherAsChild(userID string, parentMessageID string, messageID string) {
	updateMessageParentRefs(UserID(userID), messageID, parentMessageID)

	originMessage, err := GetConversationDataByMessageID(messageID)
	if err != nil {
		return
	}

	originParentMessage, err := GetConversationDataByMessageID(parentMessageID)
	if err != nil {
		return
	}

	message := originMessage
	message.Parent = parentMessageID
	updateConversationDataByMessageID(messageID, message)

	parentMessage := originParentMessage
	if !IsStrInArray(parentMessage.Children, messageID) {
		parentMessage.Children = append(parentMessage.Children, messageID)
	}
	updateConversationDataByMessageID(parentMessageID, parentMessage)
}

func SetMessage(userID string, parentMessageID string, messageID string, content string, isUser bool) {
	message, err := GetConversationDataByMessageID(messageID)
	if err != nil {
		message = Message{
			ID:         messageID,
			Content:    content,
			CreateTime: time.Now(),
			IsUser:     isUser,
		}
	}

	message.Content = content
	message.IsUser = isUser

	uid := UserID(userID)
	updateMessageParentRefs(uid, messageID, parentMessageID)

	if parentMessageID != "" {
		message.Parent = parentMessageID
	}

	updateConversationDataByMessageID(messageID, message)
	LinkMessageToOtherAsChild(userID, parentMessageID, messageID)

	if define.DEV_MODE {
		Debug(uid)
	}
}
