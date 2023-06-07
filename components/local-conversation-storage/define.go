package LocalConversationStorage

import "time"

// [<UserSpace> User Space|
// 	[<conversation> Conversation 1...N|
// 		[Message A] -> [Message B]
// 		[Message B] -> [Message C]
// 		[Message C] -> [Message ...]
// 	]
// ]

// [<Message> Message|
//     [Conversation ID]
//     [Message ID]
//     [Parent ID]
// ]

type Message struct {
	ID         string // The unique ID of the message.
	Parent     string
	Children   []string
	CreateTime time.Time
	Content    string // MESSAGE BODY
	IsUser     bool
}

type UserID string

// storage the links between parent message and child message
type LinkReferences map[string]string

// storage the messages
type MessagesData map[string]Message

// storage the user's messages
type UserMessages map[UserID]LinkReferences

// storage the user's conversations
type UserConversations map[UserID][]string
