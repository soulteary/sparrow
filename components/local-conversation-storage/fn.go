package LocalConversationStorage

import (
	"errors"
	"fmt"

	"github.com/soulteary/sparrow/internal/define"
)

func getRefsByUserID(uid UserID) LinkReferences {
	refs, exist := Refs[uid]
	if refs == nil || !exist {
		return make(LinkReferences)
	}
	return refs
}

func updateRefsByUserID(uid UserID, refs LinkReferences) {
	Refs[uid] = refs

	data, _ := define.MakeJSON(refs)
	SaveMeta(string(uid), "refs.json", data)
}

func getParentRef(refs LinkReferences, messageID string) (ref string, exist bool) {
	ref, exist = refs[messageID]
	if !exist {
		return "", false
	}
	return ref, true
}

func updateMessageParentRefs(uid UserID, messageID string, parentMessageID string) {
	refs := getRefsByUserID(uid)
	refs[messageID] = parentMessageID
	updateRefsByUserID(uid, refs)
}

func getConversationsByUserID(uid UserID) []string {
	conversations, exist := Conversations[uid]
	if conversations == nil || !exist {
		return make([]string, 0)
	}
	return conversations
}

func updateConversationByUserID(uid UserID, conversations []string) {
	Conversations[uid] = conversations

	data, _ := define.MakeJSON(conversations)
	SaveMeta(string(uid), "conversations.json", data)
}

func GetConversationDataByMessageID(messageID string) (Message, error) {
	data, exist := Data[messageID]
	if !exist {
		return Message{}, errors.New("message not found")
	}
	return data, nil
}

func updateConversationDataByMessageID(messageID string, data Message) {
	Data[messageID] = data

	jsonData, _ := define.MakeJSON(data)
	SaveMessage(messageID, jsonData)
}

func IsStrInArray(strs []string, s string) bool {
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
	fmt.Println(ret)
}
