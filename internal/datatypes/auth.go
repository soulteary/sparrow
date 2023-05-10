package datatypes

type API_SESSION struct {
	User         API_SESSION_USER `json:"user"`
	Expires      string           `json:"expires"`
	AccessToken  string           `json:"accessToken"`
	AuthProvider string           `json:"authProvider"`
}

type API_SESSION_USER struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Image        string   `json:"image"`
	Picture      string   `json:"picture"`
	Mfa          bool     `json:"mfa"`
	Groups       []string `json:"groups"`
	IntercomHash string   `json:"intercom_hash"`
}
