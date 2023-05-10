package datatypes

type AIPluginResponse struct {
	Count int        `json:"count"`
	Items []AIPlugin `json:"items"`
}

type AIPlugin struct {
	ID            string               `json:"id"`
	Domain        string               `json:"domain"`
	Namespace     string               `json:"namespace"`
	Status        string               `json:"status"`
	Manifest      AIPluginManifest     `json:"manifest"`
	OauthClientID any                  `json:"oauth_client_id"`
	UserSettings  AIPluginUserSettings `json:"user_settings"`
	Categories    []any                `json:"categories"`
}

type AIPluginUserSettings struct {
	IsInstalled     bool `json:"is_installed"`
	IsAuthenticated bool `json:"is_authenticated"`
}

type AIPluginAuth struct {
	Type string `json:"type"`
}
type AIPluginAPI struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type AIPluginManifest struct {
	SchemaVersion       string       `json:"schema_version"`
	NameForModel        string       `json:"name_for_model"`
	NameForHuman        string       `json:"name_for_human"`
	DescriptionForModel string       `json:"description_for_model"`
	DescriptionForHuman string       `json:"description_for_human"`
	Auth                AIPluginAuth `json:"auth"`
	API                 AIPluginAPI  `json:"api"`
	LogoURL             string       `json:"logo_url"`
	ContactEmail        string       `json:"contact_email"`
	LegalInfoURL        string       `json:"legal_info_url"`
}
