package define

import (
	"os"
	"strings"
)

var (
	DEV_MODE = false // Development mode
)

var (
	APP_PORT            = GetPortString("APP_PORT", DEFAULT_APP_PORT)
	WEB_CLIENT_HOSTNAME = GetHostName("WEB_CLIENT_HOSTNAME", DEFAULT_WEB_CLIENT_HOSTNAME)
	RESPONSE_SPEED      = GetGenerateSpeed("RESPONSE_SPEED", 10)

	ENABLE_HISTORY_LIST      = GetBool("ENABLE_HISTORY_LIST", true)       // Enable history list
	ENABLE_I18N              = GetBool("ENABLE_I18N", true)               // Enable i18n
	ENABLE_DATA_CONTROL      = GetBool("ENABLE_DATA_CONTROL", false)      // Enable the data control
	ENABLE_MODEL_SWITCH      = GetBool("ENABLE_MODEL_SWITCH", true)       // Enable the model switch
	ENABLE_PAID_SUBSCRIPTION = GetBool("ENABLE_PAID_SUBSCRIPTION", false) // Enable the subscription
	ENABLE_PLUGIN            = GetBool("ENABLE_PLUGIN", true)             // Enable the plugin

	ENABLE_OPENAI_OFFICIAL_MODEL = GetBool("ENABLE_OPENAI_OFFICIAL_MODEL", true) // Enable the official model
	ENABLE_OPENAI_ONLY_3_5       = GetBool("ENABLE_OPENAI_ONLY_3_5", true)       // Only Enable the 3.5 model

	ENABLE_OPENAI_API       = GetBool("ENABLE_OPENAI_API", false)                                 // Enable OpenAI 3.5 API
	OPENAI_API_KEY          = strings.TrimSpace(os.Getenv("OPENAI_API_KEY"))                      // OpenAI API Key
	ENABLE_OPENAI_API_PROXY = GetBool("OPENAI_API_PROXY_ENABLE", false)                           // Enable OpenAI API Proxy
	OPENAI_API_PROXY_ADDR   = GetHostName("OPENAI_API_PROXY_ADDR", DEFAULT_OPENAI_API_PROXY_ADDR) // OpenAI API Proxy Address
)
