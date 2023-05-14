package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetFeatures() []string {
	features := []string{
		// datatypes.FEATURE_LOG_STATSIG_EVENTS,  // Disable the default statistical analysis function
		// datatypes.FEATURE_LOG_INTERCOM_EVENTS, // Disable the default statistical analysis function
		datatypes.FEATURE_DFW_MESSAGE_FEEDBACK,
		datatypes.FEATURE_DFW_INLINE_MESSAGE_REGEN,
		datatypes.FEATURE_SYSTEM_MESSAGE,
		datatypes.FEATURE_SHOW_EXISTING_USER_AGE_CONFIRM_MODAL,
		datatypes.FEATURE_ALLOW_BETA_FEATURES,
	}

	if define.ENABLE_MODEL_SWITCH || define.ENABLE_MODEL_SWITCH_0512 {
		features = append(features, datatypes.FEATURE_MODEL_SWITCHER)
		features = append(features, datatypes.FEATURE_MODEL_PREVIEWER)
	}
	if define.ENABLE_MODEL_SWITCH_0512 {
		features = append(features, datatypes.FEATURE_MODEL_SWITCHER_0512)
	}

	if define.ENABLE_DATA_CONTROL {
		features = append(features, datatypes.FEATURE_DATA_DELETION_ENABLE)
		features = append(features, datatypes.FEATURE_DATA_EXPORT)
		features = append(features, datatypes.FEATURE_DATA_CONTROL)
	}

	if define.DEV_MODE {
		features = append(features, datatypes.FEATURE_DEBUG)
		features = append(features, datatypes.FEATURE_SHAREABLE_LINKS)
		// OpenAI internal debugging interface
		// features = append(features, datatypes.FEATURE_SYSTEM_MESSAGE2)
	}

	if !define.ENABLE_HISTORY_LIST {
		features = append(features, datatypes.FEATURE_DISABLE_HISTORY)
		features = append(features, datatypes.FEATURE_SCROLL_HISTORY)
	} else {
		features = append(features, datatypes.FEATURE_BUCKETED_HISTORY)
	}

	if define.ENABLE_PLUGIN {
		features = append(features, datatypes.FEATURE_PLIGIN_BROWSING) // No need after 0427
		features = append(features, datatypes.FEATURE_PLIGIN_CODE)
		features = append(features, datatypes.FEATURE_PLUGINS_AVAILABLE)
		features = append(features, datatypes.FEATURE_PLUGINS_BROWSING_AVAILABLE)

		features = append(features, datatypes.FEATURE_PLIGIN_PLUGIN)
		features = append(features, datatypes.FEATURE_PLIGIN_PLUGIN_ADMIN)
		features = append(features, datatypes.FEATURE_PLIGIN_PLUGIN_DEV)
	} else {
		features = append(features, datatypes.FEATURE_PLUGINS_DISABLED)
		features = append(features, datatypes.FEATURE_PLUGINS_BROWSING_DISABLED)
	}
	return features
}

func Info() datatypes.AccountCheck {
	plan := datatypes.AccountCheckPlan{
		AccountUserRole:                "account-owner",
		SubscriptionExpiresAtTimestamp: nil, // or: "null", 199199199199
	}

	if !define.ENABLE_PAID_SUBSCRIPTION {
		plan.IsPaidSubscriptionActive = true
		plan.WasPaidCustomer = true
		plan.HasCustomerObject = true
		plan.SubscriptionPlan = "chatgptplusplan" // or: "chatgptplusfreeplan"
	} else {
		plan.IsPaidSubscriptionActive = false
		plan.WasPaidCustomer = false
		plan.HasCustomerObject = false
		plan.SubscriptionPlan = ""
	}

	return datatypes.AccountCheck{
		UserCountry: define.MOCK_USER_REGION,
		Features:    GetFeatures(),
		AccountPlan: plan,
	}
}

func AccountCheck(c *gin.Context) {
	c.JSON(http.StatusOK, Info())
}
