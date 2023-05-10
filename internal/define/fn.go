package define

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func GetBool(envKey string, def bool) bool {
	env := strings.ToLower(strings.TrimSpace(os.Getenv(envKey)))
	if env == "" {
		return def
	}
	return env == "on" || env == "true" || env == "1"
}

func GetPortString(envKey string, def int) string {
	env := strings.ToLower(strings.TrimSpace(os.Getenv(envKey)))
	if env != "" {
		var REGEXP_PURE_NUMBER = regexp.MustCompile(`^\d{4,5}$`)
		appPortMatchFromEnv := REGEXP_PURE_NUMBER.FindAllString(env, -1)
		if len(appPortMatchFromEnv) == 1 {
			return fmt.Sprintf(":%s", appPortMatchFromEnv[0])
		}
	}
	return fmt.Sprintf(":%d", def)
}

func GetHostName(envKey string, def string) string {
	env := strings.ToLower(strings.TrimSpace(os.Getenv(envKey)))
	if env == "" {
		return def
	}

	u, err := url.Parse(env)
	if !(err == nil && u.Scheme != "" && u.Host != "") {
		return def
	}

	if u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "ws" && u.Scheme != "wss" {
		return def
	}

	return strings.TrimSuffix(fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path), "/")
}

func GenerateRandomString(length int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var result string
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(err)
		}
		result += string(letters[n.Int64()])
	}
	return result
}
