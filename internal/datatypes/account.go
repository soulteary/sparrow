package datatypes

type Deactivate struct {
	Status string `json:"status"`
}

type DataExport struct {
	Status string `json:"status"`
}

type AccountCheck struct {
	AccountPlan AccountCheckPlan `json:"account_plan"`
	UserCountry string           `json:"user_country"`
	Features    []string         `json:"features"`
}

type AccountCheckPlan struct {
	IsPaidSubscriptionActive       bool   `json:"is_paid_subscription_active"`
	SubscriptionPlan               string `json:"subscription_plan"`
	AccountUserRole                string `json:"account_user_role"`
	WasPaidCustomer                bool   `json:"was_paid_customer"`
	HasCustomerObject              bool   `json:"has_customer_object"`
	SubscriptionExpiresAtTimestamp any    `json:"subscription_expires_at_timestamp"`
}

const (
	// stats
	FEATURE_LOG_STATSIG_EVENTS  = "log_statsig_events"  // Enabled by default, statistical analysis report
	FEATURE_LOG_INTERCOM_EVENTS = "log_intercom_events" // Enabled by default, statistical analysis report

	// data controls
	FEATURE_DATA_DELETION_ENABLE = "data_deletion_enabled" // Enabled by default, allowing account deletion
	FEATURE_DATA_EXPORT          = "data_export_enabled"   // Enabled by default, allows data to be exported
	FEATURE_DATA_CONTROL         = "data_controls_enabled" // Enabled by default, allows users to control data

	// messages
	FEATURE_DFW_MESSAGE_FEEDBACK     = "dfw_message_feedback"                // Enabled by default
	FEATURE_DFW_INLINE_MESSAGE_REGEN = "dfw_inline_message_regen_comparison" // Enabled by default
	FEATURE_SYSTEM_MESSAGE           = "system_message"                      // Enabled by default

	// account
	FEATURE_ONFOFF_STATUE_ACCOUNT                = "oneoff_status_account"                     // If you purchase a service during a service interruption, you will be prompted for a refund
	FEATURE_SHOW_EXISTING_USER_AGE_CONFIRM_MODAL = "show_existing_user_age_confirmation_modal" // 23.05.08 Added, display age confirmation pop-up window

	// models
	FEATURE_MODEL_SWITCHER      = "model_switcher"              // The model can be switched in the interface, and the Plus account is enabled by default
	FEATURE_MODEL_SWITCHER_0512 = "new_model_switcher_20230512" // New UI model switcher, OpenAI 23.05.12 Added
	FEATURE_MODEL_PREVIEWER     = "model_preview"               // Remind the limit when using the preview model, the Plus account is enabled by default
	FEATURE_ALLOW_BETA_FEATURES = "beta_features"               // 23.05.14 Added, allow use beta features

	// misc
	FEATURE_DISABLE_UPGRADE_UI = "disable_upgrade_ui"      // Enabled by default except for Plus accounts
	FEATURE_DISABLE_HISTORY    = "disable_history"         // Disable the session history, only the interface is reflected
	FEATURE_BUCKETED_HISTORY   = "bucketed_history"        // Enabled by default, Display history in buckets
	FEATURE_SCROLL_HISTORY     = "infinite_scroll_history" // Enabled by default, infinite scroll history

	FEATURE_SHAREABLE_LINKS = "shareable_links" // 23.05.08 Added, conversation sharing feat

	// plugins
	FEATURE_PLUGINS_AVAILABLE = "plugins_available" // Plug-in permissions
	FEATURE_PLIGIN_ENABLE     = "tools3"            // Plug-in permissions
	FEATURE_PLUGINS_DISABLED  = "plugins_disabled"  // Plug-in permissions
	// plugins, browsing
	FEATURE_PLIGIN_BROWSING            = "tools"              // Plug-in permissions
	FEATURE_PLUGINS_BROWSING_AVAILABLE = "browsing_available" // Plug-in permissions
	FEATURE_PLUGINS_BROWSING_DISABLED  = "browsing_disabled"  // Plug-in permissions
	// plugins, code interpreter
	FEATURE_PLIGIN_CODE_INTERPRETER = "tools2" // Plug-in permissions
	// plugins, debug
	FEATURE_PLIGIN_PLUGIN_ADMIN = "tools3_admin" // Plug-in permissions
	FEATURE_PLIGIN_PLUGIN_DEV   = "tools3_dev"   // Plug-in permissions
	FEATURE_DEBUG               = "debug"        // Developer permissions, debug mode

	FEATURE_SYSTEM_MESSAGE2 = "system_message2"
)
