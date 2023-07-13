package mock

import (
	"fmt"

	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetConversationLimit() datatypes.ConversationLimit {
	const INVOKE_MAX_CAP = 100
	const INVOKE_HOUR_INTERVAL = 3
	conversationLimit := datatypes.ConversationLimit{
		MessageCap:       INVOKE_MAX_CAP,
		MessageCapWindow: INVOKE_HOUR_INTERVAL * 60,
		MessageDisclaimer: datatypes.ConversationLimitMessageDisclaimer{
			ModelSwitcher: "You've reached the GPT-4 cap, which gives all ChatGPT Plus users a chance to try the model.\n\nPlease check back soon.",
			Textarea:      fmt.Sprintf("GPT-4 currently has a cap of %d messages every %d hours.", INVOKE_MAX_CAP, INVOKE_HOUR_INTERVAL),
		},
	}

	if define.ENABLE_I18N {
		conversationLimit.MessageDisclaimer.ModelSwitcher = "您已用尽模型调用配额。\n\n请稍后再试。"
		conversationLimit.MessageDisclaimer.Textarea = fmt.Sprintf("GPT-4 目前每 %d 小时允许 %d 次调用。", INVOKE_HOUR_INTERVAL, INVOKE_MAX_CAP)
	}
	return conversationLimit
}
