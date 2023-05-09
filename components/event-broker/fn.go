package EventBroker

import (
	"fmt"
	"strings"
)

func IsLastMessage(payload any) bool {
	return strings.EqualFold(strings.TrimSpace(fmt.Sprintf("%s", payload)), "[DONE]")
}
