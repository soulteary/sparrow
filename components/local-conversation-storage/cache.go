package LocalConversationStorage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/soulteary/sparrow/internal/define"
)

var Conversations = make(UserConversations)
var Data = make(MessagesData)
var Refs = make(UserMessages)

func init() {
	LoadMeta()
	LoadMessages()
}

func InitStorage(filename string) {
	os.MkdirAll(filename, os.ModePerm)
}

func SaveMeta(uid string, filename string, data string) {
	if !define.ENABLE_HISTORY_LIST {
		return
	}
	dir := filepath.Join(define.LOCAL_CONVERSATION_STORAGE_PATH, uid)
	InitStorage(dir)
	os.WriteFile(filepath.Join(dir, filename), []byte(data), os.ModePerm)
}

func SaveMessage(messageId string, data string) {
	if !define.ENABLE_HISTORY_LIST {
		return
	}
	dir := filepath.Join(define.LOCAL_CONVERSATION_STORAGE_PATH, "messages")
	InitStorage(dir)
	os.WriteFile(filepath.Join(dir, messageId+".json"), []byte(data), os.ModePerm)
}

func ParseConversations(uid string) (data []string, err error) {
	buf, err := os.ReadFile(filepath.Join(define.LOCAL_CONVERSATION_STORAGE_PATH, uid, "conversations.json"))
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func ParseRefs(uid string) (data map[string]string, err error) {
	buf, err := os.ReadFile(filepath.Join(define.LOCAL_CONVERSATION_STORAGE_PATH, uid, "refs.json"))
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func LoadMeta() {
	if !define.ENABLE_HISTORY_LIST {
		return
	}
	dirs, _ := os.ReadDir(define.LOCAL_CONVERSATION_STORAGE_PATH)
	users := []string{}
	for _, dir := range dirs {
		if dir.IsDir() && dir.Name() != "messages" {
			users = append(users, dir.Name())
		}
	}
	for _, user := range users {
		uid := UserID(user)
		conversations, _ := ParseConversations(user)
		Conversations[uid] = conversations
		refs, _ := ParseRefs(user)
		Refs[uid] = refs
	}
}

func LoadMessages() {
	if !define.ENABLE_HISTORY_LIST {
		return
	}
	files, _ := os.ReadDir(filepath.Join(define.LOCAL_CONVERSATION_STORAGE_PATH, "messages"))
	messages := []string{}
	for _, items := range files {
		if !items.IsDir() {
			messages = append(messages, items.Name())
		}
	}

	for _, message := range messages {
		buf, _ := os.ReadFile(filepath.Join(define.LOCAL_CONVERSATION_STORAGE_PATH, "messages", message))
		var data Message
		json.Unmarshal(buf, &data)
		Data[data.ID] = data
	}
}
