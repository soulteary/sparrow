package mock

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func AuthSession() datatypes.API_SESSION {
	data := datatypes.API_SESSION{
		User: datatypes.API_SESSION_USER{
			ID:           define.MOCK_USER_ID,
			Name:         define.MOCK_USER_NAME,
			Email:        define.MOCK_USER_EMAIL,
			Image:        define.MOCK_USER_IMAGE,
			Picture:      define.MOCK_USER_IMAGE,
			Mfa:          false,
			Groups:       []string{},
			IntercomHash: define.GenerateRandomString(64),
		},
		Expires:     "2099-12-31T23:59:59.000Z",
		AccessToken: define.MOCK_USER_TOKEN,
	}

	if define.ENABLE_PLUGIN {
		data.User.Groups = append(data.User.Groups, "chatgpt-plugin-partners")
		data.AuthProvider = "auth0"
	}
	return data
}
