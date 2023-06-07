package datatypes

type GeneralConversationHistory struct {
	Title             string                         `json:"title"`
	CreateTime        float64                        `json:"create_time"`
	UpdateTime        float64                        `json:"update_time"`
	Mapping           map[string]ConversationHistory `json:"mapping"`
	ModerationResults []string                       `json:"moderation_results"`
	CurrentNode       string                         `json:"current_node"`
}

type PluginConversationHistory struct {
	Title             string                         `json:"title"`
	CreateTime        float64                        `json:"create_time"`
	UpdateTime        float64                        `json:"update_time"`
	Mapping           map[string]ConversationHistory `json:"mapping"`
	ModerationResults []string                       `json:"moderation_results"`
	CurrentNode       string                         `json:"current_node"`
	PluginIds         []string                       `json:"plugin_ids"`
}

type ConversationHistory struct {
	ID       string   `json:"id"`
	Message  any      `json:"message,omitempty"`
	Parent   string   `json:"parent,omitempty"`
	Children []string `json:"children"`
}

type ConversationMessage struct {
	ID         string  `json:"id"`
	Author     any     `json:"author"`
	CreateTime float64 `json:"create_time,omitempty"`
	UpdateTime float64 `json:"update_time,omitempty"`
	Content    any     `json:"content"`
	EndTurn    any     `json:"end_turn,omitempty"`
	Weight     float64 `json:"weight"`
	Metadata   any     `json:"metadata"`
	Recipient  string  `json:"recipient,omitempty"`
}

// update date to official api, 2023.03.02
type PostConversationMessage struct {
	ID         string                `json:"id"`
	Author     GeneralMessageAuthor  `json:"author"`
	CreateTime float64               `json:"create_time"`
	UpdateTime float64               `json:"update_time,omitempty"`
	Content    GeneralMessageContent `json:"content"`
	EndTurn    bool                  `json:"end_turn,omitempty"`
	Weight     float64               `json:"weight"`
	Metadata   any                   `json:"metadata"`
	Recipient  string                `json:"recipient"`
}

type GeneralMessageAuthor struct {
	Role     string `json:"role"`
	Metadata struct {
	} `json:"metadata"`
}

type PluginMessageAuthor struct {
	Role     string `json:"role"`
	Name     string `json:"name"`
	Metadata struct {
	} `json:"metadata"`
}

type GeneralMessageContent struct {
	ContentType string   `json:"content_type,omitempty"`
	Parts       []string `json:"parts,omitempty"`
}

type PluginMessageContent struct {
	ContentType  string            `json:"content_type"`
	Text         string            `json:"text"`
	ToolsSection map[string]string `json:"tools_section"`
}

type EmptyConversationMessageMetaBody struct{}

type GeneralConversationMessageMeta struct {
	ModelSlug     string `json:"model_slug,omitempty"`
	FinishDetails struct {
		Type string `json:"type,omitempty"`
		Stop string `json:"stop,omitempty"`
	} `json:"finish_details,omitempty"`
	Timestamp string `json:"timestamp_,omitempty"` // TODO missing `_` at the end
}

type PluginConversationMessageMeta struct {
	InvokedPlugin struct {
		Type      string `json:"type"`
		Namespace string `json:"namespace"`
		PluginID  string `json:"plugin_id"`
	} `json:"invoked_plugin"`
	ModelSlug string `json:"model_slug"`
	Timestamp string `json:"timestamp_"`
}

type ConversationMessageGeneratedMetaBody struct {
	MessageType string `json:"message_type"`
	ModelSlug   string `json:"model_slug"`
}

type ConversationMessageMetaTS struct {
	Timestamp string `json:"timestamp,omitempty"`
}

type UpdateConversationResponse struct {
	Success bool `json:"success"`
}

// update date to official api, 2023.03.02
type Conversation struct {
	Action                     string                    `json:"action"`
	Messages                   []PostConversationMessage `json:"messages"`
	ConversationID             string                    `json:"conversation_id,omitempty"`
	ParentMessageID            string                    `json:"parent_message_id"`
	Model                      string                    `json:"model"`
	TimezoneOffsetMin          int                       `json:"timezone_offset_min"`
	HistoryAndTrainingDisabled bool                      `json:"history_and_training_disabled"`
}

type ConversationMessageGenerated struct {
	Message        ConversationMessage `json:"message"`
	ConversationID string              `json:"conversation_id"`
	Error          any                 `json:"error"`
}

type ConversationMessageGeneratedMessageMetaType struct {
	Type string `json:"type"`
}

type ConversationMessageGeneratedMessageMeta struct {
	MessageType   string                                      `json:"message_type"`
	ModelSlug     string                                      `json:"model_slug"`
	FinishDetails ConversationMessageGeneratedMessageMetaType `json:"finish_details"`
}
