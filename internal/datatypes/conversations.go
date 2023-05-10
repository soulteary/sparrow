package datatypes

import "time"

type ConversationsList struct {
	Items                   []ConversationListItem `json:"items"`
	Total                   int                    `json:"total"`
	Limit                   int                    `json:"limit"`
	Offset                  int                    `json:"offset"`
	HasMissingConversations bool                   `json:"has_missing_conversations"`
}

type ConversationListItem struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
