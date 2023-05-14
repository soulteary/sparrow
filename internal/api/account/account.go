package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

// Official default function
func GetDefaultFeatures() []string {
	return []string{
		// datatypes.FEATURE_LOG_STATSIG_EVENTS,  // Disable the default statistical analysis function
		// datatypes.FEATURE_LOG_INTERCOM_EVENTS, // Disable the default statistical analysis function
		datatypes.FEATURE_DFW_MESSAGE_FEEDBACK,
		datatypes.FEATURE_DFW_INLINE_MESSAGE_REGEN,
		datatypes.FEATURE_SYSTEM_MESSAGE,
		datatypes.FEATURE_SHOW_EXISTING_USER_AGE_CONFIRM_MODAL,
	}
}

func GetDataControlsFeatures() []string {
	if define.ENABLE_DATA_CONTROL {
		return []string{
			datatypes.FEATURE_DATA_CONTROL,
			datatypes.FEATURE_DATA_EXPORT,
			datatypes.FEATURE_DATA_DELETION_ENABLE,
		}
	}
	return []string{}
}

func GetHistoryFeatures() []string {
	if define.ENABLE_HISTORY_LIST {
		return []string{
			datatypes.FEATURE_BUCKETED_HISTORY,
			datatypes.FEATURE_SCROLL_HISTORY,
		}
	}
	return []string{datatypes.FEATURE_DISABLE_HISTORY}
}

func GetUIFeatures() []string {
	if define.ENABLE_NEW_UI {
		return []string{
			datatypes.FEATURE_MODEL_SWITCHER_0512,
			datatypes.FEATURE_MODEL_PREVIEWER,
		}
	}
	return []string{datatypes.FEATURE_MODEL_PREVIEWER}
}

func GetModelSwitcherFeatures() []string {
	if define.ENABLE_MODEL_SWITCH {
		return []string{
			datatypes.FEATURE_MODEL_SWITCHER,
		}
	}
	return []string{}
}

func GetPluginFeatures() []string {
	features := []string{}

	if define.ENABLE_PLUGIN_BROWSING {
		if define.NEW_TOOLS_USER {
			features = append(features, datatypes.FEATURE_PLUGINS_BROWSING_AVAILABLE)
		} else {
			features = append(features, datatypes.FEATURE_PLUGINS_BROWSING_AVAILABLE)
			features = append(features, datatypes.FEATURE_PLIGIN_BROWSING)
		}
	} else {
		features = append(features, datatypes.FEATURE_PLUGINS_BROWSING_DISABLED)
	}

	if define.ENABLE_PLUGIN_CODE_INTERPRETER {
		features = append(features, datatypes.FEATURE_PLIGIN_CODE_INTERPRETER)
	}

	if define.ENABLE_PLUGIN {
		if define.NEW_TOOLS_USER {
			features = append(features, datatypes.FEATURE_PLUGINS_AVAILABLE)
		} else {
			features = append(features, datatypes.FEATURE_PLUGINS_AVAILABLE)
			features = append(features, datatypes.FEATURE_PLIGIN_ENABLE)
		}
	} else {
		features = append(features, datatypes.FEATURE_PLUGINS_DISABLED)
	}
	return features
}

func GetFeatures() []string {
	features := GetDefaultFeatures()
	features = append(features, datatypes.FEATURE_ALLOW_BETA_FEATURES)
	features = append(features, GetDataControlsFeatures()...)
	features = append(features, GetHistoryFeatures()...)

	features = append(features, GetUIFeatures()...)
	features = append(features, GetModelSwitcherFeatures()...)
	features = append(features, GetPluginFeatures()...)

	if define.DEV_MODE {
		features = append(features, datatypes.FEATURE_DEBUG)
		features = append(features, datatypes.FEATURE_SHAREABLE_LINKS)
		// OpenAI internal debugging interface
		// features = append(features, datatypes.FEATURE_SYSTEM_MESSAGE2)
		features = append(features, datatypes.FEATURE_PLIGIN_PLUGIN_ADMIN)
		features = append(features, datatypes.FEATURE_PLIGIN_PLUGIN_DEV)
	}
	return features
}

func AccountCheck(c *gin.Context) {
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

	data := datatypes.AccountCheck{
		UserCountry: define.MOCK_USER_REGION,
		Features:    GetFeatures(),
		AccountPlan: plan,
	}
	c.JSON(http.StatusOK, data)
}
