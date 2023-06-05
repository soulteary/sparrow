package LocalConversationStorage

import (
	"os"
	"time"

	"github.com/soulteary/sparrow/internal/define"
)

func InitStorage(filename string) {
	os.MkdirAll(filename, os.ModePerm)
}

// Determine whether the parent message exists
func IsParentMessageExist(userID string, parentMessageID string) bool {
	refs := GetRefsByUserID(UserID(userID))
	_, err := GetRefsToParent(refs, parentMessageID)
	return err != nil
}

// Update the message to the message list
func AddConversationToUserConversationList(userID string, conversationID string) {
	uid := UserID(userID)
	conversations := GetConversationsByUserID(uid)
	if !StrInArray(conversations, conversationID) {
		conversations = append(conversations, conversationID)
	}
	UpdateConversationByUserID(uid, conversations)
}

// Get the message list of the user
func GetConversationListByUserID(userID string) (conversations []string) {
	return GetConversationsByUserID(UserID(userID))
}

// Clear the message list of the user
func ClearConversationListByUserID(userID string) {
	uid := UserID(userID)
	UpdateConversationByUserID(uid, make([]string, 0))
	refs := GetRefsByUserID(uid)
	for _, refID := range refs {
		UpdateConversationDataByMessageID(refID, Message{})
	}
	UpdateRefsByUserID(uid, make(LinkReferences))
}

// Get the conversation info by conversation id
func GetConversationInfoByID(userID string, conversationID string) (conversationInfo Message) {
	refs := GetRefsByUserID(UserID(userID))
	messageID, err := GetRefsToParent(refs, conversationID)
	if err != nil {
		return conversationInfo
	}
	conversationInfo, err = ConversationDataByMessageID(messageID)
	if err != nil {
		return conversationInfo
	}
	return conversationInfo
}

// Create root message and return the root message id
func CreateRootMessages(userID string, childID string) (rootMessageID string) {
	rootMessageID = define.GenerateUUID()
	UpdateConversationDataByMessageID(rootMessageID, Message{
		ID:         rootMessageID,
		Children:   []string{childID},
		Content:    "Root Message",
		CreateTime: time.Now(),
	})
	UpdateMessageParentRefs(UserID(userID), rootMessageID, rootMessageID)
	return rootMessageID
}

// Set the message as roote message
func SetRootMessage(userID string, parentMessageID string, messageID string, content string) (rootMessageID string) {
	rootMessageID = CreateRootMessages(userID, parentMessageID)
	AddConversationToUserConversationList(userID, rootMessageID)

	uid := UserID(userID)
	refs := GetRefsByUserID(uid)
	refs[messageID] = parentMessageID
	UpdateRefsByUserID(uid, refs)

	UpdateConversationDataByMessageID(parentMessageID, Message{
		ID:         parentMessageID,
		Parent:     rootMessageID,
		Children:   []string{messageID},
		Content:    content,
		CreateTime: time.Now(),
	})
	LinkMessageToOtherAsChild(userID, rootMessageID, parentMessageID)

	UpdateConversationDataByMessageID(messageID, Message{
		ID:         messageID,
		Parent:     parentMessageID,
		Children:   []string{},
		Content:    content,
		CreateTime: time.Now(),
	})
	LinkMessageToOtherAsChild(userID, parentMessageID, messageID)
	return rootMessageID
}

// Link a message to other message as child
func LinkMessageToOtherAsChild(userID string, parentMessageID string, messageID string) {
	refs := GetRefsByUserID(UserID(userID))
	refs[messageID] = parentMessageID
	UpdateRefsByUserID(UserID(userID), refs)

	originMessage, err := ConversationDataByMessageID(messageID)
	if err != nil {
		return
	}

	originParentMessage, err := ConversationDataByMessageID(parentMessageID)
	if err != nil {
		return
	}

	message := originMessage
	message.Parent = parentMessageID
	UpdateConversationDataByMessageID(messageID, message)

	parentMessage := originParentMessage
	if !StrInArray(parentMessage.Children, messageID) {
		parentMessage.Children = append(parentMessage.Children, messageID)
	}
	UpdateConversationDataByMessageID(parentMessageID, parentMessage)
}

func SetMessage(userID string, parentMessageID string, messageID string, content string) {
	refs := GetRefsByUserID(UserID(userID))
	refs[messageID] = parentMessageID
	UpdateRefsByUserID(UserID(userID), refs)

	message, err := ConversationDataByMessageID(messageID)
	if err != nil {
		message = Message{
			ID:         messageID,
			Content:    content,
			CreateTime: time.Now(),
		}
	}
	message.Content = content
	if parentMessageID != "" {
		message.Parent = parentMessageID
	}

	UpdateConversationDataByMessageID(messageID, message)
	LinkMessageToOtherAsChild(userID, parentMessageID, messageID)

	Debug(UserID(userID))
}
