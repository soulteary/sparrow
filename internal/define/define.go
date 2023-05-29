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

	ENABLE_HISTORY_LIST      = GetBool("ENABLE_HISTORY_LIST", false)      // Enable history list
	ENABLE_I18N              = GetBool("ENABLE_I18N", false)              // Enable i18n
	ENABLE_DATA_CONTROL      = GetBool("ENABLE_DATA_CONTROL", false)      // Enable the data control
	ENABLE_PAID_SUBSCRIPTION = GetBool("ENABLE_PAID_SUBSCRIPTION", false) // Enable the subscription

	ENABLE_NEW_UI       = GetBool("FEATURE_NEW_UI", false)      // Enable the new UI
	ENABLE_MODEL_SWITCH = GetBool("ENABLE_MODEL_SWITCH", false) // Enable the model switch
	NEW_TOOLS_USER      = GetBool("NEW_TOOLS_USER", false)      // New tools user

	ENABLE_PLUGIN                  = GetBool("ENABLE_PLUGIN", false)                  // Enable the plugin
	ENABLE_PLUGIN_PLUGIN_DEV       = GetBool("ENABLE_PLUGIN_PLUGIN_DEV", false)       // Enable the plugin
	ENABLE_PLUGIN_BROWSING         = GetBool("ENABLE_PLUGIN_BROWSING", false)         // Enable the plugin, Browsing
	ENABLE_PLUGIN_CODE_INTERPRETER = GetBool("ENABLE_PLUGIN_CODE_INTERPRETER", false) // Enable the plugin, CodeInterpreter
)

var (
	ENABLE_OPENAI_OFFICIAL_MODEL = GetBool("ENABLE_OPENAI_OFFICIAL_MODEL", false) // Enable the official model
)

var (
	ENABLE_OPENAI_API       = GetBool("ENABLE_OPENAI_API", false)                                 // Enable OpenAI 3.5 API
	ENABLE_OPENAI_API_ONLY  = GetBool("ENABLE_OPENAI_API_ONLY", false)                            // Only Enable the 3.5 model
	OPENAI_API_KEY          = strings.TrimSpace(os.Getenv("OPENAI_API_KEY"))                      // OpenAI API Key
	ENABLE_OPENAI_API_PROXY = GetBool("OPENAI_API_PROXY_ENABLE", false)                           // Enable OpenAI API Proxy
	OPENAI_API_PROXY_ADDR   = GetHostName("OPENAI_API_PROXY_ADDR", DEFAULT_OPENAI_API_PROXY_ADDR) // OpenAI API Proxy Address
)

var (
	ENABLE_MIDJOURNEY      = GetBool("ENABLE_MIDJOURNEY", false)                            // Enable Midjourney
	ENABLE_MIDJOURNEY_ONLY = GetBool("ENABLE_MIDJOURNEY_ONLY", false)                       // Enable Midjourney only
	MIDJOURNEY_API_SECRET  = GetSecret("MIDJOURNEY_API_SECRET", "YOUR_MIDJOURNEY_SECRET")   // Midjourney API Secret
	MIDJOURNEY_API_ADDR    = GetHostName("MIDJOURNEY_API_URL", DEFAULT_MIDJOURNEY_API_ADDR) // Midjourney API URL
)

var (
	ENABLE_FLAGSTUDIO      = GetBool("ENABLE_FLAGSTUDIO", false)                       // Enable Flagstudio
	ENABLE_FLAGSTUDIO_ONLY = GetBool("ENABLE_FLAGSTUDIO_ONLY", false)                  // Enable Flagstudio only
	FLAGSTUDIO_API_KEY     = GetSecret("FLAGSTUDIO_API_KEY", "YOUR_FLAGSTUDIO_SECRET") // Flagstudio API Token
)

var (
	ENABLE_CLAUDE      = GetBool("ENABLE_CLAUDE", false)                        // Enable Claude
	ENABLE_CLAUDE_ONLY = GetBool("ENABLE_CLAUDE_ONLY", false)                   // Enable Claude only
	CLAUDE_API_SECRET  = GetSecret("CLAUDE_API_SECRET", "YOUR_CLAUDE_SECRET")   // Claude API Secret
	CLAUDE_API_ADDR    = GetHostName("CLAUDE_API_URL", DEFAULT_CLAUDE_API_ADDR) // Claude API URL
)

var (
	ENABLE_GITHUB_TOP = GetBool("ENABLE_GITHUB_TOP", false) // Enable Github Top
)
