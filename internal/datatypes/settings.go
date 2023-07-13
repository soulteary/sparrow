package datatypes

type BetaFeatures struct {
	Browsing        bool `json:"browsing"`
	CodeInterpreter bool `json:"code_interpreter"`
	Plugins         bool `json:"plugins"`
	ChatPreferences bool `json:"chat_preferences"` // added 0713
}
