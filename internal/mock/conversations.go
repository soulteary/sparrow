package mock

import (
	"fmt"

	lcs "github.com/soulteary/sparrow/components/local-conversation-storage"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetConversationList(userID string) datatypes.ConversationsList {
	var items []datatypes.ConversationListItem
	if !define.ENABLE_HISTORY_LIST {
		items = []datatypes.ConversationListItem{}
	} else {
		items = []datatypes.ConversationListItem{}

		conversationIDs := lcs.GetConversationListByUserID(userID)
		if len(conversationIDs) == 0 {
			items = []datatypes.ConversationListItem{}
		} else {
			for i, conversationID := range conversationIDs {
				conversationInfo := lcs.GetConversationInfoByID(userID, conversationID)

				if conversationInfo.ID == conversationID {
					items = append(items, datatypes.ConversationListItem{
						ID:          conversationID,
						Title:       fmt.Sprintf("Conversation %d", i),
						CreateTime:  conversationInfo.CreateTime,
						Mapping:     nil,
						CurrentNode: nil,
					})
				}

			}
		}
	}

	return datatypes.ConversationsList{
		Total:                   len(items),
		Limit:                   20,
		Offset:                  0,
		HasMissingConversations: false,
		Items:                   items,
	}
}

func ClearConversationList(userID string) datatypes.ConversationsList {
	if define.ENABLE_HISTORY_LIST {
		lcs.ClearConversationListByUserID(userID)
	}
	return GetConversationList(userID)
}

func ClearConversationByID(userID string, conversationID string) {
	if define.ENABLE_HISTORY_LIST {
		lcs.ClearConversationIdByUserID(userID, conversationID)
	}
}
