package define

import "time"

const (
	DEFAULT_APP_PORT              = 8091                     // Application listening port
	DEFAULT_WEB_CLIENT_HOSTNAME   = "http://localhost:8090"  // The resource domain name used by the web client
	DEFAULT_EVENT_BROKER_PATIENCE = time.Second * 1          // The maximum time to wait for the event broker to respond
	DEFAULT_OPENAI_API_PROXY_ADDR = ""                       // OpenAI API Proxy Address
	DEFAULT_MIDJOURNEY_API_ADDR   = "ws://localhost:8092/ws" // Midjourney API URL
	DEFAULT_CLAUDE_API_ADDR       = "ws://localhost:8093/ws" // Claude API URL
)
