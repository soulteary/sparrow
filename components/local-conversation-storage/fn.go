package LocalConversationStorage

import (
	"errors"
	"fmt"

	"github.com/soulteary/sparrow/internal/define"
)

func GetRefsByUserID(uid UserID) LinkReferences {
	refs, ok := Refs[uid]
	if refs == nil || !ok {
		return make(LinkReferences)
	}
	return refs
}

func UpdateRefsByUserID(uid UserID, refs LinkReferences) {
	Refs[uid] = refs
}

func GetRefsToParent(refs LinkReferences, messageID string) (string, error) {
	id, ok := refs[messageID]
	if !ok {
		return "", errors.New("message not found")
	}
	return id, nil
}

func UpdateMessageParentRefs(uid UserID, messageID string, parentMessageID string) {
	refs := GetRefsByUserID(uid)
	refs[messageID] = parentMessageID
	UpdateRefsByUserID(uid, refs)
}

func GetConversationsByUserID(uid UserID) []string {
	conversations, ok := Conversations[uid]
	if conversations == nil || !ok {
		return make([]string, 0)
	}
	return conversations
}

func UpdateConversationByUserID(uid UserID, conversations []string) {
	Conversations[uid] = conversations
}

func ConversationDataByMessageID(messageID string) (Message, error) {
	data, ok := Data[messageID]
	if !ok {
		return Message{}, errors.New("message not found")
	}
	return data, nil
}

func UpdateConversationDataByMessageID(messageID string, data Message) {
	Data[messageID] = data
}

func StrInArray(strs []string, s string) bool {
	exist := false
	for _, v := range strs {
		if v == s {
			exist = true
			break
		}
	}
	return exist
}

func Debug(uid UserID) {
	ret, _ := define.MakeJSON(Data)
	fmt.Println(uid)
	fmt.Println(ret)
}
