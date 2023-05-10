package datatypes

type PageInfo struct {
	Props struct {
		PageProps struct {
			User struct {
				ID           string   `json:"id"`
				Name         string   `json:"name"`
				Email        string   `json:"email"`
				Image        string   `json:"image"`
				Picture      string   `json:"picture"`
				Idp          string   `json:"idp"` // 23.05.08 added
				Mfa          bool     `json:"mfa"`
				Groups       []string `json:"groups"`
				IntercomHash string   `json:"intercom_hash"`
			} `json:"user"`
			ServiceStatus struct {
			} `json:"serviceStatus"`
			UserCountry         string `json:"userCountry"`
			GeoOk               bool   `json:"geoOk"`
			ServiceAnnouncement struct {
				Public struct {
				} `json:"public"`
				Paid struct {
				} `json:"paid"`
			} `json:"serviceAnnouncement"`
			IsUserInCanPayGroup bool   `json:"isUserInCanPayGroup"`
			SentryTraceData     string `json:"_sentryTraceData"` // 23.05.08 added
			SentryBaggage       string `json:"_sentryBaggage"`   // 23.05.08 added
		} `json:"pageProps"`
		NSsp bool `json:"__N_SSP"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		ChatID string `json:"chatId,omitempty"`
		Model  string `json:"model,omitempty"`
	} `json:"query"`
	BuildID      string   `json:"buildId"`
	IsFallback   bool     `json:"isFallback"`
	Gssp         bool     `json:"gssp"`
	ScriptLoader []string `json:"scriptLoader"`
}
