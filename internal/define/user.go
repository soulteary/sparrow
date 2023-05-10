package define

import (
	"fmt"
)

const ENABLE_MOCK = true

const (
	MOCK_USER_ID     = "user-ABCD10241024102410241024"
	MOCK_USER_EMAIL  = "soulteary@gmail.com"
	MOCK_USER_NAME   = "soulteary"
	MOCK_USER_REGION = "CN"
)

var (
	MOCK_USER_IMAGE        = WEB_CLIENT_HOSTNAME + "/assets/avatar.jpg"
	MOCK_USER_TOKEN        = fmt.Sprintf("%s.%s.%s-D-%s_g", GenerateRandomString(120), GenerateRandomString(648), GenerateRandomString(152), GenerateRandomString(185))
	MOCK_USER_IDP          = "idp-provider-name"
	MOCK_INTERCOM_HASH     = GenerateRandomString(64)
	MOCK_SENTRY_TRACE_DATA = GenerateRandomString(49) + "-1"
	MOCK_SENTRY_TRACE_ID   = "sentry-environment=production,sentry-release=" + GenerateRandomString(40) + ",sentry-transaction=%2F,sentry-public_key=" + GenerateRandomString(32) + ",sentry-trace_id=" + GenerateRandomString(32) + ",sentry-sample_rate=1"
)
