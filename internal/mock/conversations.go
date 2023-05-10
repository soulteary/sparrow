package mock

import (
	"fmt"
	"time"

	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetConversationList() datatypes.ConversationsList {
	var items []datatypes.ConversationListItem
	if !define.ENABLE_HISTORY_LIST {
		items = []datatypes.ConversationListItem{}
	} else {
		for i := 10; i < 30; i++ {
			items = append(items, datatypes.ConversationListItem{
				ID:         define.GenerateUUID(),
				Title:      fmt.Sprintf("会话名称 %d", i-10+1),
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			})
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
