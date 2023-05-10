package datatypes

type ConversationLimit struct {
	MessageCap        int                                `json:"message_cap"`
	MessageCapWindow  int                                `json:"message_cap_window"`
	MessageDisclaimer ConversationLimitMessageDisclaimer `json:"message_disclaimer"`
}

type ConversationLimitMessageDisclaimer struct {
	ModelSwitcher string `json:"model-switcher"`
	Textarea      string `json:"textarea"`
}
