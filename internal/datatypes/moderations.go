package datatypes

type Moderations struct {
	Flagged      bool   `json:"flagged"`
	Blocked      bool   `json:"blocked"`
	ModerationID string `json:"moderation_id"`
}
